package arkruntime

import (
	"context"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/encryption"
)

type E2eeClient struct {
	certificate string
	info        model.EncryptInfo
	isAICC      bool
	cipher      *encryption.KeyAgreementClient
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
		info: model.EncryptInfo{
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

func EncryptChatRequest(ctx context.Context, keyNonce []byte, request model.CreateChatCompletionRequest) error {
	key := keyNonce[:32]
	nonce := keyNonce[32:]
	err := TraversalChatCompletionMessageHandler(ctx, request.Messages, func(text string) (string, error) {
		return encryption.AesGcmEncryptBase64String(key, nonce, text)
	})
	if err != nil {
		return err
	}
	return nil
}

// TraversalChatCompletionMessageHandler traversal chat messages handler
func TraversalChatCompletionMessageHandler(ctx context.Context, msgs []*model.ChatCompletionMessage, fn func(text string) (string, error)) error {
	var err error
	for _, m := range msgs {
		if m.Content == nil {
			continue
		}
		if m.Content.StringValue != nil {
			text, err := fn(*m.Content.StringValue)
			if err != nil {
				return err
			}
			m.Content.StringValue = &text
		}
		for _, p := range m.Content.ListValue {
			if len(p.Text) != 0 {
				p.Text, err = fn(p.Text)
				if err != nil {
					return err
				}
			}
		}
	}
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
