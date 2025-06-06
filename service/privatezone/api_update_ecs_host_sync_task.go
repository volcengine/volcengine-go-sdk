// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatezone

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateEcsHostSyncTaskCommon = "UpdateEcsHostSyncTask"

// UpdateEcsHostSyncTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateEcsHostSyncTaskCommon operation. The "output" return
// value will be populated with the UpdateEcsHostSyncTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateEcsHostSyncTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateEcsHostSyncTaskCommon Send returns without error.
//
// See UpdateEcsHostSyncTaskCommon for more information on using the UpdateEcsHostSyncTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateEcsHostSyncTaskCommonRequest method.
//    req, resp := client.UpdateEcsHostSyncTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATEZONE) UpdateEcsHostSyncTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateEcsHostSyncTaskCommon,
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

// UpdateEcsHostSyncTaskCommon API operation for PRIVATE_ZONE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATE_ZONE's
// API operation UpdateEcsHostSyncTaskCommon for usage and error information.
func (c *PRIVATEZONE) UpdateEcsHostSyncTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateEcsHostSyncTaskCommonRequest(input)
	return out, req.Send()
}

// UpdateEcsHostSyncTaskCommonWithContext is the same as UpdateEcsHostSyncTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateEcsHostSyncTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATEZONE) UpdateEcsHostSyncTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateEcsHostSyncTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateEcsHostSyncTask = "UpdateEcsHostSyncTask"

// UpdateEcsHostSyncTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateEcsHostSyncTask operation. The "output" return
// value will be populated with the UpdateEcsHostSyncTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateEcsHostSyncTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateEcsHostSyncTaskCommon Send returns without error.
//
// See UpdateEcsHostSyncTask for more information on using the UpdateEcsHostSyncTask
// API call, and error handling.
//
//    // Example sending a request using the UpdateEcsHostSyncTaskRequest method.
//    req, resp := client.UpdateEcsHostSyncTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATEZONE) UpdateEcsHostSyncTaskRequest(input *UpdateEcsHostSyncTaskInput) (req *request.Request, output *UpdateEcsHostSyncTaskOutput) {
	op := &request.Operation{
		Name:       opUpdateEcsHostSyncTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateEcsHostSyncTaskInput{}
	}

	output = &UpdateEcsHostSyncTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateEcsHostSyncTask API operation for PRIVATE_ZONE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATE_ZONE's
// API operation UpdateEcsHostSyncTask for usage and error information.
func (c *PRIVATEZONE) UpdateEcsHostSyncTask(input *UpdateEcsHostSyncTaskInput) (*UpdateEcsHostSyncTaskOutput, error) {
	req, out := c.UpdateEcsHostSyncTaskRequest(input)
	return out, req.Send()
}

// UpdateEcsHostSyncTaskWithContext is the same as UpdateEcsHostSyncTask with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateEcsHostSyncTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATEZONE) UpdateEcsHostSyncTaskWithContext(ctx volcengine.Context, input *UpdateEcsHostSyncTaskInput, opts ...request.Option) (*UpdateEcsHostSyncTaskOutput, error) {
	req, out := c.UpdateEcsHostSyncTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EcsRegionForUpdateEcsHostSyncTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountID *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EcsRegionForUpdateEcsHostSyncTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EcsRegionForUpdateEcsHostSyncTaskInput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *EcsRegionForUpdateEcsHostSyncTaskInput) SetAccountID(v string) *EcsRegionForUpdateEcsHostSyncTaskInput {
	s.AccountID = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *EcsRegionForUpdateEcsHostSyncTaskInput) SetRegion(v string) *EcsRegionForUpdateEcsHostSyncTaskInput {
	s.Region = &v
	return s
}

type UpdateEcsHostSyncTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EcsRegions []*EcsRegionForUpdateEcsHostSyncTaskInput `type:"list" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	// ZID is a required field
	ZID *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateEcsHostSyncTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateEcsHostSyncTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEcsHostSyncTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateEcsHostSyncTaskInput"}
	if s.ZID == nil {
		invalidParams.Add(request.NewErrParamRequired("ZID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEcsRegions sets the EcsRegions field's value.
func (s *UpdateEcsHostSyncTaskInput) SetEcsRegions(v []*EcsRegionForUpdateEcsHostSyncTaskInput) *UpdateEcsHostSyncTaskInput {
	s.EcsRegions = v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateEcsHostSyncTaskInput) SetStatus(v string) *UpdateEcsHostSyncTaskInput {
	s.Status = &v
	return s
}

// SetZID sets the ZID field's value.
func (s *UpdateEcsHostSyncTaskInput) SetZID(v int64) *UpdateEcsHostSyncTaskInput {
	s.ZID = &v
	return s
}

type UpdateEcsHostSyncTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateEcsHostSyncTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateEcsHostSyncTaskOutput) GoString() string {
	return s.String()
}
