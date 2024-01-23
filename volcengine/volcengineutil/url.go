package volcengineutil

import (
	"code.byted.org/eps-platform/volcengine-innersdk-golang/volcengine/volcengineerr"
)

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

type Endpoint struct {
	//UseSSL bool
	//Locate bool
	//UseInternal                 bool
	CustomerEndpoint string
	//CustomerDomainIgnoreService bool

}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (c *Endpoint) WithCustomerEndpoint(customerEndpoint string) *Endpoint {
	c.CustomerEndpoint = customerEndpoint
	return c
}

type ServiceInfo struct {
	Service string
	Region  string
}

const (
	endpoint = "open.volcengineapi.com"

	//internalUrl = "open.volcengineapi.com"
	//http  = "http"
	//https = "https"

)

func (c *Endpoint) GetEndpoint() string {
	if c.CustomerEndpoint != "" {
		return c.CustomerEndpoint
	} else {
		return endpoint
	}
}

const (
	separator      = "."
	endpointSuffix = separator + "commones.org"
)

const (
	regionCodeCNBeijing  = "cn-beijing"
	regionCodeCNShanghai = "cn-shanghai"
	regionCodeCNNantong  = "cn-nantong"
)
const (
	ErrorCodeDefaultEndpointServiceNotExist        = "DefaultEndpointServiceNotExistError"
	ErrorCodeDefaultEndpointRegionEndpointNotExist = "DefaultEndpointRegionEndpointNotExistError"
)

var ErrorDefaultEndpointServiceNotExist = volcengineerr.New(ErrorCodeDefaultEndpointServiceNotExist, "service info not exist when fetch default endpoint", nil)
var ErrorDefaultEndpointRegionEndpointNotExist = volcengineerr.New(ErrorCodeDefaultEndpointRegionEndpointNotExist, "region endpoint info not exist when fetch default endpoint", nil)

type RegionEndpointMap map[string]string

type ServiceEndpointInfo struct {
	Service        string
	IsGlobal       bool
	GlobalEndpoint string
	RegionEndpointMap
}

var defaultEndpoint = map[string]*ServiceEndpointInfo{
	"iam": {
		Service:           "iam",
		IsGlobal:          true,
		GlobalEndpoint:    "iam" + endpointSuffix,
		RegionEndpointMap: nil,
	},
	//"vpc": {
	//	Service:        "vpc",
	//	IsGlobal:       false,
	//	GlobalEndpoint: "",
	//	RegionEndpointMap: map[string]string{
	//		regionCodeCNBeijing:  regionCodeCNBeijing + separator + "vpc" + endpointSuffix,
	//		regionCodeCNShanghai: regionCodeCNShanghai + separator + "vpc" + endpointSuffix,
	//		regionCodeCNNantong:  regionCodeCNNantong + separator + "vpc" + endpointSuffix,
	//	},
	//},
	//"fw_center": {
	//	Service:        "fw_center",
	//	IsGlobal:       false,
	//	GlobalEndpoint: "",
	//	RegionEndpointMap: map[string]string{
	//		regionCodeCNBeijing:  regionCodeCNBeijing + separator + "fw-center" + endpointSuffix,
	//		regionCodeCNShanghai: regionCodeCNShanghai + separator + "fw-center" + endpointSuffix,
	//		regionCodeCNNantong:  regionCodeCNNantong + separator + "fw-center" + endpointSuffix,
	//	},
	//},
}

func GetDefaultEndpointByServiceInfo(service string, regionCode string) (*string, volcengineerr.Error) {
	defaultEndpointInfo, sExist := defaultEndpoint[service]
	if sExist == false {
		return nil, ErrorDefaultEndpointServiceNotExist
	}

	isGlobal := defaultEndpointInfo.IsGlobal
	if isGlobal {
		return &defaultEndpointInfo.GlobalEndpoint, nil
	} else {
		regionEndpointMp := defaultEndpointInfo.RegionEndpointMap
		regionEndpointStr, rExist := regionEndpointMp[regionCode]
		if !rExist {
			return nil, ErrorDefaultEndpointRegionEndpointNotExist
		}

		return &regionEndpointStr, nil
	}
}
