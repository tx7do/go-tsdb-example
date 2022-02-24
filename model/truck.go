package model

type Trucks struct {
	TruckId      int    `json:"truck_id" bson:"truck_id,omitempty"`
	Make         string `json:"make,omitempty" bson:"make,omitempty"`
	Model        string `json:"model,omitempty" bson:"model,omitempty"`
	WeightClass  string `json:"weight_class,omitempty" bson:"weight_class,omitempty"`
	DateAcquired uint64 `json:"date_acquired,omitempty" bson:"date_acquired,omitempty"`
	ActiveStatus bool   `json:"active_status,omitempty" bson:"active_status,omitempty"`
}

type TruckReading struct {
	LastTime  uint64  `json:"ts,omitempty" bson:"ts,omitempty"`
	TruckId   int     `json:"truck_id" bson:"truck_id,omitempty"`
	Milage    int     `json:"milage,omitempty" bson:"milage,omitempty"`
	Fuel      int     `json:"fuel,omitempty" bson:"fuel,omitempty"`
	Latitude  float64 `json:"latitude" bson:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
}

type TruckLog struct {
	TruckId   int     `json:"truck_id" bson:"truck_id,omitempty"`
	LastTime  uint64  `json:"last_time,omitempty" bson:"last_time,omitempty"`
	Milage    int     `json:"milage,omitempty" bson:"milage,omitempty"`
	Fuel      int     `json:"fuel,omitempty" bson:"fuel,omitempty"`
	Latitude  float64 `json:"latitude" bson:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
}
