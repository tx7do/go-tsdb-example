package model

import "time"

type User struct {
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Phone  string    `json:"phone"`
	Birth  time.Time `json:"birth"`
	Height float32   `json:"height"`
	Smoke  bool      `json:"smoke"`
	Home   string    `json:"home"`
}

// UserMapping 定义用户mapping
const UserMapping = `
{
    "mappings":{
        "properties":{
            "name":{
                "type":"text"
            },
            "age":{
                "type":"byte"
            },
            "phone":{
                "type":"text"
            },
            "birth":{
                "type":"date"
            },
            "height":{
                "type":"float"
            },
            "smoke":{
                "type":"boolean"
            },
            "home":{
                "type":"geo_point"
            }
        }
    }
}`
