package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyRouteEntry() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifyRouteEntryInput := &vpc.ModifyRouteEntryInput{
		RouteEntryId:   volcengine.String("rte-3tj9gw2pwq3vj******"),
		RouteEntryName: volcengine.String("test"),
	}

	resp, err := svc.ModifyRouteEntry(modifyRouteEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
