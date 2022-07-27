package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.Token("myToken"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("should be connected now")

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	//Simple Async Ephemeral Consumer
	js.Subscribe("apps.1.update.manual", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})

	// js.Subscribe(">", func(m *nats.Msg) {
	// 	fmt.Printf("Received a JetStream message: %s\n", m.Subject)
	// })

	//this is necessary because otherwise main may finish running before the consumer has a chance to
	runtime.Goexit()

}
