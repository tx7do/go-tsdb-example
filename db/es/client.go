package es

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

type Options struct {
	Addresses           []string
	Username            string
	Password            string
	SnifferEnabled      bool
	HealthCheckInterval string
}

type Client struct {
	cli *elastic.Client
	opt *Options
}

func NewClient(o *Options) *Client {
	c := &Client{}

	//ES 实例：对应 MySQL 实例中的一个 Database。
	//Index 对应 MySQL 中的 Table 。
	//Document 对应 MySQL 中表的记录。

	duration, err := time.ParseDuration(o.HealthCheckInterval)
	cli, err := elastic.NewClient(
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(o.Addresses...),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth(o.Username, o.Password),
		// 启用gzip压缩
		elastic.SetGzip(true),
		elastic.SetSniff(o.SnifferEnabled),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(duration),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ES ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, " ", log.LstdFlags)),
	)
	if err != nil {
		return nil
	}

	c.cli = cli
	c.opt = o

	return c
}

func (c *Client) IndexExists(ctx context.Context, indexName string) bool {
	exist, err := c.cli.IndexExists(indexName).Do(ctx)
	if err != nil {
		return true
	}
	return exist
}

// CreateIndex 创建一条索引
// @param[in] mapping 如果为空("")则表示不创建模型
func (c *Client) CreateIndex(ctx context.Context, indexName, mapping string) error {
	exist, err := c.cli.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("索引已经存在，无需重复创建")
	}

	// 创建索引
	builder := c.cli.CreateIndex(indexName)
	if len(mapping) != 0 {
		builder.BodyString(mapping)
	}
	_, err = builder.Do(ctx)

	if err != nil {
		return err
	}

	return nil
}

// DeleteIndex 删除一条索引
func (c *Client) DeleteIndex(ctx context.Context, indexName string) error {
	exist, err := c.cli.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("index already deleted")
	}

	deleteIndex, err := c.cli.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !deleteIndex.Acknowledged {
		return errors.New("delete index is not Acknowledged")
	}
	return nil
}

// DeleteData 删除一条数据
func (c *Client) DeleteData(ctx context.Context, indexName, id string) error {
	// 根据id删除一条数据
	_, err := c.cli.Delete().
		Index(indexName).
		Id(id).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

// InsertData 插入一条数据
func (c *Client) InsertData(ctx context.Context, indexName, id string, data interface{}) error {
	// Index a tweet (using JSON serialization)
	builder := c.cli.Index()
	builder.
		Index(indexName).
		BodyJson(data)

	if len(id) > 0 {
		builder.Id(id)
	}

	_, err := builder.Do(ctx)
	if err != nil {
		return err
	}
	_, err = c.cli.Flush().Index(indexName).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

// BatchInsertData 批量插入数据
func (c *Client) BatchInsertData(ctx context.Context, indexName string, dataSet []interface{}) error {
	bulk := c.cli.Bulk().Index(indexName)
	if bulk == nil {
		return errors.New("create bulk failed")
	}
	for i := 0; i < len(dataSet); i++ {
		doc := elastic.NewBulkIndexRequest().Doc(dataSet[i])
		bulk.Add(doc)
	}
	if bulk.NumberOfActions() < 0 {
		return errors.New("no data need save")
	}
	if _, err := bulk.Do(ctx); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateValues(ctx context.Context, indexName string, pk string, doc interface{}) bool {
	esUpdate := c.cli.Update()
	if len(indexName) > 0 {
		esUpdate = esUpdate.Index(indexName)
	}
	if len(pk) > 0 {
		esUpdate = esUpdate.Id(pk)
	}

	_, err := esUpdate.Doc(doc).Do(ctx)
	if err != nil {
		return false
	}
	return true
}

// QueryById 查询数据
func (c *Client) QueryById(ctx context.Context, indexName string, id string, out interface{}) error {
	got, err := c.cli.Get().
		Index(indexName).
		Id(id).
		Do(ctx)
	if err != nil {
		return err
	}

	data, err := got.Source.MarshalJSON()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return err
	}

	return nil
}

// Search 查询数据
// @param ctx 上下文
// @param indexName 索引名
// @param sortBy 排序
// @param from 分页的页码
// @param pageSize 分页每页的行数
// @param out 输出结果
func (c *Client) Search(ctx context.Context, indexName string, sortBy map[string]bool, from, pageSize int, out interface{}) error {
	searchResult := c.cli.Search().Index(indexName)

	// 排序
	for k, v := range sortBy {
		searchResult = searchResult.Sort(k, v)
	}

	//if len(query) > 0 {
	//	for _, v := range query {
	//		searchResult = searchResult.Query(v)
	//	}
	//}

	// 分页
	if pageSize > 0 {
		searchResult = searchResult.From(from)
		searchResult = searchResult.Size(pageSize)
	}

	//esResult, err := searchResult.Do(ctx) // execute
	//if err != nil {
	//	return err
	//}

	//	log.Println("wwwwww", es_result.Aggregations)
	//if esResult.TotalHits() > 0 {
	//	return nil, esResult.Hits.Hits
	//} else {
	//	return nil, nil
	//}

	return nil
}
