package engine

import (
	"fmt"

	"github.com/janrockdev/blockchain/utils"
)

type Transaction struct {
	Data []byte

	PublicKey utils.PublicKey
	Signature *utils.Signature
}

func (tx *Transaction) Sign(privKey utils.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PublicKey = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signed")
	}

	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("transaction signature is invalid")
	}

	return nil
}
