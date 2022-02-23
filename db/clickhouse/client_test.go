package clickhouse

import (
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var (
	client *Client
	ctx    context.Context
)

func init() {
	rand.Seed(time.Now().Unix())

	ctx = context.Background()

	host := "localhost:9000"
	username := ""
	password := ""
	dbName := "default"
	client = NewClient(host, username, password, dbName)

	_ = client.Connect()
}

func TestNewClickHouseDBClient(t *testing.T) {
	assert.NotNil(t, client)
}
