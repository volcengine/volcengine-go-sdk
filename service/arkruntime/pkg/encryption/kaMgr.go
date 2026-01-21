package encryption

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const (
	aesKeySize   = 32
	aesNonceSize = 12
)

type KeyAgreementClient struct {
	cert      *x509.Certificate
	publicKey *ecdsa.PublicKey
}

func MarshalBinary(pub *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(pub.Curve, pub.X, pub.Y)
}

func ECDHMarshalBinary(pub *ecdsa.PublicKey) []byte {
	byteLen := (pub.Curve.Params().BitSize + 7) / 8
	ecdh := make([]byte, byteLen)
	pub.X.FillBytes(ecdh[:])
	return ecdh
}

// GenerateKey is the constructor for Private-Public key pair
func GenerateKey(curve elliptic.Curve) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	rand := rand.Reader
	privateBytes, x, y, err := elliptic.GenerateKey(curve, rand)
	if err != nil {
		return nil, nil, err
	}
	public := ecdsa.PublicKey{X: x, Y: y, Curve: curve}
	private := ecdsa.PrivateKey{PublicKey: public, D: new(big.Int).SetBytes(privateBytes)}
	return &private, &public, nil
}

// NewP256KeyAgreementClient Load cert and extract public key
func NewP256KeyAgreementClient(pemString string) (*KeyAgreementClient, error) {
	var err error
	p256NistKey := &KeyAgreementClient{}
	p256NistKey.cert, err = ReadCertFromString(pemString)
	if err != nil {
		return nil, err
	}
	switch p256NistKey.cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		pub := p256NistKey.cert.PublicKey.(*ecdsa.PublicKey)
		p256NistKey.publicKey = pub
		return p256NistKey, nil
	default:
		return nil, fmt.Errorf("parsed public key is not a ecdsa public key")
	}
}

// GenerateECIESKeyPair generate ECIES key pair and return the tuple (keyNonce, session token, error)
func (k *KeyAgreementClient) GenerateECIESKeyPair() ([]byte, string, error) {
	hash := sha256.New
	// Generate an ephemeral elliptic curve scalar and point
	r, R, err := GenerateKey(k.publicKey.Curve)
	if err != nil {
		return nil, "", err
	}
	// Compute shared DH key
	dh := &ecdsa.PublicKey{Curve: k.publicKey.Curve}
	dh.X, dh.Y = k.publicKey.Curve.ScalarMult(k.publicKey.X, k.publicKey.Y, r.D.Bytes())
	// Derive symmetric key and nonce via HKDF
	len := aesKeySize + aesNonceSize
	buf, err := deriveKeyBasic(hash, dh, len)
	if err != nil {
		return nil, "", err
	}
	token := MarshalBinary(R)
	return buf, base64.StdEncoding.EncodeToString(token), nil
}

// AesGcmEncryptBytes encrypt message using AES-GCM
func AesGcmEncryptBytes(key, nonce, plain []byte) ([]byte, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	aesgcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm from aes: %w", err)
	}
	return aesgcm.Seal(nil, nonce, plain, nil), nil
}

// AesGcmEncryptBase64String Encrypt message from base64 string to string using AES-GCM
func AesGcmEncryptBase64String(key, nonce []byte, plaintext string) (string, error) {
	plainBytes := []byte(plaintext)
	c, err := AesGcmEncryptBytes(key, nonce, plainBytes)
	return base64.StdEncoding.EncodeToString(c), err
}

// AesGcmDecryptBytes Decrypt message using AES-GCM
func AesGcmDecryptBytes(key, nonce, cipherBytes []byte) ([]byte, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	aesgcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm from aes: %w", err)
	}
	plainBytes, err := aesgcm.Open(nil, nonce, cipherBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt message: %w", err)
	}
	return plainBytes, nil
}

// AesGcmDecryptBase64String decrypt message(base64.std.string) using AES-GCM
func AesGcmDecryptBase64String(key, nonce []byte, ciphertext string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("provided ciphered text is not a valid base64 string: %w", err)
	}
	plainBytes, err := AesGcmDecryptBytes(key, nonce, cipherBytes)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

func ReadEcdsaPrivKeyFromBytes(pemBytes []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("no PEM data is found")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	return key, nil
}

func ReadEcdsaPrivKeyFromString(key string) (*ecdsa.PrivateKey, error) {
	return ReadEcdsaPrivKeyFromBytes([]byte(key))
}

func ReadCertFromBytes(pemBytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("no PEM data is found")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}

	return cert, nil
}

func ReadCertFromString(pemString string) (*x509.Certificate, error) {
	return ReadCertFromBytes([]byte(pemString))
}

