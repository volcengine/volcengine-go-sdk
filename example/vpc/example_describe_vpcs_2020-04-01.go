package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeVpcs() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	describeVpcsInput := &vpc.DescribeVpcsInput{
		VpcIds: volcengine.StringSlice([]string{"vpc-2ff3****zwc1s5oxru"}),
		VpcName: volcengine.String("test"),
	}

	resp, err := svc.DescribeVpcs(describeVpcsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
