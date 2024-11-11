package engine

import (
	"fmt"
	"testing"
	"time"

	"github.com/janrockdev/blockchain/types"
	"github.com/janrockdev/blockchain/utils"
	"github.com/stretchr/testify/assert"
)

type mockHasher struct{}

func (m *mockHasher) Hash(b *Block) types.Hash {
	return types.Hash{0x01, 0x02, 0x03, 0x04}
}

func TestBlock_Hash(t *testing.T) {
	header := &Header{
		Version:       1,
		DataHash:      types.Hash{0x00},
		PrevBlockHash: types.Hash{0x00},
		Height:        1,
		Timestamp:     1234567890,
	}

	block := &Block{
		Header:       header,
		Transactions: nil,
		Validator:    utils.PublicKey{},
		Signature:    &utils.Signature{},
		hash:         types.Hash{},
	}

	hasher := &mockHasher{}

	// Test that the hash is calculated correctly
	expectedHash := types.Hash{0x01, 0x02, 0x03, 0x04}
	actualHash := block.Hash(hasher)
	assert.Equal(t, expectedHash, actualHash, "The hash should be correctly calculated")

	// Test that the cached hash is returned
	cachedHash := block.Hash(hasher)
	assert.Equal(t, expectedHash, cachedHash, "The cached hash should be returned")
}

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	tx := Transaction{
		Data: []byte("hello"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestHashBlock(t *testing.T) {
	block := randomBlock(0)
	fmt.Println(block.Hash(BlockHasher{}))
}
