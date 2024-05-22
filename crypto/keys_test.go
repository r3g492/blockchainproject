package crypto

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	assert.Equal(t, len(privateKey.Bytes()), privateKeyLen)

	publicKey := privateKey.PublicKey()
	assert.Equal(t, len(publicKey.Bytes()), publicKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	msg := []byte("hello")
	signature := privateKey.Sign(msg)

	// Test with the correct message
	assert.True(t, signature.Verify(publicKey, msg))

	// Test with the wrong message
	assert.False(t, signature.Verify(publicKey, []byte("world")))

	// Test with the wrong public key
	pk := GeneratePrivateKey()
	assert.False(t, signature.Verify(pk.PublicKey(), msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	address := publicKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))

	fmt.Println("Public Key:", hex.EncodeToString(publicKey.Bytes()))
	fmt.Println("Address:", hex.EncodeToString(address.Bytes()))
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "8f8ce97f48d39fc81120cba56a341156cb98e83d55f1b9bbb1bc09977e6d062e"
		privateKey = NewPrivateKeyFromString(seed)
		addressStr = "08b32c049a297584c1bdc35afc08b3b8a7bcc220"
	)

	assert.Equal(t, privateKeyLen, len(privateKey.Bytes()))
	address := privateKey.PublicKey().Address()
	fmt.Println("Address:", address.String())
	assert.Equal(t, addressStr, address.String())
}
