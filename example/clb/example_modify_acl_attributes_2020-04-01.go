package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyAclAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	modifyAclAttributesInput := &clb.ModifyAclAttributesInput{
		AclId: volcengine.String("acl-3cj44nv0jhhxc6c6rrtet****"),
		AclName: volcengine.String("baaa"),
	}

	resp, err := svc.ModifyAclAttributes(modifyAclAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
