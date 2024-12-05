// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateContactsCommon = "CreateContacts"

// CreateContactsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateContactsCommon operation. The "output" return
// value will be populated with the CreateContactsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateContactsCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateContactsCommon Send returns without error.
//
// See CreateContactsCommon for more information on using the CreateContactsCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateContactsCommonRequest method.
//    req, resp := client.CreateContactsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) CreateContactsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateContactsCommon,
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

// CreateContactsCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation CreateContactsCommon for usage and error information.
func (c *VOLCOBSERVE) CreateContactsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateContactsCommonRequest(input)
	return out, req.Send()
}

// CreateContactsCommonWithContext is the same as CreateContactsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateContactsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) CreateContactsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateContactsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateContacts = "CreateContacts"

// CreateContactsRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateContacts operation. The "output" return
// value will be populated with the CreateContactsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateContactsCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateContactsCommon Send returns without error.
//
// See CreateContacts for more information on using the CreateContacts
// API call, and error handling.
//
//    // Example sending a request using the CreateContactsRequest method.
//    req, resp := client.CreateContactsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) CreateContactsRequest(input *CreateContactsInput) (req *request.Request, output *CreateContactsOutput) {
	op := &request.Operation{
		Name:       opCreateContacts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateContactsInput{}
	}

	output = &CreateContactsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateContacts API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation CreateContacts for usage and error information.
func (c *VOLCOBSERVE) CreateContacts(input *CreateContactsInput) (*CreateContactsOutput, error) {
	req, out := c.CreateContactsRequest(input)
	return out, req.Send()
}

// CreateContactsWithContext is the same as CreateContacts with the addition of
// the ability to pass a context and additional request options.
//
// See CreateContacts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) CreateContactsWithContext(ctx volcengine.Context, input *CreateContactsInput, opts ...request.Option) (*CreateContactsOutput, error) {
	req, out := c.CreateContactsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateContactsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Email is a required field
	Email *string `type:"string" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	Phone *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateContactsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateContactsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateContactsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateContactsInput"}
	if s.Email == nil {
		invalidParams.Add(request.NewErrParamRequired("Email"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEmail sets the Email field's value.
func (s *CreateContactsInput) SetEmail(v string) *CreateContactsInput {
	s.Email = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateContactsInput) SetName(v string) *CreateContactsInput {
	s.Name = &v
	return s
}

// SetPhone sets the Phone field's value.
func (s *CreateContactsInput) SetPhone(v string) *CreateContactsInput {
	s.Phone = &v
	return s
}

type CreateContactsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s CreateContactsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateContactsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *CreateContactsOutput) SetData(v []*string) *CreateContactsOutput {
	s.Data = v
	return s
}
