package timescaledb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	cli *pgxpool.Pool

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
func (c *Client) Connect(ctx context.Context) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.Username, c.Password, c.Address, c.DBName)

	cli, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return err
	}

	c.cli = cli

	return nil
}

func (c *Client) ExecuteSQL(ctx context.Context, sql string, arguments ...interface{}) error {
	_, err := c.cli.Exec(ctx, sql, arguments...)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Query(ctx context.Context, sql string, args ...interface{}) ([]byte, error) {
	rows, err := c.cli.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}

	tableData := make([]map[string]interface{}, 0)
	for rows.Next() {
		va, _ := rows.Values()
		//fmt.Println(va)
		entry := make(map[string]interface{})
		for i, v := range columns {
			entry[v] = va[i]
		}
		tableData = append(tableData, entry)
	}

	jsonData, _ := json.Marshal(tableData)

	return jsonData, nil
}
