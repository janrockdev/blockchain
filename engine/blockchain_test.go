package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesisBlock(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesisBlock(t)
	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(uint32(i + 1))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, uint32(lenBlocks), bc.Height())
	assert.Equal(t, lenBlocks+1, len(bc.headers))
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(89)))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesisBlock(t)
	assert.True(t, bc.HasBlock(0))
}
