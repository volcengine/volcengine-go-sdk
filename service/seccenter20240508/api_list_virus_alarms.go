// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListVirusAlarmsCommon = "ListVirusAlarms"

// ListVirusAlarmsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListVirusAlarmsCommon operation. The "output" return
// value will be populated with the ListVirusAlarmsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListVirusAlarmsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListVirusAlarmsCommon Send returns without error.
//
// See ListVirusAlarmsCommon for more information on using the ListVirusAlarmsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListVirusAlarmsCommonRequest method.
//    req, resp := client.ListVirusAlarmsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListVirusAlarmsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListVirusAlarmsCommon,
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

// ListVirusAlarmsCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListVirusAlarmsCommon for usage and error information.
func (c *SECCENTER20240508) ListVirusAlarmsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListVirusAlarmsCommonRequest(input)
	return out, req.Send()
}

// ListVirusAlarmsCommonWithContext is the same as ListVirusAlarmsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListVirusAlarmsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListVirusAlarmsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListVirusAlarmsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListVirusAlarms = "ListVirusAlarms"

// ListVirusAlarmsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListVirusAlarms operation. The "output" return
// value will be populated with the ListVirusAlarmsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListVirusAlarmsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListVirusAlarmsCommon Send returns without error.
//
// See ListVirusAlarms for more information on using the ListVirusAlarms
// API call, and error handling.
//
//    // Example sending a request using the ListVirusAlarmsRequest method.
//    req, resp := client.ListVirusAlarmsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListVirusAlarmsRequest(input *ListVirusAlarmsInput) (req *request.Request, output *ListVirusAlarmsOutput) {
	op := &request.Operation{
		Name:       opListVirusAlarms,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVirusAlarmsInput{}
	}

	output = &ListVirusAlarmsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListVirusAlarms API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListVirusAlarms for usage and error information.
func (c *SECCENTER20240508) ListVirusAlarms(input *ListVirusAlarmsInput) (*ListVirusAlarmsOutput, error) {
	req, out := c.ListVirusAlarmsRequest(input)
	return out, req.Send()
}

// ListVirusAlarmsWithContext is the same as ListVirusAlarms with the addition of
// the ability to pass a context and additional request options.
//
// See ListVirusAlarms for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListVirusAlarmsWithContext(ctx volcengine.Context, input *ListVirusAlarmsInput, opts ...request.Option) (*ListVirusAlarmsOutput, error) {
	req, out := c.ListVirusAlarmsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttributionListForListVirusAlarmsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AttributionListForListVirusAlarmsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttributionListForListVirusAlarmsOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *AttributionListForListVirusAlarmsOutput) SetType(v string) *AttributionListForListVirusAlarmsOutput {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *AttributionListForListVirusAlarmsOutput) SetValue(v string) *AttributionListForListVirusAlarmsOutput {
	s.Value = &v
	return s
}

type ClusterForListVirusAlarmsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClusterArea *string `type:"string" json:",omitempty"`

	ClusterID *string `type:"string" json:",omitempty"`

	ClusterName *string `type:"string" json:",omitempty"`

	ClusterTags []*string `type:"list" json:",omitempty"`

	RuleType1st *string `type:"string" json:",omitempty"`

	RuleType2nd *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ClusterForListVirusAlarmsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ClusterForListVirusAlarmsOutput) GoString() string {
	return s.String()
}

// SetClusterArea sets the ClusterArea field's value.
func (s *ClusterForListVirusAlarmsOutput) SetClusterArea(v string) *ClusterForListVirusAlarmsOutput {
	s.ClusterArea = &v
	return s
}

// SetClusterID sets the ClusterID field's value.
func (s *ClusterForListVirusAlarmsOutput) SetClusterID(v string) *ClusterForListVirusAlarmsOutput {
	s.ClusterID = &v
	return s
}

// SetClusterName sets the ClusterName field's value.
func (s *ClusterForListVirusAlarmsOutput) SetClusterName(v string) *ClusterForListVirusAlarmsOutput {
	s.ClusterName = &v
	return s
}

// SetClusterTags sets the ClusterTags field's value.
func (s *ClusterForListVirusAlarmsOutput) SetClusterTags(v []*string) *ClusterForListVirusAlarmsOutput {
	s.ClusterTags = v
	return s
}

