package jeddak_secure_channel

import (
	"sync"
	"errors"
	"io"
	"log"
	"os"
	"time"
)

// Client manages secure channel communication, including server attestation and encryption.
type Client struct {
	config     ClientConfig
	ticker     *time.Ticker
	done       chan struct{}
	closeOnce  sync.Once
	sessionKey *ClientSessionKey
}

// EncryptResult contains the serialized encrypted message and response decryption key.
type EncryptResult struct {
	Ciphertext  string
	ResponseKey ResponseKey
}

// NewClient creates a new Client with the given configuration.
// If AttestInterval is non-nil and positive, periodic server attestation is enabled.
func NewClient(config ClientConfig) *Client {
	c := &Client{
		config: config,
		done:   make(chan struct{}),
	}

	// Start periodic attestation if interval is set
	if config.AttestInterval != nil && *config.AttestInterval > 0 {
		interval := time.Duration(*config.AttestInterval * float64(time.Second))
		c.ticker = time.NewTicker(interval)

		// Run attestation immediately, then on each tick
		go func() {
			if err := c.AttestServer(); err != nil {
				log.Printf("Initial attestation failed: %v", err)
			}

			for {
				select {
				case <-c.ticker.C:
					if err := c.AttestServer(); err != nil {
						log.Printf("Periodic attestation failed: %v", err)
					}
				case <-c.done:
					return
				}
			}
		}()
	}

	return c
}

// Close stops periodic attestation and cleans up resources.
func (c *Client) Close() error {
	c.closeOnce.Do(func() {
		close(c.done)
    })
	if c.ticker != nil {
		c.ticker.Stop()
	}
	return nil
}

// AttestServer triggers immediate server attestation with an empty token.
func (c *Client) AttestServer() error {
	return c.AttestServerWithToken("")
}

// AttestServerWithToken triggers immediate server attestation with a specific token.
func (c *Client) AttestServerWithToken(token string) error {
	log.Println("Client attesting server")
	sessionKey, err := attestServer(token, c.config)
	if err != nil {
		return err
	}

	c.sessionKey = sessionKey
	return nil
}

// EncryptString encrypts a string and returns the ciphertext.
func (c *Client) EncryptString(plaintext string) (string, error) {
	res, err := c.EncryptStringWithResponse(plaintext)
	if err != nil {
		return "", err
	}
	return res.Ciphertext, nil
}

// EncryptBytes encrypts a byte slice and returns the ciphertext.
func (c *Client) EncryptBytes(plaintext []byte) (string, error) {
	res, err := c.EncryptBytesWithResponse(plaintext)
	if err != nil {
		return "", err
	}
	return res.Ciphertext, nil
}

// EncryptStream encrypts data from a reader and writes ciphertext to a writer.
func (c *Client) EncryptStream(source io.Reader, sink io.Writer) (string, error) {
	res, err := c.EncryptStreamWithResponse(source, sink)
	if err != nil {
		return "", err
	}
	return res.Ciphertext, nil
}

// EncryptFile encrypts a file and writes the ciphertext to another file.
func (c *Client) EncryptFile(inputPath, outputPath string) (string, error) {
	res, err := c.EncryptFileWithResponse(inputPath, outputPath)
	if err != nil {
		return "", err
	}
	return res.Ciphertext, nil
}

// EncryptStringWithResponse encrypts a string and returns the full EncryptResult.
func (c *Client) EncryptStringWithResponse(plaintext string) (*EncryptResult, error) {
	return c.EncryptBytesWithResponse([]byte(plaintext))
}

// EncryptBytesWithResponse encrypts a byte slice and returns the full EncryptResult.
func (c *Client) EncryptBytesWithResponse(plaintext []byte) (*EncryptResult, error) {
	sessionKey, err := c.getSessionKey()
	if err != nil {
		return nil, err
	}
	return sessionKey.EncryptWithResponse(plaintext)
}

// EncryptStreamWithResponse encrypts a stream and returns the full EncryptResult.
func (c *Client) EncryptStreamWithResponse(source io.Reader, sink io.Writer) (*EncryptResult, error) {
	sessionKey, err := c.getSessionKey()
	if err != nil {
		return nil, err
	}
	return sessionKey.EncryptStreamWithResponse(source, sink)
}

// EncryptFileWithResponse encrypts a file and returns the full EncryptResult.
func (c *Client) EncryptFileWithResponse(inputPath, outputPath string) (*EncryptResult, error) {
	input, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	output, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}
	defer output.Close()

	return c.EncryptStreamWithResponse(input, output)
}

// getSessionKey retrieves the current session key (thread-safe).
func (c *Client) getSessionKey() (*ClientSessionKey, error) {
	if c.sessionKey == nil {
		return nil, errors.New("session key not initialized; call AttestServer first")
	}
	return c.sessionKey, nil
}