func GetCertInfo(certPem string) (string, string, int64) {
	cert, err := ReadCertFromString(certPem)
	if err != nil {
		return "", "", 0
	}
	dns := cert.DNSNames
	if len(dns) > 1 && strings.HasPrefix(dns[0], "ring.") && strings.HasPrefix(dns[1], "key.") {
		return dns[0][5:], dns[1][4:], cert.NotAfter.UTC().Unix()
	}
	return "", "", cert.NotAfter.UTC().Unix()
}

// hkdfReader implements the io.Reader interface for HKDF expansion
// This is a simplified implementation compatible with the standard library
// Based on RFC 5869 (HKDF: HMAC-based Extract-and-Expand Key Derivation Function)
type hkdfReader struct {
	mac  hash.Hash
	hmac []byte
	info []byte
	i    uint8
	temp []byte
	out  []byte
}

// NewHKDF creates a new HKDF reader for extracting and expanding keys
func NewHKDF(hash func() hash.Hash, secret, salt, info []byte) io.Reader {
	// Extract step
	if salt == nil {
		// Use an all-zero salt if none is provided
		salt = make([]byte, hash().Size())
	}
	extract := hmac.New(hash, salt)
	extract.Write(secret)
	prk := extract.Sum(nil)

	// Expand step preparation
	return &hkdfReader{
		mac:  hmac.New(hash, prk),
		hmac: prk,
		info: info,
		i:    1,
		temp: make([]byte, hash().Size()),
		out:  nil,
	}
}

// Read implements io.Reader for hkdfReader
func (r *hkdfReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	// If we have leftover bytes from previous expansion, use them first
	if r.out != nil {
		n = copy(p, r.out)
		if len(r.out) == n {
			r.out = nil
		} else {
			r.out = r.out[n:]
		}
		return n, nil
	}

	// Calculate how many expansion rounds we need
	hashSize := r.mac.Size()
	remaining := len(p)
	written := 0

	for remaining > 0 {
		// Reset the MAC for a new round of expansion
		r.mac.Reset()

		// Write previous HMAC output if this isn't the first round
		if r.i > 1 {
			r.mac.Write(r.temp)
		}

		// Write the info and the round counter
		r.mac.Write(r.info)
		r.mac.Write([]byte{r.i})

		// Calculate the HMAC for this round
		copy(r.temp, r.mac.Sum(nil))

		// Copy as much as we can to the output buffer
		toCopy := hashSize
		if toCopy > remaining {
			toCopy = remaining
		}

		nCopied := copy(p[written:], r.temp[:toCopy])
		written += nCopied
		remaining -= nCopied

		// Increment the round counter
		r.i++

		// Check if we've reached the maximum number of rounds (255)
		if r.i > 255 {
			return len(p) - remaining, io.EOF
		}
	}

	return len(p), nil
}

func deriveKeyBasic(hash func() hash.Hash, dh *ecdsa.PublicKey, len int) ([]byte, error) {
	dhb := ECDHMarshalBinary(dh)
	hkdf := NewHKDF(hash, dhb, nil, nil)
	key := make([]byte, len)
	n, err := hkdf.Read(key)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if n != len {
		return nil, fmt.Errorf("hkdf.Read returned incomplete key: expected %d bytes, got %d bytes", len, n)
	}
	return key, nil
}

func LoadLocalCertificate(model string) (string, error) {
	if model == "" || strings.ContainsAny(model, "/\\:..") {
		return "", fmt.Errorf("invalid model parameter: %s", model)
	}
	var dir string
	home, herr := os.UserHomeDir()
	if herr == nil && home != "" {
		dir = filepath.Join(home, ".ark", "certificates")
	} else {
		return "", fmt.Errorf("failed to get user home dir. err=%w", herr)
	}

	certFileName := filepath.Clean(model) + ".pem"
	certFilePath := filepath.Join(dir, certFileName)
	fi, err := os.Stat(certFilePath)
	if err != nil {
		return "", err
	}
	lastModified := fi.ModTime()
	current := time.Now()
	if current.Sub(lastModified) <= 14*24*time.Hour {
		b, err := os.ReadFile(certFilePath)
		if err != nil {
			return "", err
		}
		certPem := string(b)
		ringID, keyID, _ := GetCertInfo(certPem)
		aiccEnabled := CheckIsModeAICC()

		if (ringID == "" || keyID == "") && !aiccEnabled {
			return certPem, nil
		}
		if ringID != "" && keyID != "" && aiccEnabled {
			return certPem, nil
		}
	}
	_ = os.Remove(certFilePath)
	return "", nil
}

