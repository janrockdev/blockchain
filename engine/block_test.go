package engine

import (
	"testing"
	"time"

	"github.com/janrockdev/blockchain/types"
	"github.com/janrockdev/blockchain/utils"
	"github.com/stretchr/testify/assert"
)

type mockHasher struct{}

// Hash returns a fixed hash
func (m *mockHasher) Hash(b *Block) types.Hash {
	return types.Hash{0x01, 0x02, 0x03, 0x04}
}

// TestBlockHash tests the Hash method of the Block struct
func TestBlockHash(t *testing.T) {
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

// randomBlock creates a random block for testing
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

	// Sign the transaction
	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	tx := Transaction{
		Data: []byte("hello"),
	}

	// Sign the transaction
	block := NewBlock(header, []Transaction{tx})
	block.Sign(utils.GeneratePrivateKey())

	return block
}

// TestSignBlock tests the Sign method of the Block struct
func TestVerifyBlock(t *testing.T) {
	// Test with valid signature
	privKey := utils.GeneratePrivateKey()
	block := randomBlock(0)
	assert.Nil(t, block.Sign(privKey))
	assert.Nil(t, block.Verify())

	// Test with no signature
	otherPrivKey := utils.GeneratePrivateKey()
	block.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, block.Verify())

	// Test with tampered data
	block.Height = 100
	assert.NotNil(t, block.Verify())
}
