package engine

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/janrockdev/blockchain/types"
	"github.com/janrockdev/blockchain/utils"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Height        uint32
	Timestamp     int64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    utils.PublicKey
	Signature    *utils.Signature

	// Cached version of the header hash
	hash types.Hash
}

func NewBlock(header *Header, txs []Transaction) *Block {
	return &Block{
		Header:       header,
		Transactions: txs,
	}
}

func (b *Block) Sign(privKey utils.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block signature is invalid")
	}

	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}
