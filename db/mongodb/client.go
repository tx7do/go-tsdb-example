package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	Address  string
	Username string
	Password string
	DBName   string
}

type Client struct {
	cli *mongo.Client
	opt *Options
}

// NewClient 创建MongoDB客户端
// https://docs.mongodb.com/drivers/go/current/quick-start/
func NewClient(o *Options) *Client {
	c := &Client{}

	var ctx = context.Background()

	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s",
			o.Username, o.Password, o.Address,
		),
	)

	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}

	c.cli = cli
	c.opt = o

	return c
}

// WriteData 插入数据
// @param [in] ctx 上下文
// @param [in] collName 集合名
// @param [in] doc 文档,最终会被转成bson
func (c *Client) WriteData(ctx context.Context, collName string, doc interface{}) error {
	coll := c.cli.Database(c.opt.DBName).Collection(collName)
	if coll == nil {
		return errors.New("collection not found")
	}

	_, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}

// QueryOneData 查询一条数据
func (c *Client) QueryOneData(ctx context.Context, collName string, filter interface{}, outResult interface{}) error {
	coll := c.cli.Database(c.opt.DBName).Collection(collName)
	if coll == nil {
		return errors.New("collection not found")
	}

	ret := coll.FindOne(ctx, filter)
	err := ret.Err()
	if err != nil {
		return err
	}

	var result bson.M
	err = ret.Decode(&result)
	output, err := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", output)

	//bsonResult, err := ret.DecodeBytes()
	//_ = bson.Unmarshal(bsonResult, &outResult)

	//var outResult interface{}
	//err = ret.Decode(&outResult)
	return err
}

// QueryData 查询多条数据
func (c *Client) QueryData(ctx context.Context, collName string, filter interface{}, outResult interface{}) error {
	coll := c.cli.Database(c.opt.DBName).Collection(collName)
	if coll == nil {
		return errors.New("collection not found")
	}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	return err
}
