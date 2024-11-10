package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/janrockdev/blockchain/types"
)

type PrivateKey struct {
	pk *ecdsa.PrivateKey
}

func (pk PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, pk.pk, data)
	if err != nil {
		return nil, err
	}

	return &Signature{r: r, s: s}, nil
}

type PublicKey struct {
	pub *ecdsa.PublicKey
}

func GeneratePrivateKey() PrivateKey {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{pk: pk}
}

func (pk PrivateKey) GeneratePublicKey() PublicKey {
	return PublicKey{pub: &pk.pk.PublicKey}
}

func (pk PublicKey) ToSlice() []byte {
	return elliptic.Marshal(pk.pub.Curve, pk.pub.X, pk.pub.Y)
}

func (pub PublicKey) Address() types.Address {
	hash := sha256.Sum256(pub.ToSlice())

	return types.AddressFromBytes(hash[len(hash)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (sig Signature) Verify(data []byte, pubKey PublicKey) bool {
	return ecdsa.Verify(pubKey.pub, data, sig.r, sig.s)
}
