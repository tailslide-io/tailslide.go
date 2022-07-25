package flagManager

import (
	"encoding/json"

	nats "github.com/nats-io/nats.go"
	natsClient "github.com/tailslide-io/tailslide/lib/natsclient"
	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type FlagManager struct {
	natsServer string
	natsClient *natsClient.NatsClient
	// redistTSClient
	flags []tailslideTypes.Flag
	userContext string

}

func NewFlagManager(natsServer, stream, appId, sdkKey, userContext, redisHost string, redistPort int) *FlagManager{
	return &FlagManager{
		natsClient: natsClient.NewNatsClient(natsServer, stream, appId, sdkKey, nil ),
		userContext: userContext,
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
}

func (manager *FlagManager) GetFlags() []tailslideTypes.Flag{
	return manager.flags
}

func (manager *FlagManager) Disconnect(){
	manager.natsClient.Disconnect()
	// manager.redistTSClient.Disconnect()
}

// func (manager *FlagManager) NewToggler(config){}

