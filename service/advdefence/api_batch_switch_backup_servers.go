// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package advdefence

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opBatchSwitchBackupServersCommon = "BatchSwitchBackupServers"

// BatchSwitchBackupServersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the BatchSwitchBackupServersCommon operation. The "output" return
// value will be populated with the BatchSwitchBackupServersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchSwitchBackupServersCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchSwitchBackupServersCommon Send returns without error.
//
// See BatchSwitchBackupServersCommon for more information on using the BatchSwitchBackupServersCommon
// API call, and error handling.
//
//    // Example sending a request using the BatchSwitchBackupServersCommonRequest method.
//    req, resp := client.BatchSwitchBackupServersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) BatchSwitchBackupServersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opBatchSwitchBackupServersCommon,
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

// BatchSwitchBackupServersCommon API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation BatchSwitchBackupServersCommon for usage and error information.
func (c *ADVDEFENCE) BatchSwitchBackupServersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.BatchSwitchBackupServersCommonRequest(input)
	return out, req.Send()
}

// BatchSwitchBackupServersCommonWithContext is the same as BatchSwitchBackupServersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See BatchSwitchBackupServersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) BatchSwitchBackupServersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.BatchSwitchBackupServersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opBatchSwitchBackupServers = "BatchSwitchBackupServers"

// BatchSwitchBackupServersRequest generates a "volcengine/request.Request" representing the
// client's request for the BatchSwitchBackupServers operation. The "output" return
// value will be populated with the BatchSwitchBackupServersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchSwitchBackupServersCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchSwitchBackupServersCommon Send returns without error.
//
// See BatchSwitchBackupServers for more information on using the BatchSwitchBackupServers
// API call, and error handling.
//
//    // Example sending a request using the BatchSwitchBackupServersRequest method.
//    req, resp := client.BatchSwitchBackupServersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) BatchSwitchBackupServersRequest(input *BatchSwitchBackupServersInput) (req *request.Request, output *BatchSwitchBackupServersOutput) {
	op := &request.Operation{
		Name:       opBatchSwitchBackupServers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchSwitchBackupServersInput{}
	}

	output = &BatchSwitchBackupServersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// BatchSwitchBackupServers API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation BatchSwitchBackupServers for usage and error information.
func (c *ADVDEFENCE) BatchSwitchBackupServers(input *BatchSwitchBackupServersInput) (*BatchSwitchBackupServersOutput, error) {
	req, out := c.BatchSwitchBackupServersRequest(input)
	return out, req.Send()
}

// BatchSwitchBackupServersWithContext is the same as BatchSwitchBackupServers with the addition of
// the ability to pass a context and additional request options.
//
// See BatchSwitchBackupServers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) BatchSwitchBackupServersWithContext(ctx volcengine.Context, input *BatchSwitchBackupServersInput, opts ...request.Option) (*BatchSwitchBackupServersOutput, error) {
	req, out := c.BatchSwitchBackupServersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BatchSwitchBackupServersInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HostList []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s BatchSwitchBackupServersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchSwitchBackupServersInput) GoString() string {
	return s.String()
}

// SetHostList sets the HostList field's value.
func (s *BatchSwitchBackupServersInput) SetHostList(v []*string) *BatchSwitchBackupServersInput {
	s.HostList = v
	return s
}

type BatchSwitchBackupServersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s BatchSwitchBackupServersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchSwitchBackupServersOutput) GoString() string {
	return s.String()
}