// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/client"
	"github.com/volcengine/volcstack-go-sdk/volcstack/client/metadata"
	"github.com/volcengine/volcstack-go-sdk/volcstack/corehandlers"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/signer/volc"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackquery"
)

// VPN provides the API operation methods for making requests to
// VPN. See this package's package overview docs
// for details on the service.
//
// VPN methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type VPN struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "vpn"       // Name of service.
	EndpointsID = ServiceName // ID to lookup a service endpoint with.
	ServiceID   = "vpn"       // ServiceID is a unique identifer of a specific service.
)

// New create int can support ssl or region locate set
func New(p client.ConfigProvider, cfgs ...*volcstack.Config) *VPN {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg volcstack.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *VPN {
	svc := &VPN{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2020-04-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	svc.Handlers.Build.PushBackNamed(corehandlers.AddHostExecEnvUserAgentHandler)
	svc.Handlers.Sign.PushBackNamed(volc.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(volcstackquery.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(volcstackquery.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(volcstackquery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(volcstackquery.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a VPN operation and runs any
// custom request initialization.
func (c *VPN) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}