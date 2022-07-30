package flagManager

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/nats.go"
	natsClient "github.com/tailslide-io/tailslide/lib/natsclient"
	redisClient "github.com/tailslide-io/tailslide/lib/redisclient"
	toggler "github.com/tailslide-io/tailslide/lib/toggler"
	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type FlagManager struct {
	natsStream     string
	natsServer     string
	natsClient     *natsClient.NatsClient
	redistTSClient *redisClient.RedisTimeSeriesClient
	flags          []tailslideTypes.Flag
	userContext    string
}

func New(config tailslideTypes.FlagManagerConfig) *FlagManager {
	return &FlagManager{
		natsClient:     natsClient.New(config.NatsServer, config.NatsStream, fmt.Sprintf("%d", config.AppId), config.SdkKey, nil),
		redistTSClient: redisClient.New(config.RedisHost, config.RedisPort),
		userContext:    config.UserContext,
	}
}

func (manager *FlagManager) InitializeFlags() {
	manager.natsClient.Callback = manager.SetFlags
	manager.natsClient.InitializeFlags()
	manager.redistTSClient.Init()
}

func (manager *FlagManager) SetFlags(message *nats.Msg) {
	var flags []tailslideTypes.Flag
	json.Unmarshal(message.Data, &flags)
	manager.flags = flags
	fmt.Println(flags)
}

func (manager *FlagManager) GetFlags() []tailslideTypes.Flag {
	return manager.flags
}

func (manager *FlagManager) Disconnect() {
	manager.natsClient.Disconnect()
	manager.redistTSClient.Disconnect()
}

func (manager *FlagManager) NewToggler(config toggler.TogglerConfig) (*toggler.Toggler, error) {
	config.GetFlags = manager.GetFlags
	config.UserContext = manager.userContext
	config.EmitRedisSignal = manager.redistTSClient.EmitRedisSignal
	return toggler.New(config)
}
