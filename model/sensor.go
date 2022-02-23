package model

import "time"

type Sensor struct {
	Id       int    `json:"id" bson:"_id,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
	Location string `json:"location,omitempty" bson:"location,omitempty"`
}

type SensorData struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Time        time.Time `json:"time" bson:"created,omitempty"`
	SensorId    int       `json:"sensor_id" bson:"sensor_id,omitempty"`
	Temperature float64   `json:"temperature" bson:"temperature,omitempty"`
	CPU         float64   `json:"cpu" bson:"cpu,omitempty"`
}

const SensorMapping = `
{
  "mappings": {
    "properties": {
      "sensor_id": {
        "type": "integer"
      },
      "temperature": {
        "type": "double"
      },
      "cpu": {
        "type": "double"
      },
      "location": {
        "type": "geo_point"
      }
    }
  }
}`