// SetRuleType1st sets the RuleType1st field's value.
func (s *ClusterForListVirusAlarmsOutput) SetRuleType1st(v string) *ClusterForListVirusAlarmsOutput {
	s.RuleType1st = &v
	return s
}

// SetRuleType2nd sets the RuleType2nd field's value.
func (s *ClusterForListVirusAlarmsOutput) SetRuleType2nd(v string) *ClusterForListVirusAlarmsOutput {
	s.RuleType2nd = &v
	return s
}

type DataForListVirusAlarmsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentGroup *string `type:"string" json:",omitempty"`

	AgentID *string `type:"string" json:",omitempty"`

	AlarmHandleResult *int32 `type:"int32" json:",omitempty"`

	AlarmHostname *string `type:"string" json:",omitempty"`

	AlarmID *string `type:"string" json:",omitempty"`

	AlarmTime *int32 `type:"int32" json:",omitempty"`

	AlertTags []*string `type:"list" json:",omitempty"`

	Args []*string `type:"list" json:",omitempty"`

	ArgsList []*string `type:"list" json:",omitempty"`

	ArgvList []*string `type:"list" json:",omitempty"`

	AttributionList []*AttributionListForListVirusAlarmsOutput `type:"list" json:",omitempty"`

	Class *string `type:"string" json:",omitempty"`

	Cluster *ClusterForListVirusAlarmsOutput `type:"structure" json:",omitempty"`

	DataType *string `type:"string" json:",omitempty"`

	ErrorReason *string `type:"string" json:",omitempty"`

	EventID *string `type:"string" json:",omitempty"`

	EventName *string `type:"string" json:",omitempty"`

	Exe *string `type:"string" json:",omitempty"`

	FileDownloadable *bool `type:"boolean" json:",omitempty"`

	FileHash *string `type:"string" json:",omitempty"`

	FilePath *string `type:"string" json:",omitempty"`

	GroupPath *string `type:"string" json:",omitempty"`

	HandleTime *int32 `type:"int32" json:",omitempty"`

	Host *HostForListVirusAlarmsOutput `type:"structure" json:",omitempty"`

	ImageName *string `type:"string" json:",omitempty"`

	InDocker *string `type:"string" json:",omitempty"`

	LLMAnalysisResult *string `type:"string" json:",omitempty"`

	LLMProcessed *bool `type:"boolean" json:",omitempty"`

	Level *string `type:"string" json:",omitempty"`

	MlpInstanceID *string `type:"string" json:",omitempty"`

	MlpPrivateIP *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	NsPid *string `type:"string" json:",omitempty"`

	OsType *string `type:"string" json:",omitempty"`

	Pid *string `type:"string" json:",omitempty"`

	ProbeHook *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	Sid *string `type:"string" json:",omitempty"`

	StackTraceFormat *string `type:"string" json:",omitempty"`

	StackTraceHash *string `type:"string" json:",omitempty"`

	Status *int32 `type:"int32" json:",omitempty"`

	TraceID *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListVirusAlarmsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListVirusAlarmsOutput) GoString() string {
	return s.String()
}

// SetAgentGroup sets the AgentGroup field's value.
func (s *DataForListVirusAlarmsOutput) SetAgentGroup(v string) *DataForListVirusAlarmsOutput {
	s.AgentGroup = &v
	return s
}

// SetAgentID sets the AgentID field's value.
func (s *DataForListVirusAlarmsOutput) SetAgentID(v string) *DataForListVirusAlarmsOutput {
	s.AgentID = &v
	return s
}

// SetAlarmHandleResult sets the AlarmHandleResult field's value.
func (s *DataForListVirusAlarmsOutput) SetAlarmHandleResult(v int32) *DataForListVirusAlarmsOutput {
	s.AlarmHandleResult = &v
	return s
}

// SetAlarmHostname sets the AlarmHostname field's value.
func (s *DataForListVirusAlarmsOutput) SetAlarmHostname(v string) *DataForListVirusAlarmsOutput {
	s.AlarmHostname = &v
	return s
}

