package influxdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	client *Client
	ctx    context.Context
)

func init() {
	ctx = context.Background()
	opt := &Options{
		Address: "http://localhost:8086",
		Token:   "admintoken123",
	}
	cli := NewClient(opt)
	client = cli
}

func TestNewInfluxClient(t *testing.T) {
	assert.NotNil(t, client)
}

func TestWriteData(t *testing.T) {
	assert.NotNil(t, client)

	err := client.WriteData(ctx, "primary", "rand-buck")
	assert.Nil(t, err)
}

func TestQueryData(t *testing.T) {
	assert.NotNil(t, client)
	err := client.QueryData(ctx, "primary", "rand-buck")
	assert.Nil(t, err)
}
