package arkruntime

import (
	"encoding/json"

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
		client.info.Version = model.CipherVersionAICCv01
		client.info.RingID = ringID
		client.info.KeyID = keyID
		client.isAICC = true
	}
	return client, nil
}

func (c *E2eeClient) GenerateECIESKeyPair() ([]byte, string, error) {
	return c.cipher.GenerateECIESKeyPair()
}

func (c *E2eeClient) GetEncryptInfo() string {
	info, err := json.Marshal(c.info)
	if err != nil {
		return ""
	}
	return string(info)
}
