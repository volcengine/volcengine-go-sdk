// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTransitRouterTrafficQosMarkingEntryCommon = "CreateTransitRouterTrafficQosMarkingEntry"

// CreateTransitRouterTrafficQosMarkingEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterTrafficQosMarkingEntryCommon operation. The "output" return
// value will be populated with the CreateTransitRouterTrafficQosMarkingEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterTrafficQosMarkingEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterTrafficQosMarkingEntryCommon Send returns without error.
//
// See CreateTransitRouterTrafficQosMarkingEntryCommon for more information on using the CreateTransitRouterTrafficQosMarkingEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterTrafficQosMarkingEntryCommonRequest method.
//    req, resp := client.CreateTransitRouterTrafficQosMarkingEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterTrafficQosMarkingEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTransitRouterTrafficQosMarkingEntryCommon,
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

// CreateTransitRouterTrafficQosMarkingEntryCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterTrafficQosMarkingEntryCommon for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterTrafficQosMarkingEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterTrafficQosMarkingEntryCommonRequest(input)
	return out, req.Send()
}

// CreateTransitRouterTrafficQosMarkingEntryCommonWithContext is the same as CreateTransitRouterTrafficQosMarkingEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterTrafficQosMarkingEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterTrafficQosMarkingEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTransitRouterTrafficQosMarkingEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTransitRouterTrafficQosMarkingEntry = "CreateTransitRouterTrafficQosMarkingEntry"

// CreateTransitRouterTrafficQosMarkingEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTransitRouterTrafficQosMarkingEntry operation. The "output" return
// value will be populated with the CreateTransitRouterTrafficQosMarkingEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTransitRouterTrafficQosMarkingEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTransitRouterTrafficQosMarkingEntryCommon Send returns without error.
//
// See CreateTransitRouterTrafficQosMarkingEntry for more information on using the CreateTransitRouterTrafficQosMarkingEntry
// API call, and error handling.
//
//    // Example sending a request using the CreateTransitRouterTrafficQosMarkingEntryRequest method.
//    req, resp := client.CreateTransitRouterTrafficQosMarkingEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) CreateTransitRouterTrafficQosMarkingEntryRequest(input *CreateTransitRouterTrafficQosMarkingEntryInput) (req *request.Request, output *CreateTransitRouterTrafficQosMarkingEntryOutput) {
	op := &request.Operation{
		Name:       opCreateTransitRouterTrafficQosMarkingEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitRouterTrafficQosMarkingEntryInput{}
	}

	output = &CreateTransitRouterTrafficQosMarkingEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTransitRouterTrafficQosMarkingEntry API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation CreateTransitRouterTrafficQosMarkingEntry for usage and error information.
func (c *TRANSITROUTER) CreateTransitRouterTrafficQosMarkingEntry(input *CreateTransitRouterTrafficQosMarkingEntryInput) (*CreateTransitRouterTrafficQosMarkingEntryOutput, error) {
	req, out := c.CreateTransitRouterTrafficQosMarkingEntryRequest(input)
	return out, req.Send()
}