// SetAlarmID sets the AlarmID field's value.
func (s *DataForListVirusAlarmsOutput) SetAlarmID(v string) *DataForListVirusAlarmsOutput {
	s.AlarmID = &v
	return s
}

// SetAlarmTime sets the AlarmTime field's value.
func (s *DataForListVirusAlarmsOutput) SetAlarmTime(v int32) *DataForListVirusAlarmsOutput {
	s.AlarmTime = &v
	return s
}

// SetAlertTags sets the AlertTags field's value.
func (s *DataForListVirusAlarmsOutput) SetAlertTags(v []*string) *DataForListVirusAlarmsOutput {
	s.AlertTags = v
	return s
}

// SetArgs sets the Args field's value.
func (s *DataForListVirusAlarmsOutput) SetArgs(v []*string) *DataForListVirusAlarmsOutput {
	s.Args = v
	return s
}

// SetArgsList sets the ArgsList field's value.
func (s *DataForListVirusAlarmsOutput) SetArgsList(v []*string) *DataForListVirusAlarmsOutput {
	s.ArgsList = v
	return s
}

// SetArgvList sets the ArgvList field's value.
func (s *DataForListVirusAlarmsOutput) SetArgvList(v []*string) *DataForListVirusAlarmsOutput {
	s.ArgvList = v
	return s
}

// SetAttributionList sets the AttributionList field's value.
func (s *DataForListVirusAlarmsOutput) SetAttributionList(v []*AttributionListForListVirusAlarmsOutput) *DataForListVirusAlarmsOutput {
	s.AttributionList = v
	return s
}

// SetClass sets the Class field's value.
func (s *DataForListVirusAlarmsOutput) SetClass(v string) *DataForListVirusAlarmsOutput {
	s.Class = &v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *DataForListVirusAlarmsOutput) SetCluster(v *ClusterForListVirusAlarmsOutput) *DataForListVirusAlarmsOutput {
	s.Cluster = v
	return s
}

// SetDataType sets the DataType field's value.
func (s *DataForListVirusAlarmsOutput) SetDataType(v string) *DataForListVirusAlarmsOutput {
	s.DataType = &v
	return s
}

// SetErrorReason sets the ErrorReason field's value.
func (s *DataForListVirusAlarmsOutput) SetErrorReason(v string) *DataForListVirusAlarmsOutput {
	s.ErrorReason = &v
	return s
}

// SetEventID sets the EventID field's value.
func (s *DataForListVirusAlarmsOutput) SetEventID(v string) *DataForListVirusAlarmsOutput {
	s.EventID = &v
	return s
}

// SetEventName sets the EventName field's value.
func (s *DataForListVirusAlarmsOutput) SetEventName(v string) *DataForListVirusAlarmsOutput {
	s.EventName = &v
	return s
}

// SetExe sets the Exe field's value.
func (s *DataForListVirusAlarmsOutput) SetExe(v string) *DataForListVirusAlarmsOutput {
	s.Exe = &v
	return s
}

// SetFileDownloadable sets the FileDownloadable field's value.
func (s *DataForListVirusAlarmsOutput) SetFileDownloadable(v bool) *DataForListVirusAlarmsOutput {
	s.FileDownloadable = &v
	return s
}

// SetFileHash sets the FileHash field's value.
func (s *DataForListVirusAlarmsOutput) SetFileHash(v string) *DataForListVirusAlarmsOutput {
	s.FileHash = &v
	return s
}

// SetFilePath sets the FilePath field's value.
func (s *DataForListVirusAlarmsOutput) SetFilePath(v string) *DataForListVirusAlarmsOutput {
	s.FilePath = &v
	return s
}

// SetGroupPath sets the GroupPath field's value.
func (s *DataForListVirusAlarmsOutput) SetGroupPath(v string) *DataForListVirusAlarmsOutput {
	s.GroupPath = &v
	return s
}

// SetHandleTime sets the HandleTime field's value.
func (s *DataForListVirusAlarmsOutput) SetHandleTime(v int32) *DataForListVirusAlarmsOutput {
	s.HandleTime = &v
	return s
}

// SetHost sets the Host field's value.
func (s *DataForListVirusAlarmsOutput) SetHost(v *HostForListVirusAlarmsOutput) *DataForListVirusAlarmsOutput {
	s.Host = v
	return s
}

