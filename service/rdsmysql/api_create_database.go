// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDatabaseCommon = "CreateDatabase"

// CreateDatabaseCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDatabaseCommon operation. The "output" return
// value will be populated with the CreateDatabaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDatabaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDatabaseCommon Send returns without error.
//
// See CreateDatabaseCommon for more information on using the CreateDatabaseCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDatabaseCommonRequest method.
//    req, resp := client.CreateDatabaseCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CreateDatabaseCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDatabaseCommon,
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

// CreateDatabaseCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateDatabaseCommon for usage and error information.
func (c *RDSMYSQL) CreateDatabaseCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDatabaseCommonRequest(input)
	return out, req.Send()
}

// CreateDatabaseCommonWithContext is the same as CreateDatabaseCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDatabaseCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CreateDatabaseCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDatabaseCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDatabase = "CreateDatabase"

// CreateDatabaseRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDatabase operation. The "output" return
// value will be populated with the CreateDatabaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDatabaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDatabaseCommon Send returns without error.
//
// See CreateDatabase for more information on using the CreateDatabase
// API call, and error handling.
//
//    // Example sending a request using the CreateDatabaseRequest method.
//    req, resp := client.CreateDatabaseRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CreateDatabaseRequest(input *CreateDatabaseInput) (req *request.Request, output *CreateDatabaseOutput) {
	op := &request.Operation{
		Name:       opCreateDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDatabaseInput{}
	}

	output = &CreateDatabaseOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDatabase API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateDatabase for usage and error information.
func (c *RDSMYSQL) CreateDatabase(input *CreateDatabaseInput) (*CreateDatabaseOutput, error) {
	req, out := c.CreateDatabaseRequest(input)
	return out, req.Send()
}

// CreateDatabaseWithContext is the same as CreateDatabase with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDatabase for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CreateDatabaseWithContext(ctx volcengine.Context, input *CreateDatabaseInput, opts ...request.Option) (*CreateDatabaseOutput, error) {
	req, out := c.CreateDatabaseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDatabaseInput struct {
	_ struct{} `type:"structure"`

	// CharacterSetName is a required field
	CharacterSetName *interface{} `type:"interface" required:"true"`

	DBName *string `min:"2" max:"64" type:"string"`

	// InstanceId is a required field
	InstanceId *interface{} `type:"interface" required:"true"`
}

// String returns the string representation
func (s CreateDatabaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDatabaseInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatabaseInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDatabaseInput"}
	if s.CharacterSetName == nil {
		invalidParams.Add(request.NewErrParamRequired("CharacterSetName"))
	}
	if s.DBName != nil && len(*s.DBName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("DBName", 2))
	}
	if s.DBName != nil && len(*s.DBName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("DBName", 64, *s.DBName))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCharacterSetName sets the CharacterSetName field's value.
func (s *CreateDatabaseInput) SetCharacterSetName(v interface{}) *CreateDatabaseInput {
	s.CharacterSetName = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *CreateDatabaseInput) SetDBName(v string) *CreateDatabaseInput {
	s.DBName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDatabaseInput) SetInstanceId(v interface{}) *CreateDatabaseInput {
	s.InstanceId = &v
	return s
}

type CreateDatabaseOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateDatabaseOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDatabaseOutput) GoString() string {
	return s.String()
}
