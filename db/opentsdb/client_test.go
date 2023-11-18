package opentsdb

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	client *Client
)

func init() {
	opt := &Options{
		Address: "127.0.0.1:4242",
	}
	cli := NewClient(opt)
	client = cli

	rand.Seed(time.Now().Unix())
}

func TestNewOpenTSDBClient(t *testing.T) {
	assert.NotNil(t, client)
}
