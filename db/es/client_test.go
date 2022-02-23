package es

import (
	"cmd/model"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	client *Client
)

const (
	userIndex   = "user"
	weiboIndex  = "weibo"
	sensorIndex = "sensor"
)

func init() {
	opt := &Options{
		Addresses:           []string{"http://localhost:9200"},
		Username:            "elastic",
		Password:            "elastic",
		SnifferEnabled:      false,
		HealthCheckInterval: "5s",
	}
	cli := NewClient(opt)
	client = cli

	rand.Seed(time.Now().Unix())
}

func TestNewESClient(t *testing.T) {
	assert.NotNil(t, client)
}

func TestCreateIndex(t *testing.T) {
	assert.NotNil(t, client)
	var esCtx = context.Background()

	{
		//_ = client.DeleteIndex(esCtx, userIndex)
		err := client.CreateIndex(esCtx, userIndex, model.UserMapping)
		assert.Nil(t, err)
	}

	{
		//_ = client.DeleteIndex(esCtx, weiboIndex)
		err := client.CreateIndex(esCtx, weiboIndex, model.WeiboMapping)
		assert.Nil(t, err)
	}

	{
		//_ = client.DeleteIndex(esCtx, sensorIndex)
		err := client.CreateIndex(esCtx, sensorIndex, model.SensorMapping)
		assert.Nil(t, err)
	}
}

func TestDeleteIndex(t *testing.T) {
	assert.NotNil(t, client)
	var esCtx = context.Background()

	err := client.DeleteIndex(esCtx, userIndex)
	assert.Nil(t, err)

	err = client.DeleteIndex(esCtx, weiboIndex)
	assert.Nil(t, err)

	err = client.DeleteIndex(esCtx, sensorIndex)
	assert.Nil(t, err)
}

func TestInsertData(t *testing.T) {
	assert.NotNil(t, client)
	var esCtx = context.Background()

	{
		// http://localhost:9200/user/_search?q=*&pretty
		loc, _ := time.LoadLocation("Local")
		birth, _ := time.ParseInLocation("2006-01-02", "1991-04-25", loc)
		userOne := model.User{
			Name:   "张三",
			Age:    23,
			Phone:  "17600000000",
			Birth:  birth,
			Height: 170.5,
			Home:   "41.40338,2.17403",
		}

		err := client.InsertData(esCtx, userIndex, "", userOne)
		assert.Nil(t, err)
	}

	{
		// http://localhost:9200/weibo/_search?q=*&pretty
		weiboOne := model.Weibo{User: "olive", Message: "打酱油的一天", Retweets: 0}

		err := client.InsertData(esCtx, weiboIndex, "", weiboOne)
		assert.Nil(t, err)
	}
}

func TestBatchInsertData(t *testing.T) {
	assert.NotNil(t, client)
	var esCtx = context.Background()

	{
		loc, _ := time.LoadLocation("Local")
		// 生日
		birthSlice := []string{"1991-04-25", "1990-01-15", "1989-11-05", "1988-01-25", "1994-10-12"}
		// 姓名
		nameSlice := []string{"李四", "张飞", "赵云", "关羽", "刘备"}

		var users []interface{}
		for i := 1; i < 20; i++ {
			birth, _ := time.ParseInLocation("2006-01-02", birthSlice[rand.Intn(len(birthSlice))], loc)
			height, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float32()+175.0), 32)
			user := model.User{
				Name:   nameSlice[rand.Intn(len(nameSlice))],
				Age:    rand.Intn(10) + 18,
				Phone:  "1760000000" + strconv.Itoa(i),
				Birth:  birth,
				Height: float32(height),
				Home:   "41.40338,2.17403",
			}
			users = append(users, user)
		}

		err := client.BatchInsertData(esCtx, userIndex, users)
		assert.Nil(t, err)
	}
}

func TestQueryData(t *testing.T) {
	assert.NotNil(t, client)
	var esCtx = context.Background()
	var user model.User
	const id = "LnQIdn4BSA3IQpYKUfDY"
	err := client.QueryById(esCtx, userIndex, id, &user)
	assert.Nil(t, err)
}
