// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCenAttachedInstancesCommon = "DescribeCenAttachedInstances"

// DescribeCenAttachedInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenAttachedInstancesCommon operation. The "output" return
// value will be populated with the DescribeCenAttachedInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenAttachedInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenAttachedInstancesCommon Send returns without error.
//
// See DescribeCenAttachedInstancesCommon for more information on using the DescribeCenAttachedInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenAttachedInstancesCommonRequest method.
//    req, resp := client.DescribeCenAttachedInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenAttachedInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCenAttachedInstancesCommon,
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

// DescribeCenAttachedInstancesCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenAttachedInstancesCommon for usage and error information.
func (c *CEN) DescribeCenAttachedInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCenAttachedInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeCenAttachedInstancesCommonWithContext is the same as DescribeCenAttachedInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenAttachedInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenAttachedInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCenAttachedInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCenAttachedInstances = "DescribeCenAttachedInstances"

// DescribeCenAttachedInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenAttachedInstances operation. The "output" return
// value will be populated with the DescribeCenAttachedInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenAttachedInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenAttachedInstancesCommon Send returns without error.
//
// See DescribeCenAttachedInstances for more information on using the DescribeCenAttachedInstances
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenAttachedInstancesRequest method.
//    req, resp := client.DescribeCenAttachedInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenAttachedInstancesRequest(input *DescribeCenAttachedInstancesInput) (req *request.Request, output *DescribeCenAttachedInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeCenAttachedInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCenAttachedInstancesInput{}
	}

	output = &DescribeCenAttachedInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCenAttachedInstances API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenAttachedInstances for usage and error information.
func (c *CEN) DescribeCenAttachedInstances(input *DescribeCenAttachedInstancesInput) (*DescribeCenAttachedInstancesOutput, error) {
	req, out := c.DescribeCenAttachedInstancesRequest(input)
	return out, req.Send()
}

// DescribeCenAttachedInstancesWithContext is the same as DescribeCenAttachedInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenAttachedInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenAttachedInstancesWithContext(ctx volcengine.Context, input *DescribeCenAttachedInstancesInput, opts ...request.Option) (*DescribeCenAttachedInstancesOutput, error) {
	req, out := c.DescribeCenAttachedInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachedInstanceForDescribeCenAttachedInstancesOutput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	CreationTime *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceOwnerId *string `type:"string"`

	InstanceRegionId *string `type:"string"`

	InstanceType *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s AttachedInstanceForDescribeCenAttachedInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachedInstanceForDescribeCenAttachedInstancesOutput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetCenId(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.CenId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetCreationTime(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.CreationTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetInstanceId(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceOwnerId sets the InstanceOwnerId field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetInstanceOwnerId(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.InstanceOwnerId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetInstanceRegionId(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetInstanceType(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.InstanceType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AttachedInstanceForDescribeCenAttachedInstancesOutput) SetStatus(v string) *AttachedInstanceForDescribeCenAttachedInstancesOutput {
	s.Status = &v
	return s
}

type DescribeCenAttachedInstancesInput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceOwnerId *string `type:"string"`

	InstanceRegionId *string `type:"string"`

	InstanceType *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenAttachedInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenAttachedInstancesInput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *DescribeCenAttachedInstancesInput) SetCenId(v string) *DescribeCenAttachedInstancesInput {
	s.CenId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeCenAttachedInstancesInput) SetInstanceId(v string) *DescribeCenAttachedInstancesInput {
	s.InstanceId = &v
	return s
}

// SetInstanceOwnerId sets the InstanceOwnerId field's value.
func (s *DescribeCenAttachedInstancesInput) SetInstanceOwnerId(v string) *DescribeCenAttachedInstancesInput {
	s.InstanceOwnerId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *DescribeCenAttachedInstancesInput) SetInstanceRegionId(v string) *DescribeCenAttachedInstancesInput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeCenAttachedInstancesInput) SetInstanceType(v string) *DescribeCenAttachedInstancesInput {
	s.InstanceType = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenAttachedInstancesInput) SetPageNumber(v int64) *DescribeCenAttachedInstancesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenAttachedInstancesInput) SetPageSize(v int64) *DescribeCenAttachedInstancesInput {
	s.PageSize = &v
	return s
}

type DescribeCenAttachedInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AttachedInstances []*AttachedInstanceForDescribeCenAttachedInstancesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenAttachedInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenAttachedInstancesOutput) GoString() string {
	return s.String()
}

// SetAttachedInstances sets the AttachedInstances field's value.
func (s *DescribeCenAttachedInstancesOutput) SetAttachedInstances(v []*AttachedInstanceForDescribeCenAttachedInstancesOutput) *DescribeCenAttachedInstancesOutput {
	s.AttachedInstances = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenAttachedInstancesOutput) SetPageNumber(v int64) *DescribeCenAttachedInstancesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenAttachedInstancesOutput) SetPageSize(v int64) *DescribeCenAttachedInstancesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCenAttachedInstancesOutput) SetTotalCount(v int64) *DescribeCenAttachedInstancesOutput {
	s.TotalCount = &v
	return s
}
