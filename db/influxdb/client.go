package influxdb

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"math/rand"
	"time"
)

type Options struct {
	Address string
	Token   string
}

type Client struct {
	opt *Options
	cli influxdb2.Client
}

func NewClient(o *Options) *Client {
	c := &Client{}

	cli := influxdb2.NewClientWithOptions(o.Address, o.Token,
		influxdb2.DefaultOptions().SetBatchSize(20))

	c.cli = cli
	c.opt = o

	return c
}

func (c *Client) CreateBucket(ctx context.Context, orgName, bucketName string) error {
	org, err := c.cli.OrganizationsAPI().FindOrganizationByName(ctx, orgName)
	if err != nil {
		fmt.Printf("ERROR. Cannot find organization")
		return nil
	}

	bucketsAPI := c.cli.BucketsAPI()
	_, err = bucketsAPI.CreateBucketWithName(ctx, org, bucketName, domain.RetentionRule{EverySeconds: 3600 * 12})
	if err != nil {
		fmt.Printf("Error. Cannot create bucket")
		return err
	}
	return nil
}

func (c *Client) WriteData(ctx context.Context, orgName, bucketName string) error {
	bucketsAPI := c.cli.BucketsAPI()
	bucket, err := bucketsAPI.FindBucketByName(ctx, bucketName)
	if err != nil {
		fmt.Printf("Error. Cannot create bucket")
		return nil
	}

	writeAPI := c.cli.WriteAPI(orgName, bucket.Name)
	// Read and log errors
	errorsCh := writeAPI.Errors()
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()

	// write some points
	for i := 0; i < 100; i++ {
		// create point
		p := influxdb2.NewPoint(
			"rand-buck",
			map[string]string{
				"contID":   fmt.Sprintf("contID_%v", i),
				"contName": fmt.Sprintf("contName_%v", i),
				"vendor":   "mobigen",
			},
			map[string]interface{}{
				"utime":    rand.Float64(),
				"stime":    rand.Float64(),
				"cutime":   rand.Float64(),
				"cstime":   rand.Float64(),
				"rxByte":   rand.Float64(),
				"rxPacket": rand.Float64(),
				"txByte":   rand.Float64(),
				"txPacket": rand.Float64(),
				"vmsize":   rand.Float64(),
				"vmrss":    rand.Float64(),
				"rssfile":  rand.Float64(),
			},
			time.Now())
		// write asynchronously
		writeAPI.WritePoint(p)
	}
	// Force all unwritten data to be sent
	writeAPI.Flush()

	return nil
}

func (c *Client) QueryData(ctx context.Context, orgName, bucketName string) error {
	// Get query client
	queryAPI := c.cli.QueryAPI(orgName)
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"primary")
	|> range(start:-8h)
	|> filter(fn:(r) =>
		r._measurement == "rand-buck" and
		r._field == "utime" or r._field == "stime")
    |> pivot(rowKey:["_time"], columnKey:["_field"], valueColumn: "_value")
    |> map(fn: (r) => ({ r with _value: r.utime + r.stime}))
	|> yield(name: "_results")`)

	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Access data
			fmt.Printf("Time: %v\n", result.Record().Time())
			fmt.Printf("ContainerName: %v   |   ", result.Record().ValueByKey("contName"))
			fmt.Printf("utime + stime: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		fmt.Printf("ERROR. Cannot serve qeury result\n")
	}

	return nil
}
