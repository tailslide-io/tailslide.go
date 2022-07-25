package main

import (
	"fmt"
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
	flagManager "github.com/tailslide-io/tailslide/lib/flagmanager"
)

func main(){
	config := flagManager.FlagManagerConfig{
		NatsServer : "localhost:4222",
		Stream : "flags",
		AppId : "1",
		SdkKey : "myToken",
		UserContext : "375d39e6-9c3f-4f58-80bd-e5960b710295",
		RedisHost : "null",
		RedisPort : 6379,

	}
	manager := flagManager.New(config)
	manager.InitializeFlags()


	fmt.Println("I am past getting last message")

	runtime.Goexit()
}

func printMessage(msg *nats.Msg) {
	log.Println(string(msg.Data))
} 