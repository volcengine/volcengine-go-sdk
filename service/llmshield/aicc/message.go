package jeddak_secure_channel

import (
	"encoding/json"
	"io"
)

// EncryptedMessage represents an encrypted message containing nonce, mac, key, and ciphertext.
// All byte fields are serialized as Base64 strings in JSON.
type EncryptedMessage struct {
	Nonce      []byte `json:"nonce"`
	Mac        []byte `json:"mac"`
	Key        []byte `json:"key"`        // Can be nil
	Ciphertext []byte `json:"ciphertext"` // Can be nil
}

// Deserialize reads JSON from a reader and parses it into an EncryptedMessage.
// It automatically decodes Base64 strings from JSON into byte slices.
func deserializeEncryptedMessage(r io.Reader) (*EncryptedMessage, error) {
	var msg EncryptedMessage
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// Serialize converts the EncryptedMessage to a JSON string.
// Byte slices are automatically encoded as Base64 strings in JSON.
func (e *EncryptedMessage) Serialize() (string, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
