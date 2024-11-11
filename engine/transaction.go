package engine

import (
	"io"

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

func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
