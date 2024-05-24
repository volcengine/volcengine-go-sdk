// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateObjectGroupCommon = "CreateObjectGroup"

// CreateObjectGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateObjectGroupCommon operation. The "output" return
// value will be populated with the CreateObjectGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateObjectGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateObjectGroupCommon Send returns without error.
//
// See CreateObjectGroupCommon for more information on using the CreateObjectGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateObjectGroupCommonRequest method.
//    req, resp := client.CreateObjectGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) CreateObjectGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateObjectGroupCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateObjectGroupCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation CreateObjectGroupCommon for usage and error information.
func (c *VOLCOBSERVE) CreateObjectGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateObjectGroupCommonRequest(input)
	return out, req.Send()
}

// CreateObjectGroupCommonWithContext is the same as CreateObjectGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateObjectGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) CreateObjectGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateObjectGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateObjectGroup = "CreateObjectGroup"

// CreateObjectGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateObjectGroup operation. The "output" return
// value will be populated with the CreateObjectGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateObjectGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateObjectGroupCommon Send returns without error.
//
// See CreateObjectGroup for more information on using the CreateObjectGroup
// API call, and error handling.
//
//    // Example sending a request using the CreateObjectGroupRequest method.
//    req, resp := client.CreateObjectGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) CreateObjectGroupRequest(input *CreateObjectGroupInput) (req *request.Request, output *CreateObjectGroupOutput) {
	op := &request.Operation{
		Name:       opCreateObjectGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateObjectGroupInput{}
	}

	output = &CreateObjectGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateObjectGroup API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation CreateObjectGroup for usage and error information.
func (c *VOLCOBSERVE) CreateObjectGroup(input *CreateObjectGroupInput) (*CreateObjectGroupOutput, error) {
	req, out := c.CreateObjectGroupRequest(input)
	return out, req.Send()
}

// CreateObjectGroupWithContext is the same as CreateObjectGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateObjectGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) CreateObjectGroupWithContext(ctx volcengine.Context, input *CreateObjectGroupInput, opts ...request.Option) (*CreateObjectGroupOutput, error) {
	req, out := c.CreateObjectGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateObjectGroupInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `type:"string" required:"true"`

	Objects []*ObjectForCreateObjectGroupInput `type:"list"`
}

// String returns the string representation
func (s CreateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateObjectGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateObjectGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateObjectGroupInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetName sets the Name field's value.
func (s *CreateObjectGroupInput) SetName(v string) *CreateObjectGroupInput {
	s.Name = &v
	return s
}

// SetObjects sets the Objects field's value.
func (s *CreateObjectGroupInput) SetObjects(v []*ObjectForCreateObjectGroupInput) *CreateObjectGroupInput {
	s.Objects = v
	return s
}

type CreateObjectGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data *string `type:"string"`
}

// String returns the string representation
func (s CreateObjectGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateObjectGroupOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *CreateObjectGroupOutput) SetData(v string) *CreateObjectGroupOutput {
	s.Data = &v
	return s
}

type ObjectForCreateObjectGroupInput struct {
	_ struct{} `type:"structure"`

	Dimensions map[string][]*string `type:"map"`

	Id *string `type:"string"`

	Namespace *string `type:"string"`

	Region *string `type:"string"`
}

// String returns the string representation
func (s ObjectForCreateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ObjectForCreateObjectGroupInput) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *ObjectForCreateObjectGroupInput) SetDimensions(v map[string][]*string) *ObjectForCreateObjectGroupInput {
	s.Dimensions = v
	return s
}

// SetId sets the Id field's value.
func (s *ObjectForCreateObjectGroupInput) SetId(v string) *ObjectForCreateObjectGroupInput {
	s.Id = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *ObjectForCreateObjectGroupInput) SetNamespace(v string) *ObjectForCreateObjectGroupInput {
	s.Namespace = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ObjectForCreateObjectGroupInput) SetRegion(v string) *ObjectForCreateObjectGroupInput {
	s.Region = &v
	return s
}