// SetImageName sets the ImageName field's value.
func (s *DataForListVirusAlarmsOutput) SetImageName(v string) *DataForListVirusAlarmsOutput {
	s.ImageName = &v
	return s
}

// SetInDocker sets the InDocker field's value.
func (s *DataForListVirusAlarmsOutput) SetInDocker(v string) *DataForListVirusAlarmsOutput {
	s.InDocker = &v
	return s
}

// SetLLMAnalysisResult sets the LLMAnalysisResult field's value.
func (s *DataForListVirusAlarmsOutput) SetLLMAnalysisResult(v string) *DataForListVirusAlarmsOutput {
	s.LLMAnalysisResult = &v
	return s
}

// SetLLMProcessed sets the LLMProcessed field's value.
func (s *DataForListVirusAlarmsOutput) SetLLMProcessed(v bool) *DataForListVirusAlarmsOutput {
	s.LLMProcessed = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *DataForListVirusAlarmsOutput) SetLevel(v string) *DataForListVirusAlarmsOutput {
	s.Level = &v
	return s
}

// SetMlpInstanceID sets the MlpInstanceID field's value.
func (s *DataForListVirusAlarmsOutput) SetMlpInstanceID(v string) *DataForListVirusAlarmsOutput {
	s.MlpInstanceID = &v
	return s
}

// SetMlpPrivateIP sets the MlpPrivateIP field's value.
func (s *DataForListVirusAlarmsOutput) SetMlpPrivateIP(v string) *DataForListVirusAlarmsOutput {
	s.MlpPrivateIP = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListVirusAlarmsOutput) SetName(v string) *DataForListVirusAlarmsOutput {
	s.Name = &v
	return s
}

// SetNsPid sets the NsPid field's value.
func (s *DataForListVirusAlarmsOutput) SetNsPid(v string) *DataForListVirusAlarmsOutput {
	s.NsPid = &v
	return s
}

// SetOsType sets the OsType field's value.
func (s *DataForListVirusAlarmsOutput) SetOsType(v string) *DataForListVirusAlarmsOutput {
	s.OsType = &v
	return s
}

// SetPid sets the Pid field's value.
func (s *DataForListVirusAlarmsOutput) SetPid(v string) *DataForListVirusAlarmsOutput {
	s.Pid = &v
	return s
}

// SetProbeHook sets the ProbeHook field's value.
func (s *DataForListVirusAlarmsOutput) SetProbeHook(v string) *DataForListVirusAlarmsOutput {
	s.ProbeHook = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForListVirusAlarmsOutput) SetRegion(v string) *DataForListVirusAlarmsOutput {
	s.Region = &v
	return s
}

// SetSid sets the Sid field's value.
func (s *DataForListVirusAlarmsOutput) SetSid(v string) *DataForListVirusAlarmsOutput {
	s.Sid = &v
	return s
}

// SetStackTraceFormat sets the StackTraceFormat field's value.
func (s *DataForListVirusAlarmsOutput) SetStackTraceFormat(v string) *DataForListVirusAlarmsOutput {
	s.StackTraceFormat = &v
	return s
}

// SetStackTraceHash sets the StackTraceHash field's value.
func (s *DataForListVirusAlarmsOutput) SetStackTraceHash(v string) *DataForListVirusAlarmsOutput {
	s.StackTraceHash = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForListVirusAlarmsOutput) SetStatus(v int32) *DataForListVirusAlarmsOutput {
	s.Status = &v
	return s
}

// SetTraceID sets the TraceID field's value.
func (s *DataForListVirusAlarmsOutput) SetTraceID(v string) *DataForListVirusAlarmsOutput {
	s.TraceID = &v
	return s
}

// SetType sets the Type field's value.
func (s *DataForListVirusAlarmsOutput) SetType(v string) *DataForListVirusAlarmsOutput {
	s.Type = &v
	return s
}

type HostForListVirusAlarmsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentID *string `type:"string" json:",omitempty"`

	AgentTags []*string `type:"list" json:",omitempty"`

	Hostname *string `type:"string" json:",omitempty"`

	InnerIPList []*string `type:"list" json:",omitempty"`

	OuterIPList []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s HostForListVirusAlarmsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HostForListVirusAlarmsOutput) GoString() string {
	return s.String()
}

