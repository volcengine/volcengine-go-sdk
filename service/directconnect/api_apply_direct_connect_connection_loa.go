// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opApplyDirectConnectConnectionLoaCommon = "ApplyDirectConnectConnectionLoa"

// ApplyDirectConnectConnectionLoaCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ApplyDirectConnectConnectionLoaCommon operation. The "output" return
// value will be populated with the ApplyDirectConnectConnectionLoaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ApplyDirectConnectConnectionLoaCommon Request to send the API call to the service.
// the "output" return value is not valid until after ApplyDirectConnectConnectionLoaCommon Send returns without error.
//
// See ApplyDirectConnectConnectionLoaCommon for more information on using the ApplyDirectConnectConnectionLoaCommon
// API call, and error handling.
//
//    // Example sending a request using the ApplyDirectConnectConnectionLoaCommonRequest method.
//    req, resp := client.ApplyDirectConnectConnectionLoaCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ApplyDirectConnectConnectionLoaCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opApplyDirectConnectConnectionLoaCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ApplyDirectConnectConnectionLoaCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation ApplyDirectConnectConnectionLoaCommon for usage and error information.
func (c *DIRECTCONNECT) ApplyDirectConnectConnectionLoaCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ApplyDirectConnectConnectionLoaCommonRequest(input)
	return out, req.Send()
}

// ApplyDirectConnectConnectionLoaCommonWithContext is the same as ApplyDirectConnectConnectionLoaCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ApplyDirectConnectConnectionLoaCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ApplyDirectConnectConnectionLoaCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ApplyDirectConnectConnectionLoaCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opApplyDirectConnectConnectionLoa = "ApplyDirectConnectConnectionLoa"

// ApplyDirectConnectConnectionLoaRequest generates a "volcengine/request.Request" representing the
// client's request for the ApplyDirectConnectConnectionLoa operation. The "output" return
// value will be populated with the ApplyDirectConnectConnectionLoaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ApplyDirectConnectConnectionLoaCommon Request to send the API call to the service.
// the "output" return value is not valid until after ApplyDirectConnectConnectionLoaCommon Send returns without error.
//
// See ApplyDirectConnectConnectionLoa for more information on using the ApplyDirectConnectConnectionLoa
// API call, and error handling.
//
//    // Example sending a request using the ApplyDirectConnectConnectionLoaRequest method.
//    req, resp := client.ApplyDirectConnectConnectionLoaRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ApplyDirectConnectConnectionLoaRequest(input *ApplyDirectConnectConnectionLoaInput) (req *request.Request, output *ApplyDirectConnectConnectionLoaOutput) {
	op := &request.Operation{
		Name:       opApplyDirectConnectConnectionLoa,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ApplyDirectConnectConnectionLoaInput{}
	}

	output = &ApplyDirectConnectConnectionLoaOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ApplyDirectConnectConnectionLoa API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation ApplyDirectConnectConnectionLoa for usage and error information.
func (c *DIRECTCONNECT) ApplyDirectConnectConnectionLoa(input *ApplyDirectConnectConnectionLoaInput) (*ApplyDirectConnectConnectionLoaOutput, error) {
	req, out := c.ApplyDirectConnectConnectionLoaRequest(input)
	return out, req.Send()
}