// CreateTransitRouterTrafficQosMarkingEntryWithContext is the same as CreateTransitRouterTrafficQosMarkingEntry with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTransitRouterTrafficQosMarkingEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) CreateTransitRouterTrafficQosMarkingEntryWithContext(ctx volcengine.Context, input *CreateTransitRouterTrafficQosMarkingEntryInput, opts ...request.Option) (*CreateTransitRouterTrafficQosMarkingEntryOutput, error) {
	req, out := c.CreateTransitRouterTrafficQosMarkingEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTransitRouterTrafficQosMarkingEntryInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// DestinationPortEnd is a required field
	DestinationPortEnd *int32 `type:"int32" required:"true"`

	// DestinationPortStart is a required field
	DestinationPortStart *int32 `type:"int32" required:"true"`

	MatchDscp *int32 `type:"int32"`

	// Priority is a required field
	Priority *int32 `type:"int32" required:"true"`

	Protocol *string `type:"string"`

	// RemarkingDscp is a required field
	RemarkingDscp *int32 `type:"int32" required:"true"`

	// SourceCidrBlock is a required field
	SourceCidrBlock *string `type:"string" required:"true"`

	// SourcePortEnd is a required field
	SourcePortEnd *int32 `type:"int32" required:"true"`

	// SourcePortStart is a required field
	SourcePortStart *int32 `type:"int32" required:"true"`

	TransitRouterTrafficQosMarkingEntryName *string `type:"string"`

	// TransitRouterTrafficQosMarkingPolicyId is a required field
	TransitRouterTrafficQosMarkingPolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTransitRouterTrafficQosMarkingEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterTrafficQosMarkingEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTransitRouterTrafficQosMarkingEntryInput"}
	if s.DestinationCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationCidrBlock"))
	}
	if s.DestinationPortEnd == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationPortEnd"))
	}
	if s.DestinationPortStart == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationPortStart"))
	}
	if s.Priority == nil {
		invalidParams.Add(request.NewErrParamRequired("Priority"))
	}
	if s.RemarkingDscp == nil {
		invalidParams.Add(request.NewErrParamRequired("RemarkingDscp"))
	}
	if s.SourceCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceCidrBlock"))
	}
	if s.SourcePortEnd == nil {
		invalidParams.Add(request.NewErrParamRequired("SourcePortEnd"))
	}
	if s.SourcePortStart == nil {
		invalidParams.Add(request.NewErrParamRequired("SourcePortStart"))
	}
	if s.TransitRouterTrafficQosMarkingPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterTrafficQosMarkingPolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetClientToken(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetDescription(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetDestinationCidrBlock(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDestinationPortEnd sets the DestinationPortEnd field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetDestinationPortEnd(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.DestinationPortEnd = &v
	return s
}

// SetDestinationPortStart sets the DestinationPortStart field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetDestinationPortStart(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.DestinationPortStart = &v
	return s
}

// SetMatchDscp sets the MatchDscp field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetMatchDscp(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.MatchDscp = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetPriority(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetProtocol(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.Protocol = &v
	return s
}

// SetRemarkingDscp sets the RemarkingDscp field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetRemarkingDscp(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.RemarkingDscp = &v
	return s
}

// SetSourceCidrBlock sets the SourceCidrBlock field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetSourceCidrBlock(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.SourceCidrBlock = &v
	return s
}

// SetSourcePortEnd sets the SourcePortEnd field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetSourcePortEnd(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.SourcePortEnd = &v
	return s
}

// SetSourcePortStart sets the SourcePortStart field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetSourcePortStart(v int32) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.SourcePortStart = &v
	return s
}

// SetTransitRouterTrafficQosMarkingEntryName sets the TransitRouterTrafficQosMarkingEntryName field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetTransitRouterTrafficQosMarkingEntryName(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.TransitRouterTrafficQosMarkingEntryName = &v
	return s
}

// SetTransitRouterTrafficQosMarkingPolicyId sets the TransitRouterTrafficQosMarkingPolicyId field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryInput) SetTransitRouterTrafficQosMarkingPolicyId(v string) *CreateTransitRouterTrafficQosMarkingEntryInput {
	s.TransitRouterTrafficQosMarkingPolicyId = &v
	return s
}

type CreateTransitRouterTrafficQosMarkingEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	TransitRouterTrafficQosMarkingEntryId *string `type:"string"`
}

// String returns the string representation
func (s CreateTransitRouterTrafficQosMarkingEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTransitRouterTrafficQosMarkingEntryOutput) GoString() string {
	return s.String()
}

// SetTransitRouterTrafficQosMarkingEntryId sets the TransitRouterTrafficQosMarkingEntryId field's value.
func (s *CreateTransitRouterTrafficQosMarkingEntryOutput) SetTransitRouterTrafficQosMarkingEntryId(v string) *CreateTransitRouterTrafficQosMarkingEntryOutput {
	s.TransitRouterTrafficQosMarkingEntryId = &v
	return s
}
