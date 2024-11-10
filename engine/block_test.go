package engine

import (
	"bytes"
	"testing"
	"time"

	"github.com/janrockdev/blockchain/types"
)

func TestBlockHash(t *testing.T) {
	header := &Header{
		Version:   1,
		PrevBlock: types.Hash{0x01, 0x02, 0x03, 0x04},
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     9999,
	}

	block := &Block{
		Header:       *header,
		Transactions: []Transaction{}, // empty transactions
	}

	hash := block.Hash()

	if hash.IsZero() {
		t.Errorf("Hash is zero")
	}
}

func TestHeader_EncodeDecodeBinary(t *testing.T) {
	originalHeader := &Header{
		Version:   1,
		PrevBlock: types.Hash{0x01, 0x02, 0x03, 0x04},
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     9999,
	}

	var buf bytes.Buffer
	if err := originalHeader.EncodeBinary(&buf); err != nil {
		t.Fatalf("EncodeBinary failed: %v", err)
	}

	decodedHeader := &Header{}
	if err := decodedHeader.DecodeBinary(&buf); err != nil {
		t.Fatalf("DecodeBinary failed: %v", err)
	}

	if originalHeader.Version != decodedHeader.Version {
		t.Errorf("Version mismatch: got %v, want %v", decodedHeader.Version, originalHeader.Version)
	}
	if originalHeader.PrevBlock != decodedHeader.PrevBlock {
		t.Errorf("PrevBlock mismatch: got %v, want %v", decodedHeader.PrevBlock, originalHeader.PrevBlock)
	}
	if originalHeader.Timestamp != decodedHeader.Timestamp {
		t.Errorf("Timestamp mismatch: got %v, want %v", decodedHeader.Timestamp, originalHeader.Timestamp)
	}
	if originalHeader.Height != decodedHeader.Height {
		t.Errorf("Height mismatch: got %v, want %v", decodedHeader.Height, originalHeader.Height)
	}
	if originalHeader.Nonce != decodedHeader.Nonce {
		t.Errorf("Nonce mismatch: got %v, want %v", decodedHeader.Nonce, originalHeader.Nonce)
	}
}
