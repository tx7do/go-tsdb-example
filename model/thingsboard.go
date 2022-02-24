package model

// AttributeKv 属性键值
type AttributeKv struct {
	EntityType          string  `json:"entity_type" bson:"entity_type,omitempty"`
	EntityId            string  `json:"entity_id" bson:"_id,omitempty"`
	AttributeType       string  `json:"attribute_type" bson:"attribute_type,omitempty"`
	AttributeKey        string  `json:"attribute_key" bson:"attribute_key,omitempty"`
	LastUpdateTimestamp int64   `json:"last_update_ts,omitempty" bson:"last_update_ts,omitempty"`
	BoolV               bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV             string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV               int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV             float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV               string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

// TsKv 遥测数据 历史数据 (时序数据)
type TsKv struct {
	EntityId  string   `json:"entity_id" bson:"_id,omitempty"`
	Key       int      `json:"key" bson:"key,omitempty"`
	Timestamp int64    `json:"ts,omitempty" bson:"ts,omitempty"`
	BoolV     *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV   *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV     *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV   *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV     *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

// TsKvDictionary 键值字典
type TsKvDictionary struct {
	Key   string `json:"key" bson:"_id,omitempty"`
	KeyId int    `json:"key_id,omitempty" bson:"key_id,omitempty"`
}

// TsKvLatest 遥测数据 最新数据
type TsKvLatest struct {
	EntityId  string  `json:"entity_id" bson:"_id,omitempty"`
	Key       int     `json:"key" bson:"key,omitempty"`
	Timestamp int64   `json:"ts,omitempty" bson:"ts,omitempty"`
	BoolV     bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV   string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV     int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV   float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV     string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

type GpsData struct {
	Latitude        float64 `json:"latitude" bson:"latitude,omitempty"`
	Longitude       float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
	BatteryLevel    int     `json:"batteryLevel,omitempty" bson:"batteryLevel,omitempty"`
	BatteryCharging bool    `json:"batteryCharging,omitempty" bson:"batteryCharging,omitempty"`
}
