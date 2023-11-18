package mongodb

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go-tsdb-example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://cloud.tencent.com/document/product/652/14631#.E6.99.AE.E9.80.9A.E8.81.9A.E5.90.88
// https://www.jianshu.com/p/30ba571b2937

var (
	client *Client
	ctx    context.Context
)

func init() {
	rand.Seed(time.Now().UnixNano())

	ctx = context.Background()
	opt := &Options{
		Address:  "localhost:27017",
		Username: "root",
		Password: "123456",
		DBName:   "test",
	}
	cli := NewClient(opt)
	client = cli
}

func TestNewMongoDBClient(t *testing.T) {
	assert.NotNil(t, client)
}

func TestInsertSensor(t *testing.T) {
	assert.NotNil(t, client)

	sensorTypes := []string{"a", "a", "b", "b"}
	sensorLocations := []string{"floor", "ceiling", "floor", "ceiling"}

	var sensor model.Sensor
	for i := range sensorTypes {
		sensor.Id = i + 1 // mongodb _id不能为0
		sensor.Type = sensorTypes[i]
		sensor.Location = sensorLocations[i]
		err := client.WriteData(ctx, "sensor", sensor)
		assert.Nil(t, err)
	}
}

func randomTimestamp(min, max time.Time) time.Time {
	sec := rand.Int63n(max.Unix()-min.Unix()) + min.Unix()
	return time.Unix(sec, 0)
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func stringToTimestamp(str string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	return theTime
}

func TestInsertSensorData(t *testing.T) {
	assert.NotNil(t, client)

	minTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	maxTime := time.Date(2021, 12, 30, 0, 0, 0, 0, time.UTC)

	var sensorData model.SensorData
	for i := 0; i < 100; i++ {
		sensorData.Id = primitive.NewObjectIDFromTimestamp(randomTimestamp(minTime, maxTime)).Hex()
		sensorData.SensorId = randomInt(1, 3)
		sensorData.Temperature = rand.Float64() * 100
		sensorData.CPU = rand.Float64()
		fmt.Println(sensorData)
		err := client.WriteData(ctx, "sensor_data", sensorData)
		assert.Nil(t, err)
	}
}

func TestQueryOneData(t *testing.T) {
	assert.NotNil(t, client)

	type Param struct {
		Id string `json:"id" bson:"_id,omitempty"`
	}

	var param Param
	param.Id = "6193dc9047f928e7ad67aa00"

	var err error
	var sensorData model.SensorData

	err = client.QueryOneData(ctx, "sensor_data", param, sensorData)
	assert.Nil(t, err)
}

type RangeFilter string
type TermsFilter string
type SortMode string
type SortDirection string

const (
	GteRange RangeFilter = "gte" // 大于或等于
	GtRange  RangeFilter = "gt"  // 大于
	LteRange RangeFilter = "lte" // 小于或等于
	LtRange  RangeFilter = "lt"  // 小于

	FilterTerm  TermsFilter = "filter"   // 类似于 AND
	MustNotTerm TermsFilter = "must_not" // 类似于 NOT
	ShouldTerm  TermsFilter = "should"   // 类似于 OR

	MinSortMode    SortMode = "min"
	MaxSortMode    SortMode = "max"
	SumSortMode    SortMode = "sum"
	AvgSortMode    SortMode = "avg"
	MedianSortMode SortMode = "median"

	DescendingDirection SortDirection = "desc"
	AscendingDirection  SortDirection = "asc"
)

func TestQueryData(t *testing.T) {
	assert.NotNil(t, client)

	filter := bson.D{
		{"sensor_id",
			bson.D{{
				"$gt", 1,
			}},
		},
	}

	var err error
	var sensorData model.SensorData

	err = client.QueryData(ctx, "sensor_data", filter, sensorData)
	assert.Nil(t, err)
}
