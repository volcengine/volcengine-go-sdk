package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func UploadCertificate(svc *clb.CLB) {
	input := &clb.UploadCertificateInput{
		CertificateName: volcstack.String("cert-name"),
		Description:     volcstack.String("cert description"),
		PrivateKey:      volcstack.String("xxx"),
		PublicKey:       volcstack.String("xxx"),
	}

	resp, err := svc.UploadCertificate(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeCertificates(svc *clb.CLB) {
	input := &clb.DescribeCertificatesInput{
		PageSize:   volcstack.Int64(100),
		PageNumber: volcstack.Int64(1),
	}

	resp, err := svc.DescribeCertificates(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteCertificate(svc *clb.CLB) {
	input := &clb.DeleteCertificateInput{
		CertificateId: volcstack.String("cert-xxx"),
	}

	resp, err := svc.DeleteCertificate(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}