func SaveToLocalCertificate(model, certPem string) error {
	if model == "" || strings.ContainsAny(model, "/\\:..") {
		return fmt.Errorf("invalid model parameter: %s", model)
	}
	var dir string
	home, herr := os.UserHomeDir()
	if herr == nil && home != "" {
		dir = filepath.Join(home, ".ark", "certificates")
	} else {
		return fmt.Errorf("failed to get user home dir. err=%w", herr)
	}

	certFileName := filepath.Clean(model) + ".pem"
	certFilePath := filepath.Join(dir, certFileName)
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}
	return os.WriteFile(certFilePath, []byte(certPem), 0o644)
}

func CheckIsModeAICC() bool {
	return os.Getenv("VOLC_ARK_ENCRYPTION") == "AICC"
}

func EncryptChatRequest(ctx context.Context, keyNonce []byte, request model.CreateChatCompletionRequest) error {
	err := ProcessChatCompletionRequest(ctx, request.Messages, func(text string) (string, error) {
		if len(keyNonce) != aesKeySize+aesNonceSize {
			return "", fmt.Errorf("keyNonce must be %d bytes long, got %d", aesKeySize+aesNonceSize, len(keyNonce))
		}
		return AesGcmEncryptBase64String(keyNonce[:aesKeySize], keyNonce[aesKeySize:], text)
	})
	if err != nil {
		return err
	}
	return nil
}

// ProcessChatCompletionRequest process chat completion request
func ProcessChatCompletionRequest(ctx context.Context, msgs []*model.ChatCompletionMessage, fn func(text string) (string, error)) error {
	var err error
	for _, m := range msgs {
		if m.Content == nil {
			continue
		}
		err = ProcessChatCompletionMessageContent(m.Content, fn)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessChatCompletionMessageContent(content *model.ChatCompletionMessageContent, fn func(text string) (string, error)) error {
	var err error
	if content.StringValue != nil {
		var text string
		text, err = fn(*content.StringValue)
		if err != nil {
			return err
		}
		content.StringValue = &text
	}
	for _, p := range content.ListValue {
		if len(p.Text) != 0 {
			p.Text, err = fn(p.Text)
			if err != nil {
				return err
			}
		}
		if p.ImageURL != nil {
			p.ImageURL.URL, err = EncryptURL(p.ImageURL.URL, fn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func EncryptURL(urlString string, fn func(text string) (string, error)) (string, error) {
	res, err := url.Parse(urlString)
	if err != nil {
		return urlString, err
	}
	if StringInSlice(res.Scheme, []string{"https", "http", "file", "ftp"}) {
		// TBD: warn info
		// fmt.Println("encryption is not supported for image url, please use base64 image if you want encryption")
		return urlString, nil
	} else if res.Scheme == "data" {
		var newURL string
		newURL, err = fn(res.Opaque)
		if err != nil {
			return urlString, err
		}
		return newURL, nil
	} else {
		return urlString, fmt.Errorf("encryption is not supported for image url scheme: %s", res.Scheme)
	}
}

func StringInSlice(str string, list []string) bool {
	for _, val := range list {
		if val == str {
			return true
		}
	}
	return false
}

func DecryptChatResponse(keyNonce []byte, response model.Response) error {
	var err error
	if len(keyNonce) != aesKeySize+aesNonceSize {
		return fmt.Errorf("keyNonce must be %d bytes long, got %d", aesKeySize+aesNonceSize, len(keyNonce))
	}
	fn := func(text string) (string, error) {
		return AesGcmDecryptBase64String(keyNonce[:aesKeySize], keyNonce[aesKeySize:], text)
	}
	if cr, ok := response.(*model.ChatCompletionResponse); ok {
		for _, choice := range cr.Choices {
			if choice.FinishReason != model.FinishReasonContentFilter && choice.Message.Content != nil {
				err = ProcessChatCompletionMessageContent(choice.Message.Content, fn)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func DecryptChatStreamResponse(keyNonce []byte, response model.ChatCompletionStreamResponse) error {
	expectedLen := aesKeySize + aesNonceSize
	if len(keyNonce) != expectedLen {
		return fmt.Errorf("keyNonce must be %d bytes long, got %d", expectedLen, len(keyNonce))
	}
	var err error
	fn := func(text string) (string, error) {
		return AesGcmDecryptBase64String(keyNonce[:aesKeySize], keyNonce[aesKeySize:], text)
	}
	for _, choice := range response.Choices {
		if choice.FinishReason != model.FinishReasonContentFilter && len(choice.Delta.Content) > 0 {
			choice.Delta.Content, err = fn(choice.Delta.Content)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
