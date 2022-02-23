package model

// TsKv
type TsKv struct {
	EntityId  string  `json:"entity_id" bson:"_id,omitempty"`
	Key       int     `json:"key" bson:"key,omitempty"`
	Timestamp uint64  `json:"ts,omitempty" bson:"ts,omitempty"`
	BoolV     bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV   string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV     uint64  `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV   float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV     string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

// TsKvDictionary
type TsKvDictionary struct {
	Key   string `json:"key" bson:"_id,omitempty"`
	KeyId int    `json:"key_id,omitempty" bson:"key_id,omitempty"`
}

// TsKvLatest
type TsKvLatest struct {
	EntityId  string  `json:"entity_id" bson:"_id,omitempty"`
	Key       int     `json:"key" bson:"key,omitempty"`
	Timestamp uint64  `json:"ts,omitempty" bson:"ts,omitempty"`
	BoolV     bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV   string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV     uint64  `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV   float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV     string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

type GpsData struct {
	Latitude        float64 `json:"latitude" bson:"latitude,omitempty"`
	Longitude       float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
	BatteryLevel    int     `json:"batteryLevel,omitempty" bson:"batteryLevel,omitempty"`
	BatteryCharging bool    `json:"batteryCharging,omitempty" bson:"batteryCharging,omitempty"`
}
