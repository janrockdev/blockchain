package engine

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/janrockdev/blockchain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(b.Header); err != nil {
		panic(err)
	}
	h := sha256.Sum256(buf.Bytes())
	return types.Hash(h)
}

// type TxHasher struct{}

// // Hash will hash the whole bytes of the TX no exception.
// func (TxHasher) Hash(tx *Transaction) types.Hash {
// 	buf := new(bytes.Buffer)

// 	binary.Write(buf, binary.LittleEndian, tx.Data)
// 	binary.Write(buf, binary.LittleEndian, tx.To)
// 	binary.Write(buf, binary.LittleEndian, tx.Value)
// 	binary.Write(buf, binary.LittleEndian, tx.From)
// 	binary.Write(buf, binary.LittleEndian, tx.Nonce)

// 	return types.Hash(sha256.Sum256(buf.Bytes()))
// }
