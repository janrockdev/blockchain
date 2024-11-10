package engine

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/janrockdev/blockchain/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32
	Nonce     uint64
}

// Generic interface for encoding binary data using the io.Writer interface.
func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

type Block struct {
	Header
	Transactions []Transaction

	// cached version of the header hash
	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buffer := &bytes.Buffer{}
	b.Header.EncodeBinary(buffer)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buffer.Bytes()))
	}

	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint64(len(b.Transactions))); err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}
	var txCount uint64
	if err := binary.Read(r, binary.LittleEndian, &txCount); err != nil {
		return err
	}
	b.Transactions = make([]Transaction, txCount)
	for i := range b.Transactions {
		if err := b.Transactions[i].DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil
}
