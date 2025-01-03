// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDataFlowTasksCommon = "DescribeDataFlowTasks"

// DescribeDataFlowTasksCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDataFlowTasksCommon operation. The "output" return
// value will be populated with the DescribeDataFlowTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDataFlowTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDataFlowTasksCommon Send returns without error.
//
// See DescribeDataFlowTasksCommon for more information on using the DescribeDataFlowTasksCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDataFlowTasksCommonRequest method.
//    req, resp := client.DescribeDataFlowTasksCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DescribeDataFlowTasksCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDataFlowTasksCommon,
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

// DescribeDataFlowTasksCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DescribeDataFlowTasksCommon for usage and error information.
func (c *FILENAS) DescribeDataFlowTasksCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDataFlowTasksCommonRequest(input)
	return out, req.Send()
}

// DescribeDataFlowTasksCommonWithContext is the same as DescribeDataFlowTasksCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDataFlowTasksCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DescribeDataFlowTasksCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDataFlowTasksCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDataFlowTasks = "DescribeDataFlowTasks"

// DescribeDataFlowTasksRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDataFlowTasks operation. The "output" return
// value will be populated with the DescribeDataFlowTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDataFlowTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDataFlowTasksCommon Send returns without error.
//
// See DescribeDataFlowTasks for more information on using the DescribeDataFlowTasks
// API call, and error handling.
//
//    // Example sending a request using the DescribeDataFlowTasksRequest method.
//    req, resp := client.DescribeDataFlowTasksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DescribeDataFlowTasksRequest(input *DescribeDataFlowTasksInput) (req *request.Request, output *DescribeDataFlowTasksOutput) {
	op := &request.Operation{
		Name:       opDescribeDataFlowTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDataFlowTasksInput{}
	}

	output = &DescribeDataFlowTasksOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDataFlowTasks API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DescribeDataFlowTasks for usage and error information.
func (c *FILENAS) DescribeDataFlowTasks(input *DescribeDataFlowTasksInput) (*DescribeDataFlowTasksOutput, error) {
	req, out := c.DescribeDataFlowTasksRequest(input)
	return out, req.Send()
}

// DescribeDataFlowTasksWithContext is the same as DescribeDataFlowTasks with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDataFlowTasks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DescribeDataFlowTasksWithContext(ctx volcengine.Context, input *DescribeDataFlowTasksInput, opts ...request.Option) (*DescribeDataFlowTasksOutput, error) {
	req, out := c.DescribeDataFlowTasksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataFlowTaskForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BucketName *string `type:"string" json:",omitempty"`

	BucketPrefix *string `type:"string" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	DataFlowId *string `type:"string" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	EntryListFileKey *string `type:"string" json:",omitempty"`

	EntryListFileName *string `type:"string" json:",omitempty"`

	EntryListFileUrl *string `type:"string" json:",omitempty"`

	EvictPolicy *EvictPolicyForDescribeDataFlowTasksOutput `type:"structure" json:",omitempty"`

	ExportPolicy *ExportPolicyForDescribeDataFlowTasksOutput `type:"structure" json:",omitempty"`

	FileSystemId *string `type:"string" json:",omitempty"`

	FileSystemPath *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	ImportPolicy *ImportPolicyForDescribeDataFlowTasksOutput `type:"structure" json:",omitempty"`

	QueueExec *int64 `type:"int64" json:",omitempty"`

	QueueFailed *int64 `type:"int64" json:",omitempty"`

	QueueLen *int64 `type:"int64" json:",omitempty"`

	SameNameFilePolicy *string `type:"string" json:",omitempty" enum:"EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutput"`

	StartTime *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"ConvertEnumOfStatusForDescribeDataFlowTasksOutput"`

	Type *string `type:"string" json:",omitempty" enum:"ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutput"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataFlowTaskForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataFlowTaskForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetBucketName(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.BucketName = &v
	return s
}

// SetBucketPrefix sets the BucketPrefix field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetBucketPrefix(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.BucketPrefix = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetCreateTime(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.CreateTime = &v
	return s
}

