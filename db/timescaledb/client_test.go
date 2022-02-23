package timescaledb

import (
	"cmd/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	client *Client
	ctx    context.Context
)

// https://www.infoq.cn/article/nwftctvhwgesd-wxijcf
// https://blog.timescale.com/blog/select-the-most-recent-record-of-many-items-with-postgresql/

func init() {
	rand.Seed(time.Now().Unix())

	ctx = context.Background()

	host := "localhost:5432"
	username := "postgres"
	password := "123456"
	dbName := "postgres"
	client = NewClient(host, username, password, dbName)

	_ = client.Connect(ctx)
}

func TestNewTimeScaleDBClient(t *testing.T) {
	assert.NotNil(t, client)
}

func TestCreateTable(t *testing.T) {
	// https://docs.timescale.com/timescaledb/latest/quick-start/golang/
	queryCreateTable := `CREATE TABLE sensors (
		id SERIAL PRIMARY KEY,
		type VARCHAR(50), location VARCHAR(50)
		);`
	_ = client.ExecuteSQL(ctx, queryCreateTable)

	queryCreateHypertable := `CREATE TABLE sensor_data (
       time TIMESTAMPTZ NOT NULL,
       sensor_id INTEGER,
       temperature DOUBLE PRECISION,
       cpu DOUBLE PRECISION,
       FOREIGN KEY (sensor_id) REFERENCES sensors (id)
       );
       SELECT create_hypertable('sensor_data', 'time');       
       `
	_ = client.ExecuteSQL(ctx, queryCreateHypertable)
}

func TestInsertDataRDB(t *testing.T) {
	sensorTypes := []string{"a", "a", "b", "b"}
	sensorLocations := []string{"floor", "ceiling", "floor", "ceiling"}

	for i := range sensorTypes {
		//INSERT statement in SQL
		queryInsertMetadata := `INSERT INTO sensors (type, location) VALUES ($1, $2);`

		//Execute INSERT command
		err := client.ExecuteSQL(ctx, queryInsertMetadata, sensorTypes[i], sensorLocations[i])
		assert.Nil(t, err)

		fmt.Printf("Inserted sensor (%s, %s) into database \n", sensorTypes[i], sensorLocations[i])
	}

}

func TestInsertDataTSDB(t *testing.T) {
	queryDataGeneration := `
       SELECT generate_series(now() - interval '24 hour', now(), interval '5 minute') AS time,
       floor(random() * (3) + 1)::int as sensor_id,
       random()*100 AS temperature,
       random() AS cpu
       `
	//Execute query to generate samples for sensor_data hypertable
	jsons, err := client.Query(ctx, queryDataGeneration)
	assert.Nil(t, err)

	var results []model.SensorData
	err = json.Unmarshal(jsons, &results)
	assert.Nil(t, err)

	//SQL query to generate sample data
	queryInsertTimeSeriesData := `INSERT INTO sensor_data (time, sensor_id, temperature, cpu) VALUES ($1, $2, $3, $4);`

	//Insert contents of results slice into TimescaleDB
	for i := range results {
		var r model.SensorData
		r = results[i]
		err := client.ExecuteSQL(ctx, queryInsertTimeSeriesData, r.Time, r.SensorId, r.Temperature, r.CPU)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to insert sample into Timescale %v\n", err)
		}
	}

}

func TestQueryDataGenerateSeries(t *testing.T) {
	queryDataGeneration := `
       SELECT generate_series(now() - interval '24 hour', now(), interval '5 minute') AS time,
       floor(random() * (3) + 1)::int as sensor_id,
       random()*100 AS temperature,
       random() AS cpu
       `
	//Execute query to generate samples for sensor_data hypertable
	jsons, err := client.Query(ctx, queryDataGeneration)
	assert.Nil(t, err)

	var results []model.SensorData
	err = json.Unmarshal(jsons, &results)
	assert.Nil(t, err)

	fmt.Println(results)
}

func TestQueryDataTSDB(t *testing.T) {
	// https://docs.timescale.com/timescaledb/latest/tutorials/simulate-iot-sensor-data
	// Note the use of prepared statement placeholders $1 and $2
	queryTimeBucketFiveMin := `
       SELECT time_bucket('5 minutes', time) AS five_min, avg(cpu)
       FROM sensor_data
       JOIN sensors ON sensors.id = sensor_data.sensor_id
       WHERE sensors.location = $1 AND sensors.type = $2
       GROUP BY five_min
       ORDER BY five_min DESC;
       `

	//Execute query to generate samples for sensor_data hypertable
	jsons, err := client.Query(ctx, queryTimeBucketFiveMin, "ceiling", "a")
	assert.Nil(t, err)

	type AvgResult struct {
		Bucket time.Time `json:"five_min"`
		Avg    float64   `json:"avg"`
	}

	var results []AvgResult
	err = json.Unmarshal(jsons, &results)
	assert.Nil(t, err)
	//fmt.Println(results)

	for _, v := range results {
		fmt.Printf("Time bucket: %s | Avg: %f\n", &v.Bucket, v.Avg)
	}
}
