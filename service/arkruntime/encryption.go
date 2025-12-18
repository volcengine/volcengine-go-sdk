package arkruntime

import "github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/encryption"

type ClientE2EE struct {
	certificate string
	cipher      *encryption.KeyAgreementClient
}

func NewE2EE(certificate string) (*ClientE2EE, error) {
	cipher, err := encryption.NewP256KeyAgreementClient(certificate)
	if err != nil {
		return nil, err
	}
	client := &ClientE2EE{
		certificate: certificate,
		cipher:      cipher,
	}
	return client, nil
}
