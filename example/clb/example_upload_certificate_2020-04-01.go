package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func UploadCertificate() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	uploadCertificateInput := &clb.UploadCertificateInput{
		CertificateName: volcengine.String("mycert1"),
		PrivateKey: volcengine.String("***"),
		PublicKey: volcengine.String("***"),
	}

	resp, err := svc.UploadCertificate(uploadCertificateInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
