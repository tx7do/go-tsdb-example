package cassandra

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go-tsdb-example/model"
)

var (
	client *Client
)

func init() {
	rand.Seed(time.Now().Unix())

	host := "localhost:9042"
	username := "cassandra"
	password := "cassandra"
	dbName := "thingsboard"
	client = NewClient(host, username, password, dbName)

	_ = client.Connect()
}

func TestTelemetry(t *testing.T) {
	entityId := "ad2bfe60-7514-11ec-9a90-af0223be0666"
	timestamp := time.Now().UnixNano()

	var humidity = 56.4
	var temperature = 20.0

	{
		keyId := "humidity"
		kv := convertToTsKv(entityId, timestamp, humidity)

		err := saveOrUpdateTsKv(keyId, kv)
		assert.Nil(t, err)
		err = saveOrUpdateTsKvLatest(keyId, kv)
		assert.Nil(t, err)
	}

	{
		keyId := "temperature"
		kv := convertToTsKv(entityId, timestamp, temperature)
		err := saveOrUpdateTsKv(keyId, kv)
		assert.Nil(t, err)
		err = saveOrUpdateTsKvLatest(keyId, kv)
		assert.Nil(t, err)
	}
}

func TestToPartitionTs(t *testing.T) {
	assert.Equal(t, toPartitionTs(1645714787540161800), int64(1645714787540))
}

func convertToTsKv(entityId string, timestamp int64, value interface{}) *model.TsKv {
	entityUUid, _ := uuid.Parse(entityId)
	var kv model.TsKv
	kv.EntityId = entityUUid
	kv.Timestamp = timestamp
	switch t := value.(type) {
	case bool:
		kv.BoolV = &t
	case string:
		kv.StringV = &t
	case int64:
		kv.LongV = &t
	case float64:
		kv.DoubleV = &t
	}
	return &kv
}

func getValueOrNull(value interface{}) interface{} {
	switch t := value.(type) {
	case *bool:
		if t == nil {
			return nil
		} else {
			return *t
		}
	case *string:
		if t == nil {
			return nil
		} else {
			return *t
		}
	case *int64:
		if t == nil {
			return nil
		} else {
			return *t
		}
	case *float64:
		if t == nil {
			return nil
		} else {
			return *t
		}
	}
	return nil
}

func toPartitionTs(ts int64) int64 {
	var partition int64 = 0
	partition = time.Unix(0, ts).UnixMilli()
	return partition
}

func saveOrUpdateTsKv(key string, value *model.TsKv) error {
	partition := toPartitionTs(value.Timestamp)
	_ = saveOrUpdateTsKvPartition(value.EntityId.String(), key, partition)
	sql := `
INSERT INTO thingsboard.ts_kv_cf (entity_type, entity_id, key, partition, ts, bool_v, str_v, long_v, dbl_v, json_v)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	  `
	return client.ExecuteQuery(sql, "DEVICE", value.EntityId, key, partition, value.Timestamp,
		getValueOrNull(value.BoolV), getValueOrNull(value.StringV), getValueOrNull(value.LongV), getValueOrNull(value.DoubleV), getValueOrNull(value.JsonV))
}

func saveOrUpdateTsKvPartition(entityId, key string, partition int64) error {
	sql := `
INSERT INTO thingsboard.ts_kv_partitions_cf (entity_type, entity_id, key, partition)
VALUES (?, ?, ?, ?);
	  `
	return client.ExecuteQuery(sql, "DEVICE", entityId, key, partition)
}

func saveOrUpdateTsKvLatest(key string, value *model.TsKv) error {
	sql := `
INSERT INTO thingsboard.ts_kv_latest_cf (entity_type, entity_id, key, ts, bool_v, str_v, long_v, dbl_v, json_v)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	  `
	return client.ExecuteQuery(sql, "DEVICE", value.EntityId, key, value.Timestamp,
		getValueOrNull(value.BoolV), getValueOrNull(value.StringV), getValueOrNull(value.LongV), getValueOrNull(value.DoubleV), getValueOrNull(value.JsonV))
}
