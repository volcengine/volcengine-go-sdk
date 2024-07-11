package volcengineutil

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
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

func (c *Endpoint) GetEndpoint() string {
	if c.CustomerEndpoint != "" {
		return c.CustomerEndpoint
	} else {
		return endpoint
	}
}

const (
	separator      = "."
	openPrefix     = "open"
	endpointSuffix = separator + "volcengineapi.com"
)

var endpoint = openPrefix + endpointSuffix

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
}

func GetDefaultEndpointByServiceInfo(service string, regionCode string) (*string, volcengineerr.Error) {
	defaultEndpointInfo, sExist := defaultEndpoint[service]
	if !sExist {
		return &endpoint, nil
	}

	isGlobal := defaultEndpointInfo.IsGlobal
	if isGlobal {
		return &defaultEndpointInfo.GlobalEndpoint, nil
	}

	regionEndpointMp := defaultEndpointInfo.RegionEndpointMap
	regionEndpointStr, rExist := regionEndpointMp[regionCode]
	if !rExist {
		return &endpoint, nil
	}

	return &regionEndpointStr, nil
}
