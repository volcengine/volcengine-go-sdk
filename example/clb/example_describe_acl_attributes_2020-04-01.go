package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeAclAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeAclAttributesInput := &clb.DescribeAclAttributesInput{
		AclId: volcengine.String("acl-3cj44nv0jhhxc6c6rrtet****"),
	}

	resp, err := svc.DescribeAclAttributes(describeAclAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