// SetAgentID sets the AgentID field's value.
func (s *HostForListVirusAlarmsOutput) SetAgentID(v string) *HostForListVirusAlarmsOutput {
	s.AgentID = &v
	return s
}

// SetAgentTags sets the AgentTags field's value.
func (s *HostForListVirusAlarmsOutput) SetAgentTags(v []*string) *HostForListVirusAlarmsOutput {
	s.AgentTags = v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *HostForListVirusAlarmsOutput) SetHostname(v string) *HostForListVirusAlarmsOutput {
	s.Hostname = &v
	return s
}

// SetInnerIPList sets the InnerIPList field's value.
func (s *HostForListVirusAlarmsOutput) SetInnerIPList(v []*string) *HostForListVirusAlarmsOutput {
	s.InnerIPList = v
	return s
}

// SetOuterIPList sets the OuterIPList field's value.
func (s *HostForListVirusAlarmsOutput) SetOuterIPList(v []*string) *HostForListVirusAlarmsOutput {
	s.OuterIPList = v
	return s
}

type ListVirusAlarmsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentGroup *string `type:"string" json:",omitempty"`

	AgentGroupList []*string `type:"list" json:",omitempty"`

	AgentID *string `type:"string" json:",omitempty"`

	AgentIDList []*string `type:"list" json:",omitempty"`

	AgentTags []*string `type:"list" json:",omitempty"`

	AlarmHandleResultList []*int32 `type:"list" json:",omitempty"`

	AlarmID *string `type:"string" json:",omitempty"`

	AlertTags []*string `type:"list" json:",omitempty"`

	CloudProviders []*string `type:"list" json:",omitempty"`

	ClusterID *string `type:"string" json:",omitempty"`

	ClusterName *string `type:"string" json:",omitempty"`

	ClusterRegion *string `type:"string" json:",omitempty"`

	ClusterTags []*string `type:"list" json:",omitempty"`

	ContainerID *string `type:"string" json:",omitempty"`

	ContainerName *string `type:"string" json:",omitempty"`

	DataType *string `type:"string" json:",omitempty"`

	EventID *string `type:"string" json:",omitempty"`

	EventName *string `type:"string" json:",omitempty"`

	EventReason *string `type:"string" json:",omitempty"`

	Exe *string `type:"string" json:",omitempty"`

	FileHash *string `type:"string" json:",omitempty"`

	FilePath *string `type:"string" json:",omitempty"`

	Hostname *string `type:"string" json:",omitempty"`

	IP *string `type:"string" json:",omitempty"`

	ImageName *string `type:"string" json:",omitempty"`

	LeafGroupIDs []*string `type:"list" json:",omitempty"`

	LevelList []*string `type:"list" json:",omitempty"`

	MlpInstanceID *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	ProbeHook *string `type:"string" json:",omitempty"`

	RaspArgv *string `type:"string" json:",omitempty"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`

	Status []*int32 `type:"list" json:",omitempty"`

	TaskID *string `type:"string" json:",omitempty"`

	TimeEnd *int32 `type:"int32" json:",omitempty"`

	TimeStart *int32 `type:"int32" json:",omitempty"`

	TopGroupID *string `type:"string" json:",omitempty"`

	Type []*string `type:"list" json:",omitempty"`

	VirusType *string `type:"string" json:",omitempty"`

	WhiteListID *string `type:"string" json:",omitempty"`

	WhiteListName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListVirusAlarmsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListVirusAlarmsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVirusAlarmsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListVirusAlarmsInput"}
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

// SetAgentGroup sets the AgentGroup field's value.
func (s *ListVirusAlarmsInput) SetAgentGroup(v string) *ListVirusAlarmsInput {
	s.AgentGroup = &v
	return s
}

// SetAgentGroupList sets the AgentGroupList field's value.
func (s *ListVirusAlarmsInput) SetAgentGroupList(v []*string) *ListVirusAlarmsInput {
	s.AgentGroupList = v
	return s
}

// SetAgentID sets the AgentID field's value.
func (s *ListVirusAlarmsInput) SetAgentID(v string) *ListVirusAlarmsInput {
	s.AgentID = &v
	return s
}

