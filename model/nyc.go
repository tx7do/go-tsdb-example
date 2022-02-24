package model

type Rides struct {
	VendorId             string  `json:"vendor_id,omitempty" bson:"vendor_id,omitempty"`
	PickupDatetime       uint64  `json:"pickup_datetime,omitempty" bson:"pickup_datetime,omitempty"`
	DropoffDatetime      uint64  `json:"dropoff_datetime,omitempty" bson:"dropoff_datetime,omitempty"`
	PassengerCount       int     `json:"passenger_count,omitempty" bson:"passenger_count,omitempty"`
	TripDistance         float64 `json:"trip_distance,omitempty" bson:"trip_distance,omitempty"`
	PickupLongitude      float64 `json:"pickup_longitude,omitempty" bson:"pickup_longitude,omitempty"`
	PickupLatitude       float64 `json:"pickup_latitude,omitempty" bson:"pickup_latitude,omitempty"`
	RateCode             float64 `json:"rate_code,omitempty" bson:"rate_code,omitempty"`
	DropoffLongitude     float64 `json:"dropoff_longitude,omitempty" bson:"dropoff_longitude,omitempty"`
	DropoffLatitude      float64 `json:"dropoff_latitude,omitempty" bson:"dropoff_latitude,omitempty"`
	PaymentType          int     `json:"payment_type,omitempty" bson:"payment_type,omitempty"`
	FareAmount           float64 `json:"fare_amount,omitempty" bson:"fare_amount,omitempty"`
	Extra                float64 `json:"extra,omitempty" bson:"extra,omitempty"`
	MtaTax               float64 `json:"mta_tax,omitempty" bson:"mta_tax,omitempty"`
	TipAmount            float64 `json:"tip_amount,omitempty" bson:"tip_amount,omitempty"`
	TollsAmount          float64 `json:"tolls_amount,omitempty" bson:"tolls_amount,omitempty"`
	ImprovementSurcharge float64 `json:"improvement_surcharge,omitempty" bson:"improvement_surcharge,omitempty"`
	TotalAmount          float64 `json:"total_amount,omitempty" bson:"total_amount,omitempty"`
}

type PaymentTypes struct {
	PaymentType int    `json:"payment_type,omitempty" bson:"payment_type,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

type Rates struct {
	RateCode    int    `json:"rate_code,omitempty" bson:"rate_code,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}
