package main

import (
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func UploadCertificate(svc *clb.CLB) {
	input := &clb.UploadCertificateInput{
		CertificateName: volcengine.String("cert-name"),
		Description:     volcengine.String("cert description"),
		PrivateKey:      volcengine.String("xxx"),
		PublicKey:       volcengine.String("xxx"),
	}

	resp, err := svc.UploadCertificate(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeCertificates(svc *clb.CLB) {
	input := &clb.DescribeCertificatesInput{
		PageSize:   volcengine.Int64(100),
		PageNumber: volcengine.Int64(1),
	}

	resp, err := svc.DescribeCertificates(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteCertificate(svc *clb.CLB) {
	input := &clb.DeleteCertificateInput{
		CertificateId: volcengine.String("cert-xxx"),
	}

	resp, err := svc.DeleteCertificate(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}
