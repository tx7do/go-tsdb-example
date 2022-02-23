package cassandra

import (
	"math/rand"
	"time"
)

var (
	client *Client
)

func init() {
	rand.Seed(time.Now().Unix())

	host := "localhost:5432"
	username := "cassandra"
	password := "cassandra"
	dbName := "test"
	client = NewClient(host, username, password, dbName)

	_ = client.Connect()
}
