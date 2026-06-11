package jeddak_secure_channel

import (
	"errors"
	"io"
	"os"
	"strings"
)

// Usage specifies the intended usage of the ResponseKey (encrypt or decrypt)
type Usage int

const (
	// UsageEncrypt indicates the key is used for encryption (server side)
	UsageEncrypt Usage = iota
	// UsageDecrypt indicates the key is used for decryption (client side)
	UsageDecrypt
)

// ResponseKey wraps an AES key with specific usage constraints
type ResponseKey struct {
	key   *AesKey
	usage Usage
}

// EncryptString encrypts a string response (server use)
func (r *ResponseKey) EncryptString(response string) (string, error) {
	return r.EncryptBytes([]byte(response))
}

// EncryptBytes encrypts a bytes response (server use)
func (r *ResponseKey) EncryptBytes(response []byte) (string, error) {
	if r.usage != UsageEncrypt {
		return "", errors.New("unsupported operation: key usage is not Encrypt")
	}

	result, err := r.key.Encrypt(response)
	if err != nil {
		return "", err
	}

	message := &EncryptedMessage{
		Nonce:      result.Nonce,
		Mac:        result.Mac,
		Key:        nil,
		Ciphertext: result.Ciphertext,
	}
	return message.Serialize()
}

// EncryptStream encrypts data from source to sink (server use)
func (r *ResponseKey) EncryptStream(source io.Reader, sink io.Writer) (string, error) {
	if r.usage != UsageEncrypt {
		return "", errors.New("unsupported operation: key usage is not Encrypt")
	}

	result, err := r.key.EncryptStream(source, sink, false)
	if err != nil {
		return "", err
	}

	message := &EncryptedMessage{
		Nonce:      result.Nonce,
		Mac:        result.Mac,
		Key:        nil,
		Ciphertext: nil,
	}
	return message.Serialize()
}

// EncryptFiles encrypts input file to output file (server use)
func (r *ResponseKey) EncryptFiles(input, output string) (string, error) {
	source, err := os.Open(input)
	if err != nil {
		return "", err
	}
	defer source.Close()

	sink, err := os.Create(output)
	if err != nil {
		return "", err
	}
	defer sink.Close()

	return r.EncryptStream(source, sink)
}

// DecryptBytesString decrypts a serialized response string to bytes (client use)
func (r *ResponseKey) DecryptBytesString(response string) ([]byte, error) {
	return r.DecryptBytesReader(strings.NewReader(response))
}

// DecryptBytesReader decrypts a serialized response from reader to bytes (client use)
func (r *ResponseKey) DecryptBytesReader(response io.Reader) ([]byte, error) {
	if r.usage != UsageDecrypt {
		return nil, errors.New("unsupported operation: key usage is not Decrypt")
	}

	message, err := deserializeEncryptedMessage(response)
	if err != nil {
		return nil, err
	}

	if message.Ciphertext == nil {
		return nil, errors.New("missing ciphertext in encrypted message")
	}

	return r.key.Decrypt(message.Nonce, message.Ciphertext, message.Mac)
}

// DecryptString decrypts a serialized response string to UTF-8 string (client use)
func (r *ResponseKey) DecryptString(response string) (string, error) {
	bytes, err := r.DecryptBytesString(response)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DecryptStringReader decrypts a serialized response from reader to UTF-8 string (client use)
func (r *ResponseKey) DecryptStringReader(response io.Reader) (string, error) {
	bytes, err := r.DecryptBytesReader(response)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DecryptStream decrypts data from source to sink using serialized response (client use)
func (r *ResponseKey) DecryptStream(response io.Reader, source io.Reader, sink io.Writer) error {
	if r.usage != UsageDecrypt {
		return errors.New("unsupported operation: key usage is not Decrypt")
	}

	message, err := deserializeEncryptedMessage(response)
	if err != nil {
		return err
	}

	if message.Ciphertext != nil {
		return errors.New("unexpected ciphertext in stream message")
	}

	return r.key.DecryptStream(message.Nonce, message.Mac, source, sink)
}

// DecryptStreamString decrypts data from source to sink using response string (client use)
func (r *ResponseKey) DecryptStreamString(response string, source io.Reader, sink io.Writer) error {
	return r.DecryptStream(strings.NewReader(response), source, sink)
}

// DecryptFiles decrypts input file to output file using response string (client use)
func (r *ResponseKey) DecryptFiles(response, input, output string) error {
	source, err := os.Open(input)
	if err != nil {
		return err
	}
	defer source.Close()

	sink, err := os.Create(output)
	if err != nil {
		return err
	}
	defer sink.Close()

	return r.DecryptStreamString(response, source, sink)
}

// Usage returns the key's intended usage
func (r *ResponseKey) Usage() Usage {
	return r.usage
}

// The following types/functions are assumed to be implemented in other files:
// type AesKey struct { ... }
// type NonceSource interface { ... }
// type EncryptedMessage struct { ... }
// func NewUserProvidedNonceSource(nonce []byte) NonceSource
// func NewRandomNonceSource() NonceSource
// func (a *AesKey) Encoded() []byte
// func (a *AesKey) NonceSource() NonceSource
// func (a *AesKey) SetNonceSource(NonceSource)
// func (a *AesKey) Encrypt([]byte) (EncryptResult, error)
// func (a *AesKey) EncryptStream(io.Reader, io.Writer) (EncryptResult, error)
// func (a *AesKey) Decrypt(nonce, ciphertext, mac []byte) ([]byte, error)
// func (a *AesKey) DecryptStream(nonce, mac []byte, source io.Reader, sink io.Writer) error
// type EncryptResult struct { Nonce, Mac, Ciphertext []byte }
// func (e *EncryptedMessage) Serialize() (string, error)
// func EncryptedMessageDeserialize(io.Reader) (*EncryptedMessage, error)
