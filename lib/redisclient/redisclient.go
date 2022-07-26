package redisTimeSeriesClient

import (
	"fmt"

	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
)

type RedisTimeSeriesClient struct {
	host        string
	port        int
	redisClient *redistimeseries.Client
}

func New(host string, port int) *RedisTimeSeriesClient {
	return &RedisTimeSeriesClient{
		host: host,
		port: port,
	}
}

func (client *RedisTimeSeriesClient) Init() {
	connectionString := fmt.Sprintf("%s:%d", client.host, client.port)

	client.redisClient = redistimeseries.NewClient(connectionString, "", nil)
}

func (client *RedisTimeSeriesClient) EmitRedisSignal(flagId, appId int, status string) {
	keyName := fmt.Sprintf("%d:%s", flagId, status)
	labels := map[string]string{
		"status": status,
		"appId":  fmt.Sprintf("%d", appId),
		"flagId": fmt.Sprintf("%d", flagId),
	}
	options := redistimeseries.DefaultCreateOptions
	options.Labels = labels
	client.redisClient.AddAutoTsWithOptions(keyName, 1, options)
}
