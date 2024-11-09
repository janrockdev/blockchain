package network

import (
	"testing"
)

func TestTransportLocal_Consumer(t *testing.T) {
	addr := NetAddr("localhost:8080")
	transport := NewTransportLocal(addr)

	consumerChan := transport.Consumer()
	if consumerChan == nil {
		t.Fatalf("expected non-nil channel, got nil")
	}

	expectedCapacity := 1024
	actualCapacity := cap(consumerChan)
	if actualCapacity != expectedCapacity {
		t.Fatalf("expected channel capacity %d, got %d", expectedCapacity, actualCapacity)
	}
}
