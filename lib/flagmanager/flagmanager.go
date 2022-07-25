package flagManager

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/nats.go"
	natsClient "github.com/tailslide-io/tailslide/lib/natsclient"
	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type FlagManagerConfig struct {
	NatsServer string
	Stream string
	SdkKey string
	AppId string
	UserContext string
	RedisHost string
	RedisPort int
}

type FlagManager struct {
	natsServer string
	natsClient *natsClient.NatsClient
	// redistTSClient
	flags []tailslideTypes.Flag
	userContext string

}

func New(config FlagManagerConfig) *FlagManager{
	return &FlagManager{
		natsClient: natsClient.New(config.NatsServer, config.Stream, config.AppId, config.SdkKey, nil ),
		userContext: config.UserContext,
	}
}

func (manager *FlagManager) InitializeFlags(){
	manager.natsClient.Callback = manager.SetFlags
	manager.natsClient.InitializeFlags()
	// manager.redistTSClient.Init()
}

func (manager *FlagManager) SetFlags(message *nats.Msg){
	var flags []tailslideTypes.Flag
	json.Unmarshal(message.Data, &flags)
	manager.flags = flags
	fmt.Println(flags)
}

func (manager *FlagManager) GetFlags() []tailslideTypes.Flag{
	return manager.flags
}

func (manager *FlagManager) Disconnect(){
	manager.natsClient.Disconnect()
	// manager.redistTSClient.Disconnect()
}

// func (manager *FlagManager) NewToggler(config){}

