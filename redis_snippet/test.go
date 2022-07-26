package main

import (
	"fmt"

	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
)

func main() {
	// Connect to localhost with no password
	var client = redistimeseries.NewClient("localhost:6379", "", nil)
	fmt.Printf("%T", client)
	var keyname = "1:success"

	labels := map[string]string{
		"status": "success",
		"appId":  "1",
		"flagId": "1",
	}
	// get the default options and set the time-serie labels
	options := redistimeseries.DefaultCreateOptions
	options.Labels = labels
	client.AddAutoTsWithOptions(keyname, 1, options)
	// client.AddWithOptions(keyname, 2, 2, options)

	// Retrieve the latest data point
	latestDatapoint, _ := client.Get(keyname)

	fmt.Printf("Latest datapoint: timestamp=%d value=%f\n", latestDatapoint.Timestamp, latestDatapoint.Value)

	// fmt.Printf("Latest datapoint: timestamp=%d value=%f\n", latestDatapoint.Timestamp, latestDatapoint.Value)
}
