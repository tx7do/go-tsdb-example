package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"time"
)

type Client struct {
	cli clickhouse.Conn

	Address  string
	Username string
	Password string
	DBName   string
}

func NewClient(addr, username, password, dbName string) *Client {
	c := &Client{
		Address:  addr,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
	return c
}

// Connect 连接数据库
func (c *Client) Connect() error {
	// DSN clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&max_execution_time=60
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{c.Address},
		Auth: clickhouse.Auth{
			Database: c.DBName,
			Username: c.Username,
			Password: c.Password,
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return err
	}

	c.cli = conn

	return nil
}
