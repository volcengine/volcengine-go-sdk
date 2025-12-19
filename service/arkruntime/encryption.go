package arkruntime

import (
	"context"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/encryption"
)

type E2eeClient struct {
	certificate string
	cipher      *encryption.KeyAgreementClient
}

func NewE2eeClient(certificate string) (*E2eeClient, error) {
	cipher, err := encryption.NewP256KeyAgreementClient(certificate)
	if err != nil {
		return nil, err
	}
	client := &E2eeClient{
		certificate: certificate,
		cipher:      cipher,
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