// SetAgentIDList sets the AgentIDList field's value.
func (s *ListVirusAlarmsInput) SetAgentIDList(v []*string) *ListVirusAlarmsInput {
	s.AgentIDList = v
	return s
}

// SetAgentTags sets the AgentTags field's value.
func (s *ListVirusAlarmsInput) SetAgentTags(v []*string) *ListVirusAlarmsInput {
	s.AgentTags = v
	return s
}

// SetAlarmHandleResultList sets the AlarmHandleResultList field's value.
func (s *ListVirusAlarmsInput) SetAlarmHandleResultList(v []*int32) *ListVirusAlarmsInput {
	s.AlarmHandleResultList = v
	return s
}

// SetAlarmID sets the AlarmID field's value.
func (s *ListVirusAlarmsInput) SetAlarmID(v string) *ListVirusAlarmsInput {
	s.AlarmID = &v
	return s
}

// SetAlertTags sets the AlertTags field's value.
func (s *ListVirusAlarmsInput) SetAlertTags(v []*string) *ListVirusAlarmsInput {
	s.AlertTags = v
	return s
}

// SetCloudProviders sets the CloudProviders field's value.
func (s *ListVirusAlarmsInput) SetCloudProviders(v []*string) *ListVirusAlarmsInput {
	s.CloudProviders = v
	return s
}

// SetClusterID sets the ClusterID field's value.
func (s *ListVirusAlarmsInput) SetClusterID(v string) *ListVirusAlarmsInput {
	s.ClusterID = &v
	return s
}

// SetClusterName sets the ClusterName field's value.
func (s *ListVirusAlarmsInput) SetClusterName(v string) *ListVirusAlarmsInput {
	s.ClusterName = &v
	return s
}

// SetClusterRegion sets the ClusterRegion field's value.
func (s *ListVirusAlarmsInput) SetClusterRegion(v string) *ListVirusAlarmsInput {
	s.ClusterRegion = &v
	return s
}

// SetClusterTags sets the ClusterTags field's value.
func (s *ListVirusAlarmsInput) SetClusterTags(v []*string) *ListVirusAlarmsInput {
	s.ClusterTags = v
	return s
}

// SetContainerID sets the ContainerID field's value.
func (s *ListVirusAlarmsInput) SetContainerID(v string) *ListVirusAlarmsInput {
	s.ContainerID = &v
	return s
}

// SetContainerName sets the ContainerName field's value.
func (s *ListVirusAlarmsInput) SetContainerName(v string) *ListVirusAlarmsInput {
	s.ContainerName = &v
	return s
}

// SetDataType sets the DataType field's value.
func (s *ListVirusAlarmsInput) SetDataType(v string) *ListVirusAlarmsInput {
	s.DataType = &v
	return s
}

// SetEventID sets the EventID field's value.
func (s *ListVirusAlarmsInput) SetEventID(v string) *ListVirusAlarmsInput {
	s.EventID = &v
	return s
}

// SetEventName sets the EventName field's value.
func (s *ListVirusAlarmsInput) SetEventName(v string) *ListVirusAlarmsInput {
	s.EventName = &v
	return s
}

// SetEventReason sets the EventReason field's value.
func (s *ListVirusAlarmsInput) SetEventReason(v string) *ListVirusAlarmsInput {
	s.EventReason = &v
	return s
}

// SetExe sets the Exe field's value.
func (s *ListVirusAlarmsInput) SetExe(v string) *ListVirusAlarmsInput {
	s.Exe = &v
	return s
}

// SetFileHash sets the FileHash field's value.
func (s *ListVirusAlarmsInput) SetFileHash(v string) *ListVirusAlarmsInput {
	s.FileHash = &v
	return s
}

// SetFilePath sets the FilePath field's value.
func (s *ListVirusAlarmsInput) SetFilePath(v string) *ListVirusAlarmsInput {
	s.FilePath = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *ListVirusAlarmsInput) SetHostname(v string) *ListVirusAlarmsInput {
	s.Hostname = &v
	return s
}

// SetIP sets the IP field's value.
func (s *ListVirusAlarmsInput) SetIP(v string) *ListVirusAlarmsInput {
	s.IP = &v
	return s
}

