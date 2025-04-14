package wafruntime

import (
	"github.com/volcengine/volcengine-go-sdk/service/waf"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

// WAF provides the API operation methods for making requests to
// WAF. See this package's package overview docs
// for details on the service.
//
// WAF methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type WAFRuntime struct {
	*waf.WAF
	StreamBuf     string
	StreamSendLen int
	MsgID         *string
	defaultOut    *waf.CheckLLMResponseStreamOutput
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "waf"       // Name of service.
	EndpointsID = ServiceName // ID to lookup a service endpoint with.
	ServiceID   = "waf"       // ServiceID is a unique identifer of a specific service.
)

// New create int can support ssl or region locate set
func New(p client.ConfigProvider, cfgs ...*volcengine.Config) *WAFRuntime {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg volcengine.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *WAFRuntime {
	svc := &WAFRuntime{
		WAF: &waf.WAF{
			Client: client.New(
				cfg,
				metadata.ClientInfo{
					ServiceName:   ServiceName,
					ServiceID:     ServiceID,
					SigningName:   signingName,
					SigningRegion: signingRegion,
					Endpoint:      endpoint,
					APIVersion:    "2023-12-25",
				},
				handlers,
			),
		},
	}

	// Handlers
	svc.Handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	svc.Handlers.Build.PushBackNamed(corehandlers.AddHostExecEnvUserAgentHandler)
	svc.Handlers.Sign.PushBackNamed(volc.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(volcenginequery.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(volcenginequery.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(volcenginequery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(volcenginequery.UnmarshalErrorHandler)
	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a WAF operation and runs any
// custom request initialization.
func (s *WAFRuntime) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
