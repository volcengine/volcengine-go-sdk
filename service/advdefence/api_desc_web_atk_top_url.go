// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package advdefence

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescWebAtkTopUrlCommon = "DescWebAtkTopUrl"

// DescWebAtkTopUrlCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescWebAtkTopUrlCommon operation. The "output" return
// value will be populated with the DescWebAtkTopUrlCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescWebAtkTopUrlCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescWebAtkTopUrlCommon Send returns without error.
//
// See DescWebAtkTopUrlCommon for more information on using the DescWebAtkTopUrlCommon
// API call, and error handling.
//
//    // Example sending a request using the DescWebAtkTopUrlCommonRequest method.
//    req, resp := client.DescWebAtkTopUrlCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) DescWebAtkTopUrlCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescWebAtkTopUrlCommon,
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

// DescWebAtkTopUrlCommon API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation DescWebAtkTopUrlCommon for usage and error information.
func (c *ADVDEFENCE) DescWebAtkTopUrlCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescWebAtkTopUrlCommonRequest(input)
	return out, req.Send()
}

// DescWebAtkTopUrlCommonWithContext is the same as DescWebAtkTopUrlCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescWebAtkTopUrlCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) DescWebAtkTopUrlCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescWebAtkTopUrlCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescWebAtkTopUrl = "DescWebAtkTopUrl"

// DescWebAtkTopUrlRequest generates a "volcengine/request.Request" representing the
// client's request for the DescWebAtkTopUrl operation. The "output" return
// value will be populated with the DescWebAtkTopUrlCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescWebAtkTopUrlCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescWebAtkTopUrlCommon Send returns without error.
//
// See DescWebAtkTopUrl for more information on using the DescWebAtkTopUrl
// API call, and error handling.
//
//    // Example sending a request using the DescWebAtkTopUrlRequest method.
//    req, resp := client.DescWebAtkTopUrlRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) DescWebAtkTopUrlRequest(input *DescWebAtkTopUrlInput) (req *request.Request, output *DescWebAtkTopUrlOutput) {
	op := &request.Operation{
		Name:       opDescWebAtkTopUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescWebAtkTopUrlInput{}
	}

	output = &DescWebAtkTopUrlOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescWebAtkTopUrl API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation DescWebAtkTopUrl for usage and error information.
func (c *ADVDEFENCE) DescWebAtkTopUrl(input *DescWebAtkTopUrlInput) (*DescWebAtkTopUrlOutput, error) {
	req, out := c.DescWebAtkTopUrlRequest(input)
	return out, req.Send()
}

// DescWebAtkTopUrlWithContext is the same as DescWebAtkTopUrl with the addition of
// the ability to pass a context and additional request options.
//
// See DescWebAtkTopUrl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) DescWebAtkTopUrlWithContext(ctx volcengine.Context, input *DescWebAtkTopUrlInput, opts ...request.Option) (*DescWebAtkTopUrlOutput, error) {
	req, out := c.DescWebAtkTopUrlRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescWebAtkTopUrlInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BeginTime *int32 `type:"int32" json:",omitempty"`

	EndTime *int32 `type:"int32" json:",omitempty"`

	Hosts []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescWebAtkTopUrlInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescWebAtkTopUrlInput) GoString() string {
	return s.String()
}

// SetBeginTime sets the BeginTime field's value.
func (s *DescWebAtkTopUrlInput) SetBeginTime(v int32) *DescWebAtkTopUrlInput {
	s.BeginTime = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescWebAtkTopUrlInput) SetEndTime(v int32) *DescWebAtkTopUrlInput {
	s.EndTime = &v
	return s
}

// SetHosts sets the Hosts field's value.
func (s *DescWebAtkTopUrlInput) SetHosts(v []*string) *DescWebAtkTopUrlInput {
	s.Hosts = v
	return s
}

type DescWebAtkTopUrlOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AttackCount *int32 `type:"int32" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Path *string `type:"string" json:",omitempty"`

	Percentage *float64 `type:"float" json:",omitempty"`
}

// String returns the string representation
func (s DescWebAtkTopUrlOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescWebAtkTopUrlOutput) GoString() string {
	return s.String()
}

// SetAttackCount sets the AttackCount field's value.
func (s *DescWebAtkTopUrlOutput) SetAttackCount(v int32) *DescWebAtkTopUrlOutput {
	s.AttackCount = &v
	return s
}

// SetHost sets the Host field's value.
func (s *DescWebAtkTopUrlOutput) SetHost(v string) *DescWebAtkTopUrlOutput {
	s.Host = &v
	return s
}

// SetPath sets the Path field's value.
func (s *DescWebAtkTopUrlOutput) SetPath(v string) *DescWebAtkTopUrlOutput {
	s.Path = &v
	return s
}

// SetPercentage sets the Percentage field's value.
func (s *DescWebAtkTopUrlOutput) SetPercentage(v float64) *DescWebAtkTopUrlOutput {
	s.Percentage = &v
	return s
}