// SetDataFlowId sets the DataFlowId field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetDataFlowId(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.DataFlowId = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetEndTime(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.EndTime = &v
	return s
}

// SetEntryListFileKey sets the EntryListFileKey field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetEntryListFileKey(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.EntryListFileKey = &v
	return s
}

// SetEntryListFileName sets the EntryListFileName field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetEntryListFileName(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.EntryListFileName = &v
	return s
}

// SetEntryListFileUrl sets the EntryListFileUrl field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetEntryListFileUrl(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.EntryListFileUrl = &v
	return s
}

// SetEvictPolicy sets the EvictPolicy field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetEvictPolicy(v *EvictPolicyForDescribeDataFlowTasksOutput) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.EvictPolicy = v
	return s
}

// SetExportPolicy sets the ExportPolicy field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetExportPolicy(v *ExportPolicyForDescribeDataFlowTasksOutput) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.ExportPolicy = v
	return s
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetFileSystemId(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.FileSystemId = &v
	return s
}

// SetFileSystemPath sets the FileSystemPath field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetFileSystemPath(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.FileSystemPath = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetId(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.Id = &v
	return s
}

// SetImportPolicy sets the ImportPolicy field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetImportPolicy(v *ImportPolicyForDescribeDataFlowTasksOutput) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.ImportPolicy = v
	return s
}

// SetQueueExec sets the QueueExec field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetQueueExec(v int64) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.QueueExec = &v
	return s
}

// SetQueueFailed sets the QueueFailed field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetQueueFailed(v int64) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.QueueFailed = &v
	return s
}

// SetQueueLen sets the QueueLen field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetQueueLen(v int64) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.QueueLen = &v
	return s
}

// SetSameNameFilePolicy sets the SameNameFilePolicy field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetSameNameFilePolicy(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.SameNameFilePolicy = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetStartTime(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetStatus(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetType(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.Type = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DataFlowTaskForDescribeDataFlowTasksOutput) SetUpdateTime(v string) *DataFlowTaskForDescribeDataFlowTasksOutput {
	s.UpdateTime = &v
	return s
}

type DescribeDataFlowTasksInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DataFlowIds *string `type:"string" json:",omitempty"`

	DataFlowNames *string `type:"string" json:",omitempty"`

	Ids *string `type:"string" json:",omitempty"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDataFlowTasksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDataFlowTasksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDataFlowTasksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDataFlowTasksInput"}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDataFlowIds sets the DataFlowIds field's value.
func (s *DescribeDataFlowTasksInput) SetDataFlowIds(v string) *DescribeDataFlowTasksInput {
	s.DataFlowIds = &v
	return s
}

// SetDataFlowNames sets the DataFlowNames field's value.
func (s *DescribeDataFlowTasksInput) SetDataFlowNames(v string) *DescribeDataFlowTasksInput {
	s.DataFlowNames = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *DescribeDataFlowTasksInput) SetIds(v string) *DescribeDataFlowTasksInput {
	s.Ids = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDataFlowTasksInput) SetPageNumber(v int32) *DescribeDataFlowTasksInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDataFlowTasksInput) SetPageSize(v int32) *DescribeDataFlowTasksInput {
	s.PageSize = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDataFlowTasksInput) SetStatus(v string) *DescribeDataFlowTasksInput {
	s.Status = &v
	return s
}

type DescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DataFlowTasks []*DataFlowTaskForDescribeDataFlowTasksOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetDataFlowTasks sets the DataFlowTasks field's value.
func (s *DescribeDataFlowTasksOutput) SetDataFlowTasks(v []*DataFlowTaskForDescribeDataFlowTasksOutput) *DescribeDataFlowTasksOutput {
	s.DataFlowTasks = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDataFlowTasksOutput) SetPageNumber(v int32) *DescribeDataFlowTasksOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDataFlowTasksOutput) SetPageSize(v int32) *DescribeDataFlowTasksOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeDataFlowTasksOutput) SetTotalCount(v int32) *DescribeDataFlowTasksOutput {
	s.TotalCount = &v
	return s
}

type DimensionForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Attr *string `type:"string" json:",omitempty" enum:"EnumOfAttrForDescribeDataFlowTasksOutput"`

	StaticValues []*StaticValueForDescribeDataFlowTasksOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetAttr sets the Attr field's value.
func (s *DimensionForDescribeDataFlowTasksOutput) SetAttr(v string) *DimensionForDescribeDataFlowTasksOutput {
	s.Attr = &v
	return s
}

// SetStaticValues sets the StaticValues field's value.
func (s *DimensionForDescribeDataFlowTasksOutput) SetStaticValues(v []*StaticValueForDescribeDataFlowTasksOutput) *DimensionForDescribeDataFlowTasksOutput {
	s.StaticValues = v
	return s
}

type EvictPolicyForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FilterInfo *FilterInfoForDescribeDataFlowTasksOutput `type:"structure" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeDataFlowTasksOutput"`

	Type *string `type:"string" json:",omitempty" enum:"EnumOfTypeForDescribeDataFlowTasksOutput"`
}

