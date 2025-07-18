// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package aiotvideo20231001

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRecordListCommon = "GetRecordList"

// GetRecordListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRecordListCommon operation. The "output" return
// value will be populated with the GetRecordListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRecordListCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRecordListCommon Send returns without error.
//
// See GetRecordListCommon for more information on using the GetRecordListCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRecordListCommonRequest method.
//    req, resp := client.GetRecordListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AIOTVIDEO20231001) GetRecordListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRecordListCommon,
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

// GetRecordListCommon API operation for AIOTVIDEO20231001.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AIOTVIDEO20231001's
// API operation GetRecordListCommon for usage and error information.
func (c *AIOTVIDEO20231001) GetRecordListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRecordListCommonRequest(input)
	return out, req.Send()
}

// GetRecordListCommonWithContext is the same as GetRecordListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRecordListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AIOTVIDEO20231001) GetRecordListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRecordListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRecordList = "GetRecordList"

// GetRecordListRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRecordList operation. The "output" return
// value will be populated with the GetRecordListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRecordListCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRecordListCommon Send returns without error.
//
// See GetRecordList for more information on using the GetRecordList
// API call, and error handling.
//
//    // Example sending a request using the GetRecordListRequest method.
//    req, resp := client.GetRecordListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AIOTVIDEO20231001) GetRecordListRequest(input *GetRecordListInput) (req *request.Request, output *GetRecordListOutput) {
	op := &request.Operation{
		Name:       opGetRecordList,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRecordListInput{}
	}

	output = &GetRecordListOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetRecordList API operation for AIOTVIDEO20231001.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AIOTVIDEO20231001's
// API operation GetRecordList for usage and error information.
func (c *AIOTVIDEO20231001) GetRecordList(input *GetRecordListInput) (*GetRecordListOutput, error) {
	req, out := c.GetRecordListRequest(input)
	return out, req.Send()
}

// GetRecordListWithContext is the same as GetRecordList with the addition of
// the ability to pass a context and additional request options.
//
// See GetRecordList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AIOTVIDEO20231001) GetRecordListWithContext(ctx volcengine.Context, input *GetRecordListInput, opts ...request.Option) (*GetRecordListOutput, error) {
	req, out := c.GetRecordListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRecordListInput struct {
	_ struct{} `type:"structure"`

	// ChannelID is a required field
	ChannelID *string `type:"string" required:"true"`

	// DeviceNSID is a required field
	DeviceNSID *string `type:"string" required:"true"`

	// EndTime is a required field
	EndTime *int32 `type:"int32" required:"true"`

	// Order is a required field
	Order *string `type:"string" required:"true"`

	// RecordType is a required field
	RecordType *string `type:"string" required:"true"`

	Resolution *string `type:"string"`

	SpaceID *string `type:"string"`

	// StartTime is a required field
	StartTime *int32 `type:"int32" required:"true"`

	// StreamID is a required field
	StreamID *string `type:"string" required:"true"`

	StreamingIndex *int32 `type:"int32"`

	// TimeoutInSec is a required field
	TimeoutInSec *int32 `type:"int32" required:"true"`
}

