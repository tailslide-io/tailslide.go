package main

import (
	"runtime"

	"github.com/tailslide-io/tailslide/natsClient"
)

func main() {
	natsClient.ConnectAndListen()
	runtime.Goexit()
}