// String returns the string representation
func (s EvictPolicyForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EvictPolicyForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetFilterInfo sets the FilterInfo field's value.
func (s *EvictPolicyForDescribeDataFlowTasksOutput) SetFilterInfo(v *FilterInfoForDescribeDataFlowTasksOutput) *EvictPolicyForDescribeDataFlowTasksOutput {
	s.FilterInfo = v
	return s
}

// SetStatus sets the Status field's value.
func (s *EvictPolicyForDescribeDataFlowTasksOutput) SetStatus(v string) *EvictPolicyForDescribeDataFlowTasksOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *EvictPolicyForDescribeDataFlowTasksOutput) SetType(v string) *EvictPolicyForDescribeDataFlowTasksOutput {
	s.Type = &v
	return s
}

type ExportPolicyForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FilterInfo *FilterInfoForDescribeDataFlowTasksOutput `type:"structure" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeDataFlowTasksOutput"`

	Type *string `type:"string" json:",omitempty" enum:"EnumOfTypeForDescribeDataFlowTasksOutput"`
}

// String returns the string representation
func (s ExportPolicyForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ExportPolicyForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetFilterInfo sets the FilterInfo field's value.
func (s *ExportPolicyForDescribeDataFlowTasksOutput) SetFilterInfo(v *FilterInfoForDescribeDataFlowTasksOutput) *ExportPolicyForDescribeDataFlowTasksOutput {
	s.FilterInfo = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ExportPolicyForDescribeDataFlowTasksOutput) SetStatus(v string) *ExportPolicyForDescribeDataFlowTasksOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *ExportPolicyForDescribeDataFlowTasksOutput) SetType(v string) *ExportPolicyForDescribeDataFlowTasksOutput {
	s.Type = &v
	return s
}

type FilterInfoForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Dimensions []*DimensionForDescribeDataFlowTasksOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterInfoForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterInfoForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *FilterInfoForDescribeDataFlowTasksOutput) SetDimensions(v []*DimensionForDescribeDataFlowTasksOutput) *FilterInfoForDescribeDataFlowTasksOutput {
	s.Dimensions = v
	return s
}

type ImportPolicyForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FilterInfo *FilterInfoForDescribeDataFlowTasksOutput `type:"structure" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeDataFlowTasksOutput"`

	Type *string `type:"string" json:",omitempty" enum:"ConvertEnumOfTypeForDescribeDataFlowTasksOutput"`
}

// String returns the string representation
func (s ImportPolicyForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ImportPolicyForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetFilterInfo sets the FilterInfo field's value.
func (s *ImportPolicyForDescribeDataFlowTasksOutput) SetFilterInfo(v *FilterInfoForDescribeDataFlowTasksOutput) *ImportPolicyForDescribeDataFlowTasksOutput {
	s.FilterInfo = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ImportPolicyForDescribeDataFlowTasksOutput) SetStatus(v string) *ImportPolicyForDescribeDataFlowTasksOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *ImportPolicyForDescribeDataFlowTasksOutput) SetType(v string) *ImportPolicyForDescribeDataFlowTasksOutput {
	s.Type = &v
	return s
}

type StaticValueForDescribeDataFlowTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Relationship *string `type:"string" json:",omitempty" enum:"EnumOfRelationshipForDescribeDataFlowTasksOutput"`

	Unit *string `type:"string" json:",omitempty" enum:"EnumOfUnitForDescribeDataFlowTasksOutput"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StaticValueForDescribeDataFlowTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StaticValueForDescribeDataFlowTasksOutput) GoString() string {
	return s.String()
}

// SetRelationship sets the Relationship field's value.
func (s *StaticValueForDescribeDataFlowTasksOutput) SetRelationship(v string) *StaticValueForDescribeDataFlowTasksOutput {
	s.Relationship = &v
	return s
}

// SetUnit sets the Unit field's value.
func (s *StaticValueForDescribeDataFlowTasksOutput) SetUnit(v string) *StaticValueForDescribeDataFlowTasksOutput {
	s.Unit = &v
	return s
}

// SetValue sets the Value field's value.
func (s *StaticValueForDescribeDataFlowTasksOutput) SetValue(v string) *StaticValueForDescribeDataFlowTasksOutput {
	s.Value = &v
	return s
}

const (
	// ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutputImport is a ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutput enum value
	ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutputImport = "Import"

	// ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutputExport is a ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutput enum value
	ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutputExport = "Export"

	// ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutputEvict is a ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutput enum value
	ConvertConvertEnumOfTypeForDescribeDataFlowTasksOutputEvict = "Evict"
)

const (
	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputCreating is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputCreating = "Creating"

	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputRunning is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputRunning = "Running"

	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputFinished is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputFinished = "Finished"

	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputCanceling is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputCanceling = "Canceling"

	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputCanceled is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputCanceled = "Canceled"

	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputDeleting is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputDeleting = "Deleting"

	// ConvertEnumOfStatusForDescribeDataFlowTasksOutputError is a ConvertEnumOfStatusForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfStatusForDescribeDataFlowTasksOutputError = "Error"
)

const (
	// ConvertEnumOfTypeForDescribeDataFlowTasksOutputNone is a ConvertEnumOfTypeForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfTypeForDescribeDataFlowTasksOutputNone = "None"

	// ConvertEnumOfTypeForDescribeDataFlowTasksOutputMetadata is a ConvertEnumOfTypeForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfTypeForDescribeDataFlowTasksOutputMetadata = "Metadata"

	// ConvertEnumOfTypeForDescribeDataFlowTasksOutputMetaAndData is a ConvertEnumOfTypeForDescribeDataFlowTasksOutput enum value
	ConvertEnumOfTypeForDescribeDataFlowTasksOutputMetaAndData = "MetaAndData"
)

const (
	// EnumOfAttrForDescribeDataFlowTasksOutputFileSize is a EnumOfAttrForDescribeDataFlowTasksOutput enum value
	EnumOfAttrForDescribeDataFlowTasksOutputFileSize = "FileSize"

	// EnumOfAttrForDescribeDataFlowTasksOutputFileSuffix is a EnumOfAttrForDescribeDataFlowTasksOutput enum value
	EnumOfAttrForDescribeDataFlowTasksOutputFileSuffix = "FileSuffix"

	// EnumOfAttrForDescribeDataFlowTasksOutputSubDirectory is a EnumOfAttrForDescribeDataFlowTasksOutput enum value
	EnumOfAttrForDescribeDataFlowTasksOutputSubDirectory = "SubDirectory"

	// EnumOfAttrForDescribeDataFlowTasksOutputModifyTime is a EnumOfAttrForDescribeDataFlowTasksOutput enum value
	EnumOfAttrForDescribeDataFlowTasksOutputModifyTime = "ModifyTime"

	// EnumOfAttrForDescribeDataFlowTasksOutputAccessTime is a EnumOfAttrForDescribeDataFlowTasksOutput enum value
	EnumOfAttrForDescribeDataFlowTasksOutputAccessTime = "AccessTime"
)

const (
	// EnumOfRelationshipForDescribeDataFlowTasksOutputGe is a EnumOfRelationshipForDescribeDataFlowTasksOutput enum value
	EnumOfRelationshipForDescribeDataFlowTasksOutputGe = "GE"

	// EnumOfRelationshipForDescribeDataFlowTasksOutputCt is a EnumOfRelationshipForDescribeDataFlowTasksOutput enum value
	EnumOfRelationshipForDescribeDataFlowTasksOutputCt = "CT"

	// EnumOfRelationshipForDescribeDataFlowTasksOutputNc is a EnumOfRelationshipForDescribeDataFlowTasksOutput enum value
	EnumOfRelationshipForDescribeDataFlowTasksOutputNc = "NC"
)

const (
	// EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutputSkip is a EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutput enum value
	EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutputSkip = "Skip"

	// EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutputKeepLatest is a EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutput enum value
	EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutputKeepLatest = "KeepLatest"

	// EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutputOverWrite is a EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutput enum value
	EnumOfSameNameFilePolicyForDescribeDataFlowTasksOutputOverWrite = "OverWrite"
)

const (
	// EnumOfStatusForDescribeDataFlowTasksOutputStarting is a EnumOfStatusForDescribeDataFlowTasksOutput enum value
	EnumOfStatusForDescribeDataFlowTasksOutputStarting = "Starting"

	// EnumOfStatusForDescribeDataFlowTasksOutputRunning is a EnumOfStatusForDescribeDataFlowTasksOutput enum value
	EnumOfStatusForDescribeDataFlowTasksOutputRunning = "Running"

	// EnumOfStatusForDescribeDataFlowTasksOutputStopping is a EnumOfStatusForDescribeDataFlowTasksOutput enum value
	EnumOfStatusForDescribeDataFlowTasksOutputStopping = "Stopping"

	// EnumOfStatusForDescribeDataFlowTasksOutputStopped is a EnumOfStatusForDescribeDataFlowTasksOutput enum value
	EnumOfStatusForDescribeDataFlowTasksOutputStopped = "Stopped"

	// EnumOfStatusForDescribeDataFlowTasksOutputError is a EnumOfStatusForDescribeDataFlowTasksOutput enum value
	EnumOfStatusForDescribeDataFlowTasksOutputError = "Error"
)

const (
	// EnumOfTypeForDescribeDataFlowTasksOutputNone is a EnumOfTypeForDescribeDataFlowTasksOutput enum value
	EnumOfTypeForDescribeDataFlowTasksOutputNone = "None"

	// EnumOfTypeForDescribeDataFlowTasksOutputCustomize is a EnumOfTypeForDescribeDataFlowTasksOutput enum value
	EnumOfTypeForDescribeDataFlowTasksOutputCustomize = "Customize"
)

const (
	// EnumOfUnitForDescribeDataFlowTasksOutputHour is a EnumOfUnitForDescribeDataFlowTasksOutput enum value
	EnumOfUnitForDescribeDataFlowTasksOutputHour = "Hour"

	// EnumOfUnitForDescribeDataFlowTasksOutputDay is a EnumOfUnitForDescribeDataFlowTasksOutput enum value
	EnumOfUnitForDescribeDataFlowTasksOutputDay = "Day"

	// EnumOfUnitForDescribeDataFlowTasksOutputWeek is a EnumOfUnitForDescribeDataFlowTasksOutput enum value
	EnumOfUnitForDescribeDataFlowTasksOutputWeek = "Week"

	// EnumOfUnitForDescribeDataFlowTasksOutputMonth is a EnumOfUnitForDescribeDataFlowTasksOutput enum value
	EnumOfUnitForDescribeDataFlowTasksOutputMonth = "Month"

	// EnumOfUnitForDescribeDataFlowTasksOutputYear is a EnumOfUnitForDescribeDataFlowTasksOutput enum value
	EnumOfUnitForDescribeDataFlowTasksOutputYear = "Year"
)