// String returns the string representation
func (s GetRecordListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRecordListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRecordListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRecordListInput"}
	if s.ChannelID == nil {
		invalidParams.Add(request.NewErrParamRequired("ChannelID"))
	}
	if s.DeviceNSID == nil {
		invalidParams.Add(request.NewErrParamRequired("DeviceNSID"))
	}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.Order == nil {
		invalidParams.Add(request.NewErrParamRequired("Order"))
	}
	if s.RecordType == nil {
		invalidParams.Add(request.NewErrParamRequired("RecordType"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}
	if s.StreamID == nil {
		invalidParams.Add(request.NewErrParamRequired("StreamID"))
	}
	if s.TimeoutInSec == nil {
		invalidParams.Add(request.NewErrParamRequired("TimeoutInSec"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetChannelID sets the ChannelID field's value.
func (s *GetRecordListInput) SetChannelID(v string) *GetRecordListInput {
	s.ChannelID = &v
	return s
}

// SetDeviceNSID sets the DeviceNSID field's value.
func (s *GetRecordListInput) SetDeviceNSID(v string) *GetRecordListInput {
	s.DeviceNSID = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *GetRecordListInput) SetEndTime(v int32) *GetRecordListInput {
	s.EndTime = &v
	return s
}

// SetOrder sets the Order field's value.
func (s *GetRecordListInput) SetOrder(v string) *GetRecordListInput {
	s.Order = &v
	return s
}

// SetRecordType sets the RecordType field's value.
func (s *GetRecordListInput) SetRecordType(v string) *GetRecordListInput {
	s.RecordType = &v
	return s
}

// SetResolution sets the Resolution field's value.
func (s *GetRecordListInput) SetResolution(v string) *GetRecordListInput {
	s.Resolution = &v
	return s
}

// SetSpaceID sets the SpaceID field's value.
func (s *GetRecordListInput) SetSpaceID(v string) *GetRecordListInput {
	s.SpaceID = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *GetRecordListInput) SetStartTime(v int32) *GetRecordListInput {
	s.StartTime = &v
	return s
}

// SetStreamID sets the StreamID field's value.
func (s *GetRecordListInput) SetStreamID(v string) *GetRecordListInput {
	s.StreamID = &v
	return s
}

// SetStreamingIndex sets the StreamingIndex field's value.
func (s *GetRecordListInput) SetStreamingIndex(v int32) *GetRecordListInput {
	s.StreamingIndex = &v
	return s
}

// SetTimeoutInSec sets the TimeoutInSec field's value.
func (s *GetRecordListInput) SetTimeoutInSec(v int32) *GetRecordListInput {
	s.TimeoutInSec = &v
	return s
}

type GetRecordListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForGetRecordListOutput `type:"list"`

	Num *int32 `type:"int32"`
}

// String returns the string representation
func (s GetRecordListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRecordListOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *GetRecordListOutput) SetItems(v []*ItemForGetRecordListOutput) *GetRecordListOutput {
	s.Items = v
	return s
}

// SetNum sets the Num field's value.
func (s *GetRecordListOutput) SetNum(v int32) *GetRecordListOutput {
	s.Num = &v
	return s
}

type ItemForGetRecordListOutput struct {
	_ struct{} `type:"structure"`

	ChannelID *string `type:"string"`

	EndTime *string `type:"string"`

	EndTimeStamp *int32 `type:"int32"`

	FileSize *string `type:"string"`

	Name *string `type:"string"`

	Secrecy *string `type:"string"`

	StartTime *string `type:"string"`

	StartTimeStamp *int32 `type:"int32"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ItemForGetRecordListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForGetRecordListOutput) GoString() string {
	return s.String()
}

// SetChannelID sets the ChannelID field's value.
func (s *ItemForGetRecordListOutput) SetChannelID(v string) *ItemForGetRecordListOutput {
	s.ChannelID = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *ItemForGetRecordListOutput) SetEndTime(v string) *ItemForGetRecordListOutput {
	s.EndTime = &v
	return s
}

// SetEndTimeStamp sets the EndTimeStamp field's value.
func (s *ItemForGetRecordListOutput) SetEndTimeStamp(v int32) *ItemForGetRecordListOutput {
	s.EndTimeStamp = &v
	return s
}

// SetFileSize sets the FileSize field's value.
func (s *ItemForGetRecordListOutput) SetFileSize(v string) *ItemForGetRecordListOutput {
	s.FileSize = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForGetRecordListOutput) SetName(v string) *ItemForGetRecordListOutput {
	s.Name = &v
	return s
}

// SetSecrecy sets the Secrecy field's value.
func (s *ItemForGetRecordListOutput) SetSecrecy(v string) *ItemForGetRecordListOutput {
	s.Secrecy = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ItemForGetRecordListOutput) SetStartTime(v string) *ItemForGetRecordListOutput {
	s.StartTime = &v
	return s
}

// SetStartTimeStamp sets the StartTimeStamp field's value.
func (s *ItemForGetRecordListOutput) SetStartTimeStamp(v int32) *ItemForGetRecordListOutput {
	s.StartTimeStamp = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForGetRecordListOutput) SetType(v string) *ItemForGetRecordListOutput {
	s.Type = &v
	return s
}
