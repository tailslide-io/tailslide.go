package main

import "github.com/nats-io/nats.go"

func main() {
	// Connect to NATS
	nc, _ := nats.Connect("nats://127.0.0.1:4222")

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// Simple Stream Publisher
	js.Publish("ORDERS.scratch", []byte("hello"))
}
