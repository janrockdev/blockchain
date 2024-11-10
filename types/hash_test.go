package types

import (
	"encoding/hex"
	"testing"
)

func TestHash_IsZero(t *testing.T) {
	var zeroHash Hash
	if !zeroHash.IsZero() {
		t.Errorf("Expected zero hash to be zero")
	}

	nonZeroHash := RandomHash()
	if nonZeroHash.IsZero() {
		t.Errorf("Expected non-zero hash to be non-zero")
	}
}

func TestHash_ToSlice(t *testing.T) {
	h := RandomHash()
	slice := h.ToSlice()
	if len(slice) != 32 {
		t.Errorf("Expected slice length to be 32, got %d", len(slice))
	}
	for i := 0; i < 32; i++ {
		if slice[i] != h[i] {
			t.Errorf("Expected slice[%d] to be %d, got %d", i, h[i], slice[i])
		}
	}
}

func TestHash_String(t *testing.T) {
	h := RandomHash()
	str := h.String()
	expectedStr := hex.EncodeToString(h.ToSlice())
	if str != expectedStr {
		t.Errorf("Expected string %s, got %s", expectedStr, str)
	}
}

func TestHashFromBytes(t *testing.T) {
	bytes := RandomBytes(32)
	h := HashFromBytes(bytes)
	for i := 0; i < 32; i++ {
		if h[i] != bytes[i] {
			t.Errorf("Expected hash[%d] to be %d, got %d", i, bytes[i], h[i])
		}
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid byte length")
		}
	}()
	_ = HashFromBytes(RandomBytes(31))
}

func TestRandomBytes(t *testing.T) {
	bytes := RandomBytes(32)
	if len(bytes) != 32 {
		t.Errorf("Expected byte length to be 32, got %d", len(bytes))
	}
}

func TestRandomHash(t *testing.T) {
	h := RandomHash()
	if h.IsZero() {
		t.Errorf("Expected random hash to be non-zero")
	}
}
