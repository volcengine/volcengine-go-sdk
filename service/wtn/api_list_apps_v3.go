// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package wtn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAppsV3Common = "ListAppsV3"

// ListAppsV3CommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAppsV3Common operation. The "output" return
// value will be populated with the ListAppsV3Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAppsV3Common Request to send the API call to the service.
// the "output" return value is not valid until after ListAppsV3Common Send returns without error.
//
// See ListAppsV3Common for more information on using the ListAppsV3Common
// API call, and error handling.
//
//    // Example sending a request using the ListAppsV3CommonRequest method.
//    req, resp := client.ListAppsV3CommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WTN) ListAppsV3CommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAppsV3Common,
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

// ListAppsV3Common API operation for WTN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WTN's
// API operation ListAppsV3Common for usage and error information.
func (c *WTN) ListAppsV3Common(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAppsV3CommonRequest(input)
	return out, req.Send()
}

// ListAppsV3CommonWithContext is the same as ListAppsV3Common with the addition of
// the ability to pass a context and additional request options.
//
// See ListAppsV3Common for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WTN) ListAppsV3CommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAppsV3CommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAppsV3 = "ListAppsV3"

// ListAppsV3Request generates a "volcengine/request.Request" representing the
// client's request for the ListAppsV3 operation. The "output" return
// value will be populated with the ListAppsV3Common request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAppsV3Common Request to send the API call to the service.
// the "output" return value is not valid until after ListAppsV3Common Send returns without error.
//
// See ListAppsV3 for more information on using the ListAppsV3
// API call, and error handling.
//
//    // Example sending a request using the ListAppsV3Request method.
//    req, resp := client.ListAppsV3Request(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WTN) ListAppsV3Request(input *ListAppsV3Input) (req *request.Request, output *ListAppsV3Output) {
	op := &request.Operation{
		Name:       opListAppsV3,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAppsV3Input{}
	}

	output = &ListAppsV3Output{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAppsV3 API operation for WTN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WTN's
// API operation ListAppsV3 for usage and error information.
func (c *WTN) ListAppsV3(input *ListAppsV3Input) (*ListAppsV3Output, error) {
	req, out := c.ListAppsV3Request(input)
	return out, req.Send()
}

// ListAppsV3WithContext is the same as ListAppsV3 with the addition of
// the ability to pass a context and additional request options.
//
// See ListAppsV3 for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WTN) ListAppsV3WithContext(ctx volcengine.Context, input *ListAppsV3Input, opts ...request.Option) (*ListAppsV3Output, error) {
	req, out := c.ListAppsV3Request(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AppListForListAppsV3Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AppId *string `type:"string" json:",omitempty"`

	AppKey *string `type:"string" json:",omitempty"`

	AppName *string `type:"string" json:",omitempty"`

	CreateAt *string `type:"string" json:",omitempty"`

	SecondaryAppKey *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AppListForListAppsV3Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AppListForListAppsV3Output) GoString() string {
	return s.String()
}

// SetAppId sets the AppId field's value.
func (s *AppListForListAppsV3Output) SetAppId(v string) *AppListForListAppsV3Output {
	s.AppId = &v
	return s
}

// SetAppKey sets the AppKey field's value.
func (s *AppListForListAppsV3Output) SetAppKey(v string) *AppListForListAppsV3Output {
	s.AppKey = &v
	return s
}

// SetAppName sets the AppName field's value.
func (s *AppListForListAppsV3Output) SetAppName(v string) *AppListForListAppsV3Output {
	s.AppName = &v
	return s
}

// SetCreateAt sets the CreateAt field's value.
func (s *AppListForListAppsV3Output) SetCreateAt(v string) *AppListForListAppsV3Output {
	s.CreateAt = &v
	return s
}

// SetSecondaryAppKey sets the SecondaryAppKey field's value.
func (s *AppListForListAppsV3Output) SetSecondaryAppKey(v string) *AppListForListAppsV3Output {
	s.SecondaryAppKey = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AppListForListAppsV3Output) SetStatus(v string) *AppListForListAppsV3Output {
	s.Status = &v
	return s
}

type ListAppsV3Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AppId *string `type:"string" json:",omitempty"`

	Limit *int32 `type:"int32" json:",omitempty"`

	Offset *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	Reverse *int32 `type:"int32" json:",omitempty"`

	TagFilters []*TagFilterForListAppsV3Input `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListAppsV3Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAppsV3Input) GoString() string {
	return s.String()
}

// SetAppId sets the AppId field's value.
func (s *ListAppsV3Input) SetAppId(v string) *ListAppsV3Input {
	s.AppId = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListAppsV3Input) SetLimit(v int32) *ListAppsV3Input {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAppsV3Input) SetOffset(v int32) *ListAppsV3Input {
	s.Offset = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListAppsV3Input) SetProjectName(v string) *ListAppsV3Input {
	s.ProjectName = &v
	return s
}

// SetReverse sets the Reverse field's value.
func (s *ListAppsV3Input) SetReverse(v int32) *ListAppsV3Input {
	s.Reverse = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *ListAppsV3Input) SetTagFilters(v []*TagFilterForListAppsV3Input) *ListAppsV3Input {
	s.TagFilters = v
	return s
}

type ListAppsV3Output struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AppList []*AppListForListAppsV3Output `type:"list" json:",omitempty"`

	Limit *int32 `type:"int32" json:",omitempty"`

	Offset *int32 `type:"int32" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListAppsV3Output) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAppsV3Output) GoString() string {
	return s.String()
}

// SetAppList sets the AppList field's value.
func (s *ListAppsV3Output) SetAppList(v []*AppListForListAppsV3Output) *ListAppsV3Output {
	s.AppList = v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListAppsV3Output) SetLimit(v int32) *ListAppsV3Output {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAppsV3Output) SetOffset(v int32) *ListAppsV3Output {
	s.Offset = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListAppsV3Output) SetTotal(v int32) *ListAppsV3Output {
	s.Total = &v
	return s
}

type TagFilterForListAppsV3Input struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagFilterForListAppsV3Input) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForListAppsV3Input) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForListAppsV3Input) SetKey(v string) *TagFilterForListAppsV3Input {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForListAppsV3Input) SetValues(v []*string) *TagFilterForListAppsV3Input {
	s.Values = v
	return s
}
