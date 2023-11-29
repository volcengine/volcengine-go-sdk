// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribePrefixListAssociationsCommon = "DescribePrefixListAssociations"

// DescribePrefixListAssociationsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePrefixListAssociationsCommon operation. The "output" return
// value will be populated with the DescribePrefixListAssociationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePrefixListAssociationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePrefixListAssociationsCommon Send returns without error.
//
// See DescribePrefixListAssociationsCommon for more information on using the DescribePrefixListAssociationsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribePrefixListAssociationsCommonRequest method.
//    req, resp := client.DescribePrefixListAssociationsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribePrefixListAssociationsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribePrefixListAssociationsCommon,
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

// DescribePrefixListAssociationsCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribePrefixListAssociationsCommon for usage and error information.
func (c *VPC) DescribePrefixListAssociationsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribePrefixListAssociationsCommonRequest(input)
	return out, req.Send()
}

// DescribePrefixListAssociationsCommonWithContext is the same as DescribePrefixListAssociationsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePrefixListAssociationsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribePrefixListAssociationsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribePrefixListAssociationsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribePrefixListAssociations = "DescribePrefixListAssociations"

// DescribePrefixListAssociationsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePrefixListAssociations operation. The "output" return
// value will be populated with the DescribePrefixListAssociationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePrefixListAssociationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePrefixListAssociationsCommon Send returns without error.
//
// See DescribePrefixListAssociations for more information on using the DescribePrefixListAssociations
// API call, and error handling.
//
//    // Example sending a request using the DescribePrefixListAssociationsRequest method.
//    req, resp := client.DescribePrefixListAssociationsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribePrefixListAssociationsRequest(input *DescribePrefixListAssociationsInput) (req *request.Request, output *DescribePrefixListAssociationsOutput) {
	op := &request.Operation{
		Name:       opDescribePrefixListAssociations,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePrefixListAssociationsInput{}
	}

	output = &DescribePrefixListAssociationsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribePrefixListAssociations API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribePrefixListAssociations for usage and error information.
func (c *VPC) DescribePrefixListAssociations(input *DescribePrefixListAssociationsInput) (*DescribePrefixListAssociationsOutput, error) {
	req, out := c.DescribePrefixListAssociationsRequest(input)
	return out, req.Send()
}

// DescribePrefixListAssociationsWithContext is the same as DescribePrefixListAssociations with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePrefixListAssociations for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribePrefixListAssociationsWithContext(ctx volcengine.Context, input *DescribePrefixListAssociationsInput, opts ...request.Option) (*DescribePrefixListAssociationsOutput, error) {
	req, out := c.DescribePrefixListAssociationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribePrefixListAssociationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `min:"1" max:"100" type:"integer"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	// PrefixListId is a required field
	PrefixListId *string `type:"string" required:"true"`

	ResourceType *string `type:"string" enum:"EnumOfResourceTypeForDescribePrefixListAssociationsInput"`
}

// String returns the string representation
func (s DescribePrefixListAssociationsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePrefixListAssociationsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePrefixListAssociationsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribePrefixListAssociationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MaxResults != nil && *s.MaxResults > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("MaxResults", 100))
	}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}
	if s.PrefixListId == nil {
		invalidParams.Add(request.NewErrParamRequired("PrefixListId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribePrefixListAssociationsInput) SetMaxResults(v int64) *DescribePrefixListAssociationsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribePrefixListAssociationsInput) SetNextToken(v string) *DescribePrefixListAssociationsInput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribePrefixListAssociationsInput) SetPageNumber(v int64) *DescribePrefixListAssociationsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribePrefixListAssociationsInput) SetPageSize(v int64) *DescribePrefixListAssociationsInput {
	s.PageSize = &v
	return s
}

// SetPrefixListId sets the PrefixListId field's value.
func (s *DescribePrefixListAssociationsInput) SetPrefixListId(v string) *DescribePrefixListAssociationsInput {
	s.PrefixListId = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *DescribePrefixListAssociationsInput) SetResourceType(v string) *DescribePrefixListAssociationsInput {
	s.ResourceType = &v
	return s
}

type DescribePrefixListAssociationsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	PrefixListAssociations []*PrefixListAssociationForDescribePrefixListAssociationsOutput `type:"list"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribePrefixListAssociationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePrefixListAssociationsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribePrefixListAssociationsOutput) SetNextToken(v string) *DescribePrefixListAssociationsOutput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribePrefixListAssociationsOutput) SetPageNumber(v int64) *DescribePrefixListAssociationsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribePrefixListAssociationsOutput) SetPageSize(v int64) *DescribePrefixListAssociationsOutput {
	s.PageSize = &v
	return s
}

// SetPrefixListAssociations sets the PrefixListAssociations field's value.
func (s *DescribePrefixListAssociationsOutput) SetPrefixListAssociations(v []*PrefixListAssociationForDescribePrefixListAssociationsOutput) *DescribePrefixListAssociationsOutput {
	s.PrefixListAssociations = v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribePrefixListAssociationsOutput) SetRequestId(v string) *DescribePrefixListAssociationsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribePrefixListAssociationsOutput) SetTotalCount(v int64) *DescribePrefixListAssociationsOutput {
	s.TotalCount = &v
	return s
}

type PrefixListAssociationForDescribePrefixListAssociationsOutput struct {
	_ struct{} `type:"structure"`

	ResourceId *string `type:"string"`

	ResourceType *string `type:"string"`
}

// String returns the string representation
func (s PrefixListAssociationForDescribePrefixListAssociationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrefixListAssociationForDescribePrefixListAssociationsOutput) GoString() string {
	return s.String()
}

// SetResourceId sets the ResourceId field's value.
func (s *PrefixListAssociationForDescribePrefixListAssociationsOutput) SetResourceId(v string) *PrefixListAssociationForDescribePrefixListAssociationsOutput {
	s.ResourceId = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *PrefixListAssociationForDescribePrefixListAssociationsOutput) SetResourceType(v string) *PrefixListAssociationForDescribePrefixListAssociationsOutput {
	s.ResourceType = &v
	return s
}

const (
	// ResourceTypeForDescribePrefixListAssociationsInputVpcRouteTable is a EnumOfResourceTypeForDescribePrefixListAssociationsInput enum value
	ResourceTypeForDescribePrefixListAssociationsInputVpcRouteTable = "VpcRouteTable"

	// ResourceTypeForDescribePrefixListAssociationsInputVpcSecurityGroup is a EnumOfResourceTypeForDescribePrefixListAssociationsInput enum value
	ResourceTypeForDescribePrefixListAssociationsInputVpcSecurityGroup = "VpcSecurityGroup"
)