// ApplyDirectConnectConnectionLoaWithContext is the same as ApplyDirectConnectConnectionLoa with the addition of
// the ability to pass a context and additional request options.
//
// See ApplyDirectConnectConnectionLoa for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ApplyDirectConnectConnectionLoaWithContext(ctx volcengine.Context, input *ApplyDirectConnectConnectionLoaInput, opts ...request.Option) (*ApplyDirectConnectConnectionLoaOutput, error) {
	req, out := c.ApplyDirectConnectConnectionLoaRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ApplyDirectConnectConnectionLoaInput struct {
	_ struct{} `type:"structure"`

	// Action is a required field
	Action *string `type:"string" required:"true"`

	Bandwidth *int64 `type:"integer"`

	ClientToken *string `type:"string"`

	// CompanyName is a required field
	CompanyName *string `type:"string" required:"true"`

	// ConstructionTime is a required field
	ConstructionTime *string `type:"string" required:"true"`

	// DirectConnectConnectionId is a required field
	DirectConnectConnectionId *string `type:"string" required:"true"`

	// Engineers is a required field
	Engineers []*EngineerForApplyDirectConnectConnectionLoaInput `type:"list" required:"true"`

	// LineType is a required field
	LineType *string `type:"string" required:"true"`

	PeerLocation *string `type:"string"`

	// SystemIntegrator is a required field
	SystemIntegrator *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplyDirectConnectConnectionLoaInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ApplyDirectConnectConnectionLoaInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplyDirectConnectConnectionLoaInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ApplyDirectConnectConnectionLoaInput"}
	if s.Action == nil {
		invalidParams.Add(request.NewErrParamRequired("Action"))
	}
	if s.CompanyName == nil {
		invalidParams.Add(request.NewErrParamRequired("CompanyName"))
	}
	if s.ConstructionTime == nil {
		invalidParams.Add(request.NewErrParamRequired("ConstructionTime"))
	}
	if s.DirectConnectConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectConnectionId"))
	}
	if s.Engineers == nil {
		invalidParams.Add(request.NewErrParamRequired("Engineers"))
	}
	if s.LineType == nil {
		invalidParams.Add(request.NewErrParamRequired("LineType"))
	}
	if s.SystemIntegrator == nil {
		invalidParams.Add(request.NewErrParamRequired("SystemIntegrator"))
	}
	if s.Engineers != nil {
		for i, v := range s.Engineers {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Engineers", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAction sets the Action field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetAction(v string) *ApplyDirectConnectConnectionLoaInput {
	s.Action = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetBandwidth(v int64) *ApplyDirectConnectConnectionLoaInput {
	s.Bandwidth = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetClientToken(v string) *ApplyDirectConnectConnectionLoaInput {
	s.ClientToken = &v
	return s
}

// SetCompanyName sets the CompanyName field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetCompanyName(v string) *ApplyDirectConnectConnectionLoaInput {
	s.CompanyName = &v
	return s
}

// SetConstructionTime sets the ConstructionTime field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetConstructionTime(v string) *ApplyDirectConnectConnectionLoaInput {
	s.ConstructionTime = &v
	return s
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetDirectConnectConnectionId(v string) *ApplyDirectConnectConnectionLoaInput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetEngineers sets the Engineers field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetEngineers(v []*EngineerForApplyDirectConnectConnectionLoaInput) *ApplyDirectConnectConnectionLoaInput {
	s.Engineers = v
	return s
}

// SetLineType sets the LineType field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetLineType(v string) *ApplyDirectConnectConnectionLoaInput {
	s.LineType = &v
	return s
}

// SetPeerLocation sets the PeerLocation field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetPeerLocation(v string) *ApplyDirectConnectConnectionLoaInput {
	s.PeerLocation = &v
	return s
}

// SetSystemIntegrator sets the SystemIntegrator field's value.
func (s *ApplyDirectConnectConnectionLoaInput) SetSystemIntegrator(v string) *ApplyDirectConnectConnectionLoaInput {
	s.SystemIntegrator = &v
	return s
}

type ApplyDirectConnectConnectionLoaOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ApplyDirectConnectConnectionLoaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ApplyDirectConnectConnectionLoaOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ApplyDirectConnectConnectionLoaOutput) SetRequestId(v string) *ApplyDirectConnectConnectionLoaOutput {
	s.RequestId = &v
	return s
}

type EngineerForApplyDirectConnectConnectionLoaInput struct {
	_ struct{} `type:"structure"`

	// CertificateNo is a required field
	CertificateNo *string `type:"string" required:"true"`

	// CertificateType is a required field
	CertificateType *string `type:"string" required:"true"`

	// ContactPhone is a required field
	ContactPhone *string `type:"string" required:"true"`

	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EngineerForApplyDirectConnectConnectionLoaInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EngineerForApplyDirectConnectConnectionLoaInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EngineerForApplyDirectConnectConnectionLoaInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EngineerForApplyDirectConnectConnectionLoaInput"}
	if s.CertificateNo == nil {
		invalidParams.Add(request.NewErrParamRequired("CertificateNo"))
	}
	if s.CertificateType == nil {
		invalidParams.Add(request.NewErrParamRequired("CertificateType"))
	}
	if s.ContactPhone == nil {
		invalidParams.Add(request.NewErrParamRequired("ContactPhone"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertificateNo sets the CertificateNo field's value.
func (s *EngineerForApplyDirectConnectConnectionLoaInput) SetCertificateNo(v string) *EngineerForApplyDirectConnectConnectionLoaInput {
	s.CertificateNo = &v
	return s
}

// SetCertificateType sets the CertificateType field's value.
func (s *EngineerForApplyDirectConnectConnectionLoaInput) SetCertificateType(v string) *EngineerForApplyDirectConnectConnectionLoaInput {
	s.CertificateType = &v
	return s
}

// SetContactPhone sets the ContactPhone field's value.
func (s *EngineerForApplyDirectConnectConnectionLoaInput) SetContactPhone(v string) *EngineerForApplyDirectConnectConnectionLoaInput {
	s.ContactPhone = &v
	return s
}

// SetName sets the Name field's value.
func (s *EngineerForApplyDirectConnectConnectionLoaInput) SetName(v string) *EngineerForApplyDirectConnectConnectionLoaInput {
	s.Name = &v
	return s
}