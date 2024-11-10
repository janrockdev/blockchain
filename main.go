package main

import (
	"time"

	"github.com/janrockdev/blockchain/network"
)

func main() {
	trLocal := network.NewTransportLocal("LOCAL")
	trRemote := network.NewTransportLocal("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage("LOCAL", []byte("hello"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
