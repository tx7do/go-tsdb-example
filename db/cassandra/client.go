package cassandra

import (
	"errors"
	"github.com/gocql/gocql"
	"time"
)

type Client struct {
	cli *gocql.Session

	Address  string // <HOST:PORT>
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
	clusterConfig := gocql.NewCluster(c.Address)

	// 设置用户名密码
	clusterConfig.Authenticator = gocql.PasswordAuthenticator{
		Username: c.Username,
		Password: c.Password,
	}

	clusterConfig.Keyspace = c.DBName

	// 设置ssl
	//clusterConfig.SslOpts = &gocql.SslOptions{Config: &tls.Config{MinVersion: tls.VersionTLS12}}

	// 设置超时时间
	clusterConfig.ConnectTimeout = 10 * time.Second
	clusterConfig.Timeout = 10 * time.Second

	clusterConfig.Consistency = gocql.LocalQuorum

	// 禁止主机查找
	clusterConfig.DisableInitialHostLookup = true

	session, err := clusterConfig.CreateSession()
	if err != nil {
		return err
	}

	c.cli = session
	return nil
}

func (c *Client) Close() {
	if c.cli != nil {
		c.cli.Close()
		c.cli = nil
	}
}

// ExecuteQuery executes a query and returns an error if any
func (c *Client) ExecuteQuery(query string) error {
	if c.cli == nil {
		return errors.New("client not connect server")
	}
	return c.cli.Query(query).Exec()
}
