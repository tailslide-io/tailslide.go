package flagManager

import (
	"encoding/json"

	nats "github.com/nats-io/nats.go"
	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type FlagManager struct {
	flags []tailslideTypes.Flag

}

func (manager *FlagManager)SetFlags(message *nats.Msg){
	var flags []tailslideTypes.Flag
	json.Unmarshal(message.Data, &flags)
	manager.flags = flags
}

