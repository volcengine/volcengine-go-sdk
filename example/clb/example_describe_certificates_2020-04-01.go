package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeCertificates() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeCertificatesInput := &clb.DescribeCertificatesInput{
		CertificateIds: volcengine.StringSlice([]string{"cert-3tjuxoukkq3vj0ww****"}),
	}

	resp, err := svc.DescribeCertificates(describeCertificatesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
