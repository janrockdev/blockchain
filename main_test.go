package main

import (
	"testing"
	"time"

	"github.com/janrockdev/blockchain/network"
)

func TestMainFunction(t *testing.T) {
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
	go s.Start()

	// Allow some time for the server to start and process messages
	time.Sleep(3 * time.Second)

	// Check if the transports are connected
	// if !trLocal.IsConnectedTo("REMOTE") || !trRemote.IsConnectedTo("LOCAL") {
	// 	t.Errorf("Transports are not connected")
	// }
}
