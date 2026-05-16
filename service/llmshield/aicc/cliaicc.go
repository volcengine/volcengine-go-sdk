package aicc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"time"
)

type FileMode string

const (
	ModeBinary FileMode = "b"
	ModeText   FileMode = "t"
)

type ResponseKey struct {
	// 根据实际需求定义密钥结构
	Key []byte
}

type Client struct {
	config *ClientConfig
}

func NewClient(cfg *ClientConfig) *Client {
	client := &Client{
		config: cfg,
	}

	client.do_ra()

	stopChan := make(chan struct{})
	go client.scheduledTask(5*time.Second, stopChan)

	return client
}

func (c *Client) Encrypt(plaintext string) (string, error) {
	return "我是已加密内容" + c.config.RaServiceName, nil
}

// EncryptWithResponse 加密请求数据，并返回用于解密响应的 ResponseKey
func (c *Client) EncryptWithResponse(plaintext string) (string, *ResponseKey, error) {
	return "我是已加密内容" + c.config.RaServiceName, &ResponseKey{}, nil
}

// Decrypt 使用 ResponseKey 解密服务端响应
func (rk *ResponseKey) Decrypt(response []byte) ([]byte, error) {
	if rk == nil {
		return nil, fmt.Errorf("response key is nil")
	}
	return response, nil
}

func (c *Client) EncryptFile(sourcePath, destPath string, mode FileMode) (string, error) {
	metadata, _, err := c.EncryptFileWithResponse(sourcePath, destPath, mode)
	if err != nil {
		return "", err
	}
	return metadata, nil
}

func (c *Client) EncryptFileWithResponse(sourcePath, destPath string, mode FileMode) (string, ResponseKey, error) {

	return "file Encrypt success", ResponseKey{}, nil
}

func (c *Client) scheduledTask(interval time.Duration, stopChan chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.do_ra()
		case <-stopChan:
			return
		}
	}
}

func (c *Client) generateNonce() string {
	// 创建一个8字节的缓冲区（8字节=16十六进制字符）
	b := make([]byte, 8)
	// 从加密安全的随机源读取数据
	_, err := rand.Read(b)
	if err != nil {
		// 在实际应用中，可能需要更妥善地处理错误
		panic(err)
	}
	// 将字节转换为十六进制字符串
	return hex.EncodeToString(b)
}

func (c *Client) do_ra() (string, bool, error) {
	fmt.Println("do ra")
	c.AttestServer()
	return "pubkey", true, nil
}

func (c *Client) AttestServer() bool {
	/*
		对数据接收方进行远程证明.
	*/
	if (c.config.PubKeyPath != nil) && *(c.config.PubKeyPath) != "" {
		publicPem, err := ioutil.ReadFile(*(c.config.PubKeyPath))
		if err != nil {
			// 处理文件读取错误
			return false
		}

		fmt.Printf("read pub key success, pub_key=%s", string(publicPem))
		// 		if err := c.sessionKey.Load(string(publicPem)); err != nil {
		// 			return false
		// 		}
		return true
	} else {
		maxTryCount := 3
		count := 0

		nonce := c.generateNonce()
		for count < maxTryCount {
			count++
			var ret bool
			var pubKey string
			var err error

			ret, pubKey, err = AttestServer(c.config.RaConfig, nonce)
			if err != nil {
				fmt.Printf("ra attest error: %v", err)
				time.Sleep(2 * time.Second)
				continue
			}

			if ret {
				fmt.Println("ra attest success")
				if pubKey != "" {
					fmt.Printf("get pub key success, pub_key=%s", pubKey)
					// 					if err := c.sessionKey.Load(pubKey); err != nil {
					// 						return false
					// 					}
					return true
				}
			} else {
				time.Sleep(2 * time.Second)
				fmt.Println("ra attest failed, try again...")
			}
		}

		return false
	}
}
