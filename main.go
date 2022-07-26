package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	fmt.Println("should be connected now")

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// js.AddStream(&nats.StreamConfig{
	// 	Name:     "ORDERS",
	// 	Subjects: []string{"ORDERS.*"},
	// })

	// Simple Stream Publisher
	js.Publish("ORDERS.scratch", []byte("hello"))

	//Simple Async Ephemeral Consumer
	js.Subscribe("ORDERS.>", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})

	//this is necessary because otherwise main may finish running before the consumer has a chance to
	runtime.Goexit()

}
