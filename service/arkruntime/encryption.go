package arkruntime

import (
	"context"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/encryption"
)

type E2eeClient struct {
	certificate string
	info        EncryptInfo
	isAICC      bool
	cipher      *encryption.KeyAgreementClient
}

type EncryptInfo struct {
	Version    string `thrift:"Version,1,optional" header:"Version" json:"Version,omitempty"`
	RingID     string `thrift:"RingID,2,optional" header:"RingID" json:"RingID,omitempty"`
	KeyID      string `thrift:"KeyID,3,optional" header:"KeyID" json:"KeyID,omitempty"`
	ExpireTime int64  `thrift:"ExpireTime,4,optional" header:"ExpireTime" json:"ExpireTime,omitempty"`
}

func NewE2eeClient(certificate string) (*E2eeClient, error) {
	cipher, err := encryption.NewP256KeyAgreementClient(certificate)
	if err != nil {
		return nil, err
	}
	ringID, keyID, expireTime := encryption.GetCertInfo(certificate)
	client := &E2eeClient{
		certificate: certificate,
		cipher:      cipher,
		info: EncryptInfo{
			ExpireTime: expireTime,
		},
		isAICC: false,
	}
	if ringID != "" && keyID != "" {
		client.info.Version = "AICCv0.1"
		client.info.RingID = ringID
		client.info.KeyID = keyID
		client.isAICC = true
	}
	return client, nil
}

func (c *E2eeClient) GenerateECIESKeyPair() ([]byte, string, error) {
	return c.cipher.GenerateECIESKeyPair()
}

func EncryptChatRequest(ctx context.Context, keyNonce []byte, request model.ChatRequest) error {
	return nil
}

func DecryptChatResponse(ctx context.Context, keyNonce []byte, response model.Response) error {
	if cr, ok := response.(*model.ChatCompletionResponse); ok {
		for i := range cr.Choices {
			if cr.Choices[i] != nil {
				cr.Choices[i].Index = i
			}
		}
	}
	return nil
}
