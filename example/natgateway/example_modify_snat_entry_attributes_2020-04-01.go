package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifySnatEntryAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	modifySnatEntryAttributesInput := &natgateway.ModifySnatEntryAttributesInput{
		EipId:       volcengine.String("eip-2feaac9wtccn459gp67qe****"),
		SnatEntryId: volcengine.String("snat-2fedi096gdiww59gp680r****"),
	}

	resp, err := svc.ModifySnatEntryAttributes(modifySnatEntryAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
