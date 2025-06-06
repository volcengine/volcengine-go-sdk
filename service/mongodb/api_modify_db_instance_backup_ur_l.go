// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceBackupURLCommon = "ModifyDBInstanceBackupURL"

// ModifyDBInstanceBackupURLCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceBackupURLCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceBackupURLCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceBackupURLCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceBackupURLCommon Send returns without error.
//
// See ModifyDBInstanceBackupURLCommon for more information on using the ModifyDBInstanceBackupURLCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceBackupURLCommonRequest method.
//    req, resp := client.ModifyDBInstanceBackupURLCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) ModifyDBInstanceBackupURLCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceBackupURLCommon,
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

// ModifyDBInstanceBackupURLCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation ModifyDBInstanceBackupURLCommon for usage and error information.
func (c *MONGODB) ModifyDBInstanceBackupURLCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceBackupURLCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceBackupURLCommonWithContext is the same as ModifyDBInstanceBackupURLCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceBackupURLCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) ModifyDBInstanceBackupURLCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceBackupURLCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceBackupURL = "ModifyDBInstanceBackupURL"

// ModifyDBInstanceBackupURLRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceBackupURL operation. The "output" return
// value will be populated with the ModifyDBInstanceBackupURLCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceBackupURLCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceBackupURLCommon Send returns without error.
//
// See ModifyDBInstanceBackupURL for more information on using the ModifyDBInstanceBackupURL
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceBackupURLRequest method.
//    req, resp := client.ModifyDBInstanceBackupURLRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) ModifyDBInstanceBackupURLRequest(input *ModifyDBInstanceBackupURLInput) (req *request.Request, output *ModifyDBInstanceBackupURLOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceBackupURL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceBackupURLInput{}
	}

	output = &ModifyDBInstanceBackupURLOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceBackupURL API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation ModifyDBInstanceBackupURL for usage and error information.
func (c *MONGODB) ModifyDBInstanceBackupURL(input *ModifyDBInstanceBackupURLInput) (*ModifyDBInstanceBackupURLOutput, error) {
	req, out := c.ModifyDBInstanceBackupURLRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceBackupURLWithContext is the same as ModifyDBInstanceBackupURL with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceBackupURL for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) ModifyDBInstanceBackupURLWithContext(ctx volcengine.Context, input *ModifyDBInstanceBackupURLInput, opts ...request.Option) (*ModifyDBInstanceBackupURLOutput, error) {
	req, out := c.ModifyDBInstanceBackupURLRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceBackupURLInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// BackupId is a required field
	BackupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceBackupURLInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceBackupURLInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceBackupURLInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceBackupURLInput"}
	if s.BackupId == nil {
		invalidParams.Add(request.NewErrParamRequired("BackupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupId sets the BackupId field's value.
func (s *ModifyDBInstanceBackupURLInput) SetBackupId(v string) *ModifyDBInstanceBackupURLInput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceBackupURLInput) SetInstanceId(v string) *ModifyDBInstanceBackupURLInput {
	s.InstanceId = &v
	return s
}

type ModifyDBInstanceBackupURLOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBInstanceBackupURLOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceBackupURLOutput) GoString() string {
	return s.String()
}
