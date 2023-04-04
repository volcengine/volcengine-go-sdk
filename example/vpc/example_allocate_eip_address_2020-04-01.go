package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AllocateEipAddress() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	allocateEipAddressInput := &vpc.AllocateEipAddressInput{
		Bandwidth: volcengine.Int64(10),
		BillingType: volcengine.Int64(2),
		ISP: volcengine.String("BGP"),
	}

	resp, err := svc.AllocateEipAddress(allocateEipAddressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
