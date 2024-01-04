// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCenServiceRouteEntriesCommon = "DescribeCenServiceRouteEntries"

// DescribeCenServiceRouteEntriesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenServiceRouteEntriesCommon operation. The "output" return
// value will be populated with the DescribeCenServiceRouteEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenServiceRouteEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenServiceRouteEntriesCommon Send returns without error.
//
// See DescribeCenServiceRouteEntriesCommon for more information on using the DescribeCenServiceRouteEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenServiceRouteEntriesCommonRequest method.
//    req, resp := client.DescribeCenServiceRouteEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenServiceRouteEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCenServiceRouteEntriesCommon,
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

// DescribeCenServiceRouteEntriesCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenServiceRouteEntriesCommon for usage and error information.
func (c *CEN) DescribeCenServiceRouteEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCenServiceRouteEntriesCommonRequest(input)
	return out, req.Send()
}

// DescribeCenServiceRouteEntriesCommonWithContext is the same as DescribeCenServiceRouteEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenServiceRouteEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenServiceRouteEntriesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCenServiceRouteEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCenServiceRouteEntries = "DescribeCenServiceRouteEntries"

// DescribeCenServiceRouteEntriesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenServiceRouteEntries operation. The "output" return
// value will be populated with the DescribeCenServiceRouteEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenServiceRouteEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenServiceRouteEntriesCommon Send returns without error.
//
// See DescribeCenServiceRouteEntries for more information on using the DescribeCenServiceRouteEntries
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenServiceRouteEntriesRequest method.
//    req, resp := client.DescribeCenServiceRouteEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenServiceRouteEntriesRequest(input *DescribeCenServiceRouteEntriesInput) (req *request.Request, output *DescribeCenServiceRouteEntriesOutput) {
	op := &request.Operation{
		Name:       opDescribeCenServiceRouteEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCenServiceRouteEntriesInput{}
	}

	output = &DescribeCenServiceRouteEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCenServiceRouteEntries API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenServiceRouteEntries for usage and error information.
func (c *CEN) DescribeCenServiceRouteEntries(input *DescribeCenServiceRouteEntriesInput) (*DescribeCenServiceRouteEntriesOutput, error) {
	req, out := c.DescribeCenServiceRouteEntriesRequest(input)
	return out, req.Send()
}

// DescribeCenServiceRouteEntriesWithContext is the same as DescribeCenServiceRouteEntries with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenServiceRouteEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenServiceRouteEntriesWithContext(ctx volcengine.Context, input *DescribeCenServiceRouteEntriesInput, opts ...request.Option) (*DescribeCenServiceRouteEntriesOutput, error) {
	req, out := c.DescribeCenServiceRouteEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeCenServiceRouteEntriesInput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	CenRouteEntryIds []*string `type:"list"`

	DestinationCidrBlock *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ServiceRegionId *string `type:"string"`

	ServiceVpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeCenServiceRouteEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenServiceRouteEntriesInput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetCenId(v string) *DescribeCenServiceRouteEntriesInput {
	s.CenId = &v
	return s
}

// SetCenRouteEntryIds sets the CenRouteEntryIds field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetCenRouteEntryIds(v []*string) *DescribeCenServiceRouteEntriesInput {
	s.CenRouteEntryIds = v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetDestinationCidrBlock(v string) *DescribeCenServiceRouteEntriesInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetPageNumber(v int64) *DescribeCenServiceRouteEntriesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetPageSize(v int64) *DescribeCenServiceRouteEntriesInput {
	s.PageSize = &v
	return s
}

// SetServiceRegionId sets the ServiceRegionId field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetServiceRegionId(v string) *DescribeCenServiceRouteEntriesInput {
	s.ServiceRegionId = &v
	return s
}

// SetServiceVpcId sets the ServiceVpcId field's value.
func (s *DescribeCenServiceRouteEntriesInput) SetServiceVpcId(v string) *DescribeCenServiceRouteEntriesInput {
	s.ServiceVpcId = &v
	return s
}

type DescribeCenServiceRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ServiceRouteEntries []*ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenServiceRouteEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenServiceRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenServiceRouteEntriesOutput) SetPageNumber(v int64) *DescribeCenServiceRouteEntriesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenServiceRouteEntriesOutput) SetPageSize(v int64) *DescribeCenServiceRouteEntriesOutput {
	s.PageSize = &v
	return s
}

// SetServiceRouteEntries sets the ServiceRouteEntries field's value.
func (s *DescribeCenServiceRouteEntriesOutput) SetServiceRouteEntries(v []*ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) *DescribeCenServiceRouteEntriesOutput {
	s.ServiceRouteEntries = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCenServiceRouteEntriesOutput) SetTotalCount(v int64) *DescribeCenServiceRouteEntriesOutput {
	s.TotalCount = &v
	return s
}

type PublishToInstanceForDescribeCenServiceRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	InstanceRegionId *string `type:"string"`

	InstanceType *string `type:"string"`
}

// String returns the string representation
func (s PublishToInstanceForDescribeCenServiceRouteEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishToInstanceForDescribeCenServiceRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *PublishToInstanceForDescribeCenServiceRouteEntriesOutput) SetInstanceId(v string) *PublishToInstanceForDescribeCenServiceRouteEntriesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceRegionId sets the InstanceRegionId field's value.
func (s *PublishToInstanceForDescribeCenServiceRouteEntriesOutput) SetInstanceRegionId(v string) *PublishToInstanceForDescribeCenServiceRouteEntriesOutput {
	s.InstanceRegionId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *PublishToInstanceForDescribeCenServiceRouteEntriesOutput) SetInstanceType(v string) *PublishToInstanceForDescribeCenServiceRouteEntriesOutput {
	s.InstanceType = &v
	return s
}

type ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	PublishMode *string `type:"string"`

	PublishToInstances []*PublishToInstanceForDescribeCenServiceRouteEntriesOutput `type:"list"`

	ServiceRegionId *string `type:"string"`

	ServiceVpcId *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetCenId(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.CenId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetCreationTime(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetDescription(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetDestinationCidrBlock(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetPublishMode sets the PublishMode field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetPublishMode(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.PublishMode = &v
	return s
}

// SetPublishToInstances sets the PublishToInstances field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetPublishToInstances(v []*PublishToInstanceForDescribeCenServiceRouteEntriesOutput) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.PublishToInstances = v
	return s
}

// SetServiceRegionId sets the ServiceRegionId field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetServiceRegionId(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.ServiceRegionId = &v
	return s
}

// SetServiceVpcId sets the ServiceVpcId field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetServiceVpcId(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.ServiceVpcId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput) SetStatus(v string) *ServiceRouteEntryForDescribeCenServiceRouteEntriesOutput {
	s.Status = &v
	return s
}
