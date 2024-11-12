package utils

import (
	"testing"

	"github.com/janrockdev/blockchain/types"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	if privKey.pk == nil {
		t.Errorf("Expected private key to be generated, got nil")
	}
}

func TestGeneratePublicKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GeneratePublicKey()
	if pubKey.pub == nil {
		t.Errorf("Expected public key to be generated, got nil")
	}
}

func TestPublicKeyToSlice(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GeneratePublicKey()
	pubKeyBytes := pubKey.ToSlice()
	if len(pubKeyBytes) == 0 {
		t.Errorf("Expected public key bytes to be non-empty, got empty slice")
	}
}

func TestPublicKeyAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GeneratePublicKey()
	address := pubKey.Address()
	if len(address) != types.AddressLength {
		t.Errorf("Expected address length to be %d, got %d", types.AddressLength, len(address))
	}
	t.Logf("Generated address: %x", address)
}
func TestSign(t *testing.T) {
	privKey := GeneratePrivateKey()
	data := []byte("test data")
	signature, err := privKey.Sign(data)
	if err != nil {
		t.Fatalf("Expected no error signing data, got %v", err)
	}
	if signature == nil {
		t.Fatalf("Expected signature to be generated, got nil")
	}
	if signature.r == nil || signature.s == nil {
		t.Errorf("Expected signature components to be non-nil, got r: %v, s: %v", signature.r, signature.s)
	}
}
func TestVerify(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GeneratePublicKey()
	data := []byte("test data")
	signature, err := privKey.Sign(data)
	if err != nil {
		t.Fatalf("Expected no error signing data, got %v", err)
	}

	valid := signature.Verify(pubKey, data)
	if !valid {
		t.Errorf("Expected signature to be valid, got invalid")
	}

	// Test with tampered data
	tamperedData := []byte("tampered data")
	invalid := signature.Verify(pubKey, tamperedData)
	if invalid {
		t.Errorf("Expected signature to be invalid with tampered data, got valid")
	}
}
