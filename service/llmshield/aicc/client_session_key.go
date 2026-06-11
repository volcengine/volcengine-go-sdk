package jeddak_secure_channel

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

// ClientSessionKey manages session encryption using a server's RSA public key.
type ClientSessionKey struct {
	ServerPublicKey *rsa.PublicKey // RSA public key for encrypting AES keys
}

// LoadClientSessionKey parses an RSA public key from a PEM-encoded string.
// The PEM must contain a "PUBLIC KEY" block.
func LoadClientSessionKey(serverKeyPem string) (*ClientSessionKey, error) {
	// Decode PEM block
	block, _ := pem.Decode([]byte(serverKeyPem))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid PEM public key (expected 'PUBLIC KEY' block)")
	}

	// Parse DER-encoded public key
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Verify it's an RSA public key
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not an RSA key")
	}

	return &ClientSessionKey{ServerPublicKey: rsaPubKey}, nil
}

// EncryptWithResponse encrypts plaintext and returns an EncryptResult containing:
// - Serialized encrypted message (AES ciphertext + RSA-encrypted AES key)
// - AES key for decrypting server responses
func (c *ClientSessionKey) EncryptWithResponse(plaintext []byte) (*EncryptResult, error) {
	aesKey, err := GenerateAesKey()
	if err != nil {
		return nil, err
	}

	aesResult, err := aesKey.Encrypt(plaintext)
	if err != nil {
		return nil, err
	}

	return c.toEncryptResult(aesKey, aesResult)
}

// EncryptStreamWithResponse encrypts a stream and returns an EncryptResult.
func (c *ClientSessionKey) EncryptStreamWithResponse(source io.Reader, sink io.Writer) (*EncryptResult, error) {
	aesKey, err := GenerateAesKey()
	if err != nil {
		return nil, err
	}

	aesResult, err := aesKey.EncryptStream(source, sink, false)
	if err != nil {
		return nil, err
	}

	return c.toEncryptResult(aesKey, aesResult)
}

// toEncryptResult combines AES encryption results with RSA-encrypted AES key.
func (c *ClientSessionKey) toEncryptResult(aesKey *AesKey, aesResult *AesEncryptResult) (*EncryptResult, error) {
	// RSA-OAEP encrypt the AES key with the server's public key
	encryptedKey, err := rsa.EncryptOAEP(
		sha256.New(),      // Hash function
		rand.Reader,       // Random source
		c.ServerPublicKey, // Server's RSA public key
		aesKey.key,        // Data to encrypt (AES key)
		nil,               // Optional label (not used)
	)
	if err != nil {
		return nil, fmt.Errorf("failed to RSA-OAEP encrypt AES key: %w", err)
	}

	// Build EncryptedMessage with base64-encoded fields
	em := EncryptedMessage{
		Nonce:      aesResult.Nonce,
		Mac:        aesResult.Mac,
		Key:        encryptedKey,
		Ciphertext: aesResult.Ciphertext,
	}

	// Serialize EncryptedMessage to JSON string
	serializedEM, err := json.Marshal(em)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize EncryptedMessage: %w", err)
	}

	return &EncryptResult{
		Ciphertext:  string(serializedEM),
		ResponseKey: ResponseKey{key: aesKey, usage: UsageDecrypt},
	}, nil
}
