package network

import (
	"fmt"
	"sync"
)

type TransportLocal struct {
	addr           NetAddr
	consumeChannel chan RPC
	lock           sync.RWMutex
	peers          map[NetAddr]*TransportLocal
}

func NewTransportLocal(addr NetAddr) Transport {
	return &TransportLocal{
		addr:           addr,
		consumeChannel: make(chan RPC, 1024),
		peers:          make(map[NetAddr]*TransportLocal),
	}
}

func (t *TransportLocal) Consumer() <-chan RPC {
	return t.consumeChannel
}

func (t *TransportLocal) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*TransportLocal)
	return nil
}

func (t *TransportLocal) SendMessage(addr NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[addr] // to address
	if !ok {
		return fmt.Errorf("%s: peer not found / cound not send message to %s", t.addr, addr)
	}

	peer.consumeChannel <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

func (t *TransportLocal) Addr() NetAddr {
	return t.addr
}
