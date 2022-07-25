package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	nats "github.com/nats-io/nats.go"
	flagManager "github.com/tailslide-io/tailslide/lib/flagmanager"
	toggler "github.com/tailslide-io/tailslide/lib/toggler"
	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

func main(){
	config := tailslideTypes.FlagManagerConfig{
		NatsServer : "localhost:4222",
		Stream : "flags",
		AppId : 1,
		SdkKey : "myToken",
		UserContext : "375d39e6-9c3f-4f58-80bd-e5960b710295",
		RedisHost : "null",
		RedisPort : 6379,

	}
	manager := flagManager.New(config)
	manager.InitializeFlags()
	
	flagName := "Flag in app 1 number 1"
	flagConfig := toggler.TogglerConfig{
		FlagName: flagName,
	}
	toggler, err := manager.NewToggler(flagConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	
	for {
		if (toggler.IsFlagActive()){
			fmt.Printf(`Flag in {app_id} with name "%s" is active!`, flagName)
		} else {
			fmt.Printf(`Flag in {app_id} with name "%s" is not active!`, flagName)
		}
		time.Sleep(4 * time.Second)
	}


	fmt.Println("I am past getting last message")

	runtime.Goexit()
}

func printMessage(msg *nats.Msg) {
	log.Println(string(msg.Data))
} 