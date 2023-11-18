package opentsdb

import (
	"fmt"

	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/bluebreezecf/opentsdb-goclient/config"
)

type Options struct {
	Address  string
	Username string
	Password string
}

type Client struct {
	cli client.Client
	opt *Options
}

func NewClient(o *Options) *Client {
	c := &Client{}

	opentsdbCfg := config.OpenTSDBConfig{
		OpentsdbHost: o.Address,
	}
	cli, err := client.NewClient(opentsdbCfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}

	//0. Ping
	if err = cli.Ping(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	c.cli = cli
	c.opt = o

	return c
}
