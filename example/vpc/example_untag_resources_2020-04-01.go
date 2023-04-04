package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func UntagResources() {
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
	untagResourcesInput := &vpc.UntagResourcesInput{
		ResourceIds: volcengine.StringSlice([]string{"vpc-273w3e33y2y9s7fap8u2j****","vpc-bp15zckdt37pq72zv****"}),
		ResourceType: volcengine.String("vpc"),
		TagKeys: volcengine.StringSlice([]string{"k1","k2"}),
	}

	resp, err := svc.UntagResources(untagResourcesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
