package arkruntime

import (
	"context"
	"fmt"
	"net/url"

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
		text, err := fn(*content.StringValue)
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
			p.ImageURL.URL, err = fn(p.ImageURL.URL)
			if err != nil {
				return err
			}
		}
		if p.VideoURL != nil {
			p.VideoURL.URL, err = fn(p.VideoURL.URL)
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
		return "", err
	}
	if StringInSlice(res.Scheme, []string{"https", "http", "file", "ftp"}) {
		fmt.Println("encryption is not supported for image url, please use base64 image if you want encryption")
		return urlString, nil
	} else if res.Scheme == "data" {
		var newURL string
		newURL, err = fn(res.Opaque)
		if err != nil {
			return "", err
		}
		return newURL, nil
	} else {
		return "", fmt.Errorf("encryption is not supported for image url scheme: %s", res.Scheme)
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
	fn := func(text string) (string, error) {
		return encryption.AesGcmDecryptBase64String(keyNonce[:32], keyNonce[32:], text)
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
