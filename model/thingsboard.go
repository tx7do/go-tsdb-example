package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

// AttributeKv 属性键值
type AttributeKv struct {
	EntityType          string    `json:"entity_type" bson:"entity_type"`
	EntityId            uuid.UUID `json:"entity_id" bson:"entity_id"`
	AttributeType       string    `json:"attribute_type" bson:"attribute_type"`
	AttributeKey        string    `json:"attribute_key" bson:"attribute_key"`
	LastUpdateTimestamp int64     `json:"last_update_ts" bson:"last_update_ts"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

// TsKv 遥测数据 历史数据 (时序数据)
type TsKv struct {
	EntityId  uuid.UUID `json:"entity_id" bson:"entity_id"`
	Key       int       `json:"key" bson:"key"`
	Timestamp int64     `json:"ts" bson:"ts"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

func NewTsKv(entityId uuid.UUID, keyId int, timestamp int64,
	boolV *bool, stringV *string, longV *int64, doubleV *float64, jsonV *string) *TsKv {
	c := &TsKv{
		EntityId:  entityId,
		Key:       keyId,
		Timestamp: timestamp,

		BoolV:   boolV,
		StringV: stringV,
		LongV:   longV,
		DoubleV: doubleV,
		JsonV:   jsonV,
	}
	return c
}

func NewTsKvWithKvPair(entityId uuid.UUID, keyId int, timestamp int64, pair *KvPair) *TsKv {
	return NewTsKv(entityId, keyId, timestamp, pair.BoolV, pair.StringV, pair.LongV, pair.DoubleV, pair.JsonV)
}

func (t *TsKv) GetValue() interface{} {
	if t.BoolV != nil {
		return *t.BoolV
	}
	if t.StringV != nil {
		return *t.StringV
	}
	if t.LongV != nil {
		return *t.LongV
	}
	if t.DoubleV != nil {
		return *t.DoubleV
	}
	if t.JsonV != nil {
		return *t.JsonV
	}
	return nil
}

// TsKvDictionary 键值字典
type TsKvDictionary struct {
	KeyId   int    `json:"key_id" bson:"_id"`      // 编号
	Key     string `json:"key" bson:"key"`         // 键名,如: humidity
	Display string `json:"display" bson:"display"` // 显示名,如: 湿度
	Unit    string `json:"unit" bson:"unit"`       // 单位,如: RH
}

// TsKvLatest 遥测数据 最新数据
type TsKvLatest struct {
	EntityId  uuid.UUID `json:"entity_id" bson:"entity_id"`
	KeyId     int       `json:"key_id" bson:"key_id"`
	Key       string    `json:"key" bson:"key"`
	Timestamp int64     `json:"ts" bson:"ts"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

func NewTsKvLatest(entityId uuid.UUID, keyId int, timestamp int64,
	boolV *bool, stringV *string, longV *int64, doubleV *float64, jsonV *string) *TsKvLatest {
	c := &TsKvLatest{
		EntityId:  entityId,
		KeyId:     keyId,
		Timestamp: timestamp,

		BoolV:   boolV,
		StringV: stringV,
		LongV:   longV,
		DoubleV: doubleV,
		JsonV:   jsonV,
	}
	return c
}

func NewTsKvLatestWithKvPair(entityId uuid.UUID, keyId int, timestamp int64, pair *KvPair) *TsKvLatest {
	return NewTsKvLatest(entityId, keyId, timestamp, pair.BoolV, pair.StringV, pair.LongV, pair.DoubleV, pair.JsonV)
}

func (t *TsKvLatest) GetValue() interface{} {
	if t.BoolV != nil {
		return *t.BoolV
	}
	if t.StringV != nil {
		return *t.StringV
	}
	if t.LongV != nil {
		return *t.LongV
	}
	if t.DoubleV != nil {
		return *t.DoubleV
	}
	if t.JsonV != nil {
		return *t.JsonV
	}
	return nil
}

// KvPair 键值对
type KvPair struct {
	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"` // BOOL类型值
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`   // String类型值
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"` // Long类型值
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`   // Double类型值
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"` // JSON类型值
}
type KvMap map[string]*KvPair

func (t *KvPair) GetValue() interface{} {
	if t.BoolV != nil {
		return *t.BoolV
	}
	if t.StringV != nil {
		return *t.StringV
	}
	if t.LongV != nil {
		return *t.LongV
	}
	if t.DoubleV != nil {
		return *t.DoubleV
	}
	if t.JsonV != nil {
		return *t.JsonV
	}
	return nil
}

func MarshalKv(value interface{}) KvMap {
	th := reflect.TypeOf(value)
	vh := reflect.ValueOf(value)

	var results = make(KvMap)
	for i := 0; i < th.NumField(); i++ {
		fieldType := th.Field(i)
		fieldValue := vh.FieldByName(fieldType.Name)

		tagJson := fieldType.Tag.Get("json")
		if idx := strings.Index(tagJson, ","); idx != -1 {
			tagJson = tagJson[:idx]
		}
		if len(tagJson) == 0 {
			tagJson = fieldType.Name
		}

		switch fieldType.Type.Kind() {
		case reflect.Bool:
			v := fieldValue.Interface().(bool)
			results[tagJson] = &KvPair{BoolV: &v}

		case reflect.String:
			v := fieldValue.Interface().(string)
			results[tagJson] = &KvPair{StringV: &v}

		case reflect.Int:
			v := int64(fieldValue.Interface().(int))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Int8:
			v := int64(fieldValue.Interface().(int8))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Int16:
			v := int64(fieldValue.Interface().(int16))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Int32:
			v := int64(fieldValue.Interface().(int32))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Int64:
			v := fieldValue.Interface().(int64)
			results[tagJson] = &KvPair{LongV: &v}

		case reflect.Uint:
			v := int64(fieldValue.Interface().(uint))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Uint8:
			v := int64(fieldValue.Interface().(uint8))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Uint16:
			v := int64(fieldValue.Interface().(uint16))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Uint32:
			v := int64(fieldValue.Interface().(uint32))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Uintptr:
			v := int64(fieldValue.Interface().(uintptr))
			results[tagJson] = &KvPair{LongV: &v}
		case reflect.Uint64:
			v := int64(fieldValue.Interface().(uint64))
			results[tagJson] = &KvPair{LongV: &v}

		case reflect.Float32:
			v := float64(fieldValue.Interface().(float32))
			results[tagJson] = &KvPair{DoubleV: &v}
		case reflect.Float64:
			v := fieldValue.Interface().(float64)
			results[tagJson] = &KvPair{DoubleV: &v}

		case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
			jsonByte, _ := json.Marshal(fieldValue.Interface())
			v := string(jsonByte)
			results[tagJson] = &KvPair{JsonV: &v}
		}
	}

	return results
}
