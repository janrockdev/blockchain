package engine

import (
	"testing"

	"github.com/janrockdev/blockchain/utils"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {

	privKey := utils.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("hello"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}
