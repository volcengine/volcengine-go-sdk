package jeddak_secure_channel

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	KEY_LEN   = 32
	NONCE_LEN = 12
	MAC_LEN   = 16
)

// AesKey holds an AES-256 key and a nonce source for encryption/decryption.
type AesKey struct {
	key []byte
}

// GenerateAesKey generates a new AES-256 key.
func GenerateAesKey() (*AesKey, error) {
	key := make([]byte, KEY_LEN)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return &AesKey{
		key: key,
	}, nil
}

// AesKeyFromBytes creates an AesKey from a 32-byte key.
func AesKeyFromBytes(key []byte) (*AesKey, error) {
	if len(key) != KEY_LEN {
		return nil, invalidKeyError(len(key))
	}
	return &AesKey{
		key: key,
	}, nil
}

// GetEncoded returns the raw key bytes.
func (a *AesKey) GetEncoded() []byte {
	return a.key
}

// EncryptResult holds the output of encryption: nonce, ciphertext, and MAC.
type AesEncryptResult struct {
	Nonce      []byte
	Ciphertext []byte // nil if generated via EncryptStream with concatted=false
	Mac        []byte
}

// Encrypt encrypts plaintext and returns the nonce, ciphertext, and MAC.
func (a *AesKey) Encrypt(plaintext []byte) (*AesEncryptResult, error) {
	nonce := make([]byte, NONCE_LEN)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	aead, err := aesgcmNew(a.key)
	if err != nil {
		return nil, err
	}

	sealed := aead.Seal(nil, nonce, plaintext, nil)
	if len(sealed) < MAC_LEN {
		return nil, errors.New("seal result too short")
	}
	ciphertext := sealed[:len(sealed)-MAC_LEN]
	mac := sealed[len(sealed)-MAC_LEN:]

	return &AesEncryptResult{
		Nonce:      nonce,
		Ciphertext: ciphertext,
		Mac:        mac,
	}, nil
}

// EncryptStream encrypts data from source to sink, returning nonce and MAC.
// If concatted=true, nonce and (ciphertext+MAC) are written to sink.
func (a *AesKey) EncryptStream(source io.Reader, sink io.Writer, concatted bool) (*AesEncryptResult, error) {
	nonce := make([]byte, NONCE_LEN)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	if concatted {
		if _, err := sink.Write(nonce); err != nil {
			return nil, err
		}
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	// Compute hash subkey H
	var zeroBlock [16]byte
	h := make([]byte, 16)
	block.Encrypt(h, zeroBlock[:])

	ghash := newGHASH(h)
	initialCounter := make([]byte, 16)
	copy(initialCounter, nonce)
	binary.BigEndian.PutUint32(initialCounter[12:], 1) // CTR starts at 1
	stream := cipher.NewCTR(block, initialCounter)

	buf := make([]byte, 1024)
	totalCiphertextLen := 0

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}

		ciphertext := make([]byte, n)
		stream.XORKeyStream(ciphertext, buf[:n])

		if _, err := sink.Write(ciphertext); err != nil {
			return nil, err
		}
		ghash.update(ciphertext)
		totalCiphertextLen += n
	}

	// Compute tag
	tag := ghash.finalize(0, totalCiphertextLen)
	var initialCounter0 [16]byte
	copy(initialCounter0[:], nonce)
	encryptedInitialCounter0 := make([]byte, 16)
	block.Encrypt(encryptedInitialCounter0, initialCounter0[:])
	xorBytes(tag, tag, encryptedInitialCounter0)

	if concatted {
		if _, err := sink.Write(tag); err != nil {
			return nil, err
		}
	}

	return &AesEncryptResult{
		Nonce:      nonce,
		Ciphertext: nil,
		Mac:        tag,
	}, nil
}

// Decrypt decrypts ciphertext using nonce and MAC, returning plaintext.
func (a *AesKey) Decrypt(nonce, ciphertext, mac []byte) ([]byte, error) {
	aead, err := aesgcmNew(a.key)
	if err != nil {
		return nil, err
	}
	ciphertextWithTag := append(append([]byte(nil), ciphertext...), mac...)
	return aead.Open(nil, nonce, ciphertextWithTag, nil)
}

// DecryptStream decrypts ciphertext from source to sink using nonce and MAC.
func (a *AesKey) DecryptStream(nonce, mac []byte, source io.Reader, sink io.Writer) error {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return err
	}

	var zeroBlock [16]byte
	h := make([]byte, 16)
	block.Encrypt(h, zeroBlock[:])

	ghash := newGHASH(h)
	initialCounter := make([]byte, 16)
	copy(initialCounter, nonce)
	binary.BigEndian.PutUint32(initialCounter[12:], 1)
	stream := cipher.NewCTR(block, initialCounter)

	buf := make([]byte, 1024)
	totalCiphertextLen := 0

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		plaintext := make([]byte, n)
		stream.XORKeyStream(plaintext, buf[:n])

		if _, err := sink.Write(plaintext); err != nil {
			return err
		}
		ghash.update(buf[:n])
		totalCiphertextLen += n
	}

	// Verify tag (mac)
	computedTag := ghash.finalize(0, totalCiphertextLen)
	var initialCounter0 [16]byte
	copy(initialCounter0[:], nonce)
	encryptedInitialCounter0 := make([]byte, 16)
	block.Encrypt(encryptedInitialCounter0, initialCounter0[:])
	xorBytes(computedTag, computedTag, encryptedInitialCounter0)

	if !bytes.Equal(computedTag, mac) {
		return invalidMacError()
	}
	return nil
}

// Helper functions and types for GHASH (AES-GCM authentication)
type ghash struct {
	h [16]byte
	y [16]byte
}

func newGHASH(h []byte) *ghash {
	var hArr [16]byte
	copy(hArr[:], h)
	return &ghash{h: hArr}
}

func (g *ghash) update(data []byte) {
	for len(data) > 0 {
		var block [16]byte
		copy(block[:], data)
		xorBytes(g.y[:], g.y[:], block[:])
		g.multiplyH()
		if len(data) < 16 {
			break
		}
		data = data[16:]
	}
}

func (g *ghash) finalize(aadLen, dataLen int) []byte {
	var lenBlock [16]byte
	binary.BigEndian.PutUint64(lenBlock[:8], uint64(aadLen)*8)
	binary.BigEndian.PutUint64(lenBlock[8:], uint64(dataLen)*8)
	xorBytes(g.y[:], g.y[:], lenBlock[:])
	g.multiplyH()
	return g.y[:]
}

func (g *ghash) multiplyH() {
	var z [16]byte
	for i := 0; i < 16; i++ {
		for j := 0; j < 8; j++ {
			if (g.y[i]>>uint(7-j))&1 == 1 {
				xorBytes(z[:], z[:], g.h[:])
			}
			carry := (g.h[0] >> 7) & 1
			for k := 0; k < 15; k++ {
				g.h[k] = (g.h[k] << 1) | (g.h[k+1] >> 7)
			}
			g.h[15] = (g.h[15] << 1) ^ carry*0x87
		}
	}
	g.y = z
}

// Utility functions
func aesgcmNew(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

func xorBytes(dst, a, b []byte) {
	for i := range dst {
		dst[i] = a[i] ^ b[i]
	}
}

func invalidKeyError(length int) error {
	return fmt.Errorf("invalid key length %d, expected %d", length, KEY_LEN)
}

func invalidMacError() error {
	return fmt.Errorf("invalid MAC")
}
