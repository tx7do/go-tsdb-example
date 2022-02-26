package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalKv(t *testing.T) {
	type SubStruct struct {
		A bool    `json:"a,omitempty"`
		B string  `json:"b"`
		C int64   `json:"c"`
		D float64 `json:"d"`
	}
	type TestStruct struct {
		A bool    `json:"a,omitempty"`
		B string  `json:"b"`
		C int64   `json:"c"`
		D float64 `json:"d"`
		E SubStruct
	}

	var testStruct TestStruct
	testStruct.A = true
	testStruct.B = "test"
	testStruct.C = 1
	testStruct.D = 2

	testStruct.E.A = true
	testStruct.E.B = "test"
	testStruct.E.C = 1
	testStruct.E.D = 2

	testMap := MarshalKv(testStruct)
	assert.Equal(t, len(testMap), 5)

	for key := range testMap {
		fmt.Println(key, testMap[key])
	}

	boolV, ok := testMap["a"]
	assert.True(t, ok)
	assert.Equal(t, *boolV.BoolV, true)

	stringV, ok := testMap["b"]
	assert.True(t, ok)
	assert.Equal(t, *stringV.StringV, "test")

	longV, ok := testMap["c"]
	assert.True(t, ok)
	assert.Equal(t, *longV.LongV, int64(1))

	doubleV, ok := testMap["d"]
	assert.True(t, ok)
	assert.Equal(t, *doubleV.DoubleV, float64(2))

	jsonV, ok := testMap["E"]
	assert.True(t, ok)
	jsonStr, err := json.Marshal(testStruct.E)
	assert.Nil(t, err)
	assert.Equal(t, *jsonV.JsonV, string(jsonStr))
}