// SetImageName sets the ImageName field's value.
func (s *ListVirusAlarmsInput) SetImageName(v string) *ListVirusAlarmsInput {
	s.ImageName = &v
	return s
}

// SetLeafGroupIDs sets the LeafGroupIDs field's value.
func (s *ListVirusAlarmsInput) SetLeafGroupIDs(v []*string) *ListVirusAlarmsInput {
	s.LeafGroupIDs = v
	return s
}

// SetLevelList sets the LevelList field's value.
func (s *ListVirusAlarmsInput) SetLevelList(v []*string) *ListVirusAlarmsInput {
	s.LevelList = v
	return s
}

// SetMlpInstanceID sets the MlpInstanceID field's value.
func (s *ListVirusAlarmsInput) SetMlpInstanceID(v string) *ListVirusAlarmsInput {
	s.MlpInstanceID = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListVirusAlarmsInput) SetName(v string) *ListVirusAlarmsInput {
	s.Name = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListVirusAlarmsInput) SetPageNumber(v int32) *ListVirusAlarmsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListVirusAlarmsInput) SetPageSize(v int32) *ListVirusAlarmsInput {
	s.PageSize = &v
	return s
}

// SetProbeHook sets the ProbeHook field's value.
func (s *ListVirusAlarmsInput) SetProbeHook(v string) *ListVirusAlarmsInput {
	s.ProbeHook = &v
	return s
}

// SetRaspArgv sets the RaspArgv field's value.
func (s *ListVirusAlarmsInput) SetRaspArgv(v string) *ListVirusAlarmsInput {
	s.RaspArgv = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListVirusAlarmsInput) SetSortBy(v string) *ListVirusAlarmsInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListVirusAlarmsInput) SetSortOrder(v string) *ListVirusAlarmsInput {
	s.SortOrder = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListVirusAlarmsInput) SetStatus(v []*int32) *ListVirusAlarmsInput {
	s.Status = v
	return s
}

// SetTaskID sets the TaskID field's value.
func (s *ListVirusAlarmsInput) SetTaskID(v string) *ListVirusAlarmsInput {
	s.TaskID = &v
	return s
}

// SetTimeEnd sets the TimeEnd field's value.
func (s *ListVirusAlarmsInput) SetTimeEnd(v int32) *ListVirusAlarmsInput {
	s.TimeEnd = &v
	return s
}

// SetTimeStart sets the TimeStart field's value.
func (s *ListVirusAlarmsInput) SetTimeStart(v int32) *ListVirusAlarmsInput {
	s.TimeStart = &v
	return s
}

// SetTopGroupID sets the TopGroupID field's value.
func (s *ListVirusAlarmsInput) SetTopGroupID(v string) *ListVirusAlarmsInput {
	s.TopGroupID = &v
	return s
}

// SetType sets the Type field's value.
func (s *ListVirusAlarmsInput) SetType(v []*string) *ListVirusAlarmsInput {
	s.Type = v
	return s
}

// SetVirusType sets the VirusType field's value.
func (s *ListVirusAlarmsInput) SetVirusType(v string) *ListVirusAlarmsInput {
	s.VirusType = &v
	return s
}

// SetWhiteListID sets the WhiteListID field's value.
func (s *ListVirusAlarmsInput) SetWhiteListID(v string) *ListVirusAlarmsInput {
	s.WhiteListID = &v
	return s
}

// SetWhiteListName sets the WhiteListName field's value.
func (s *ListVirusAlarmsInput) SetWhiteListName(v string) *ListVirusAlarmsInput {
	s.WhiteListName = &v
	return s
}

type ListVirusAlarmsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListVirusAlarmsOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListVirusAlarmsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListVirusAlarmsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListVirusAlarmsOutput) SetData(v []*DataForListVirusAlarmsOutput) *ListVirusAlarmsOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListVirusAlarmsOutput) SetPageNumber(v int32) *ListVirusAlarmsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListVirusAlarmsOutput) SetPageSize(v int32) *ListVirusAlarmsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListVirusAlarmsOutput) SetTotalCount(v int32) *ListVirusAlarmsOutput {
	s.TotalCount = &v
	return s
}
