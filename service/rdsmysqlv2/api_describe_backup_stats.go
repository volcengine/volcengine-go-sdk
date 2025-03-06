// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBackupStatsCommon = "DescribeBackupStats"

// DescribeBackupStatsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupStatsCommon operation. The "output" return
// value will be populated with the DescribeBackupStatsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupStatsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupStatsCommon Send returns without error.
//
// See DescribeBackupStatsCommon for more information on using the DescribeBackupStatsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupStatsCommonRequest method.
//    req, resp := client.DescribeBackupStatsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBackupStatsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBackupStatsCommon,
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

// DescribeBackupStatsCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBackupStatsCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeBackupStatsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupStatsCommonRequest(input)
	return out, req.Send()
}

// DescribeBackupStatsCommonWithContext is the same as DescribeBackupStatsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupStatsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBackupStatsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupStatsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBackupStats = "DescribeBackupStats"

// DescribeBackupStatsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupStats operation. The "output" return
// value will be populated with the DescribeBackupStatsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupStatsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupStatsCommon Send returns without error.
//
// See DescribeBackupStats for more information on using the DescribeBackupStats
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupStatsRequest method.
//    req, resp := client.DescribeBackupStatsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBackupStatsRequest(input *DescribeBackupStatsInput) (req *request.Request, output *DescribeBackupStatsOutput) {
	op := &request.Operation{
		Name:       opDescribeBackupStats,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBackupStatsInput{}
	}

	output = &DescribeBackupStatsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBackupStats API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBackupStats for usage and error information.
func (c *RDSMYSQLV2) DescribeBackupStats(input *DescribeBackupStatsInput) (*DescribeBackupStatsOutput, error) {
	req, out := c.DescribeBackupStatsRequest(input)
	return out, req.Send()
}

// DescribeBackupStatsWithContext is the same as DescribeBackupStats with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupStats for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBackupStatsWithContext(ctx volcengine.Context, input *DescribeBackupStatsInput, opts ...request.Option) (*DescribeBackupStatsOutput, error) {
	req, out := c.DescribeBackupStatsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeBackupStatsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBackupStatsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupStatsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupStatsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBackupStatsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupStatsInput) SetInstanceId(v string) *DescribeBackupStatsInput {
	s.InstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeBackupStatsInput) SetProjectName(v string) *DescribeBackupStatsInput {
	s.ProjectName = &v
	return s
}

type DescribeBackupStatsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	UsageStats []*UsageStatForDescribeBackupStatsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBackupStatsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupStatsOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupStatsOutput) SetInstanceId(v string) *DescribeBackupStatsOutput {
	s.InstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeBackupStatsOutput) SetProjectName(v string) *DescribeBackupStatsOutput {
	s.ProjectName = &v
	return s
}

// SetUsageStats sets the UsageStats field's value.
func (s *DescribeBackupStatsOutput) SetUsageStats(v []*UsageStatForDescribeBackupStatsOutput) *DescribeBackupStatsOutput {
	s.UsageStats = v
	return s
}

type UsageStatForDescribeBackupStatsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Quantity *int64 `type:"int64" json:",omitempty"`

	StatItem *string `type:"string" json:",omitempty"`

	StatTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UsageStatForDescribeBackupStatsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageStatForDescribeBackupStatsOutput) GoString() string {
	return s.String()
}

// SetQuantity sets the Quantity field's value.
func (s *UsageStatForDescribeBackupStatsOutput) SetQuantity(v int64) *UsageStatForDescribeBackupStatsOutput {
	s.Quantity = &v
	return s
}

// SetStatItem sets the StatItem field's value.
func (s *UsageStatForDescribeBackupStatsOutput) SetStatItem(v string) *UsageStatForDescribeBackupStatsOutput {
	s.StatItem = &v
	return s
}

// SetStatTime sets the StatTime field's value.
func (s *UsageStatForDescribeBackupStatsOutput) SetStatTime(v string) *UsageStatForDescribeBackupStatsOutput {
	s.StatTime = &v
	return s
}
