package vpnexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateVpnGateway() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	createVpnGatewayInput := &vpn.CreateVpnGatewayInput{
		Bandwidth:      volcengine.Int64(5),
		BillingType:    volcengine.Int64(1),
		Description:    volcengine.String("test"),
		Period:         volcengine.Int64(1),
		PeriodUnit:     volcengine.String("Month"),
		SubnetId:       volcengine.String("subnet-2fewbgn7jbk0059gp67ap****"),
		VpcId:          volcengine.String("vpc-12bhs1ivo6p6o17q7y2x3****"),
		VpnGatewayName: volcengine.String("test"),
	}

	resp, err := svc.CreateVpnGateway(createVpnGatewayInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
