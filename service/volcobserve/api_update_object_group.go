// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateObjectGroupCommon = "UpdateObjectGroup"

// UpdateObjectGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateObjectGroupCommon operation. The "output" return
// value will be populated with the UpdateObjectGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateObjectGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateObjectGroupCommon Send returns without error.
//
// See UpdateObjectGroupCommon for more information on using the UpdateObjectGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateObjectGroupCommonRequest method.
//    req, resp := client.UpdateObjectGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) UpdateObjectGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateObjectGroupCommon,
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

// UpdateObjectGroupCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation UpdateObjectGroupCommon for usage and error information.
func (c *VOLCOBSERVE) UpdateObjectGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateObjectGroupCommonRequest(input)
	return out, req.Send()
}

// UpdateObjectGroupCommonWithContext is the same as UpdateObjectGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateObjectGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) UpdateObjectGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateObjectGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateObjectGroup = "UpdateObjectGroup"

// UpdateObjectGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateObjectGroup operation. The "output" return
// value will be populated with the UpdateObjectGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateObjectGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateObjectGroupCommon Send returns without error.
//
// See UpdateObjectGroup for more information on using the UpdateObjectGroup
// API call, and error handling.
//
//    // Example sending a request using the UpdateObjectGroupRequest method.
//    req, resp := client.UpdateObjectGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) UpdateObjectGroupRequest(input *UpdateObjectGroupInput) (req *request.Request, output *UpdateObjectGroupOutput) {
	op := &request.Operation{
		Name:       opUpdateObjectGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateObjectGroupInput{}
	}

	output = &UpdateObjectGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateObjectGroup API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation UpdateObjectGroup for usage and error information.
func (c *VOLCOBSERVE) UpdateObjectGroup(input *UpdateObjectGroupInput) (*UpdateObjectGroupOutput, error) {
	req, out := c.UpdateObjectGroupRequest(input)
	return out, req.Send()
}

// UpdateObjectGroupWithContext is the same as UpdateObjectGroup with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateObjectGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) UpdateObjectGroupWithContext(ctx volcengine.Context, input *UpdateObjectGroupInput, opts ...request.Option) (*UpdateObjectGroupOutput, error) {
	req, out := c.UpdateObjectGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DimensionConditionsForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MetaCondition *MetaConditionForUpdateObjectGroupInput `type:"structure" json:",omitempty"`

	ProjectCondition *ProjectConditionForUpdateObjectGroupInput `type:"structure" json:",omitempty"`

	TagCondition *TagConditionForUpdateObjectGroupInput `type:"structure" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionConditionsForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionConditionsForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetMetaCondition sets the MetaCondition field's value.
func (s *DimensionConditionsForUpdateObjectGroupInput) SetMetaCondition(v *MetaConditionForUpdateObjectGroupInput) *DimensionConditionsForUpdateObjectGroupInput {
	s.MetaCondition = v
	return s
}

// SetProjectCondition sets the ProjectCondition field's value.
func (s *DimensionConditionsForUpdateObjectGroupInput) SetProjectCondition(v *ProjectConditionForUpdateObjectGroupInput) *DimensionConditionsForUpdateObjectGroupInput {
	s.ProjectCondition = v
	return s
}

// SetTagCondition sets the TagCondition field's value.
func (s *DimensionConditionsForUpdateObjectGroupInput) SetTagCondition(v *TagConditionForUpdateObjectGroupInput) *DimensionConditionsForUpdateObjectGroupInput {
	s.TagCondition = v
	return s
}

// SetType sets the Type field's value.
func (s *DimensionConditionsForUpdateObjectGroupInput) SetType(v string) *DimensionConditionsForUpdateObjectGroupInput {
	s.Type = &v
	return s
}

type MetaConditionForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllDimensions *bool `type:"boolean" json:",omitempty"`

	Condition *string `type:"string" json:",omitempty"`

	Metas []*MetaForUpdateObjectGroupInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetaConditionForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetaConditionForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetAllDimensions sets the AllDimensions field's value.
func (s *MetaConditionForUpdateObjectGroupInput) SetAllDimensions(v bool) *MetaConditionForUpdateObjectGroupInput {
	s.AllDimensions = &v
	return s
}

// SetCondition sets the Condition field's value.
func (s *MetaConditionForUpdateObjectGroupInput) SetCondition(v string) *MetaConditionForUpdateObjectGroupInput {
	s.Condition = &v
	return s
}

// SetMetas sets the Metas field's value.
func (s *MetaConditionForUpdateObjectGroupInput) SetMetas(v []*MetaForUpdateObjectGroupInput) *MetaConditionForUpdateObjectGroupInput {
	s.Metas = v
	return s
}

type MetaForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Comparator *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetaForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetaForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetComparator sets the Comparator field's value.
func (s *MetaForUpdateObjectGroupInput) SetComparator(v string) *MetaForUpdateObjectGroupInput {
	s.Comparator = &v
	return s
}

// SetKey sets the Key field's value.
func (s *MetaForUpdateObjectGroupInput) SetKey(v string) *MetaForUpdateObjectGroupInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *MetaForUpdateObjectGroupInput) SetValues(v []*string) *MetaForUpdateObjectGroupInput {
	s.Values = v
	return s
}

type ObjectForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DimensionConditions *DimensionConditionsForUpdateObjectGroupInput `type:"structure" json:",omitempty"`

	Dimensions map[string][]*string `type:"map" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ObjectForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ObjectForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetDimensionConditions sets the DimensionConditions field's value.
func (s *ObjectForUpdateObjectGroupInput) SetDimensionConditions(v *DimensionConditionsForUpdateObjectGroupInput) *ObjectForUpdateObjectGroupInput {
	s.DimensionConditions = v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *ObjectForUpdateObjectGroupInput) SetDimensions(v map[string][]*string) *ObjectForUpdateObjectGroupInput {
	s.Dimensions = v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *ObjectForUpdateObjectGroupInput) SetNamespace(v string) *ObjectForUpdateObjectGroupInput {
	s.Namespace = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ObjectForUpdateObjectGroupInput) SetRegion(v string) *ObjectForUpdateObjectGroupInput {
	s.Region = &v
	return s
}

// SetType sets the Type field's value.
func (s *ObjectForUpdateObjectGroupInput) SetType(v string) *ObjectForUpdateObjectGroupInput {
	s.Type = &v
	return s
}

type ProjectConditionForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Projects []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ProjectConditionForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProjectConditionForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetProjects sets the Projects field's value.
func (s *ProjectConditionForUpdateObjectGroupInput) SetProjects(v []*string) *ProjectConditionForUpdateObjectGroupInput {
	s.Projects = v
	return s
}

type TagConditionForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Condition *string `type:"string" json:",omitempty"`

	Tags []*TagForUpdateObjectGroupInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagConditionForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagConditionForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetCondition sets the Condition field's value.
func (s *TagConditionForUpdateObjectGroupInput) SetCondition(v string) *TagConditionForUpdateObjectGroupInput {
	s.Condition = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *TagConditionForUpdateObjectGroupInput) SetTags(v []*TagForUpdateObjectGroupInput) *TagConditionForUpdateObjectGroupInput {
	s.Tags = v
	return s
}

type TagForUpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Comparator *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagForUpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForUpdateObjectGroupInput) GoString() string {
	return s.String()
}

// SetComparator sets the Comparator field's value.
func (s *TagForUpdateObjectGroupInput) SetComparator(v string) *TagForUpdateObjectGroupInput {
	s.Comparator = &v
	return s
}

// SetKey sets the Key field's value.
func (s *TagForUpdateObjectGroupInput) SetKey(v string) *TagForUpdateObjectGroupInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagForUpdateObjectGroupInput) SetValues(v []*string) *TagForUpdateObjectGroupInput {
	s.Values = v
	return s
}

type UpdateObjectGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	Objects []*ObjectForUpdateObjectGroupInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UpdateObjectGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateObjectGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateObjectGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateObjectGroupInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *UpdateObjectGroupInput) SetId(v string) *UpdateObjectGroupInput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateObjectGroupInput) SetName(v string) *UpdateObjectGroupInput {
	s.Name = &v
	return s
}

// SetObjects sets the Objects field's value.
func (s *UpdateObjectGroupInput) SetObjects(v []*ObjectForUpdateObjectGroupInput) *UpdateObjectGroupInput {
	s.Objects = v
	return s
}

type UpdateObjectGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateObjectGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateObjectGroupOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *UpdateObjectGroupOutput) SetData(v string) *UpdateObjectGroupOutput {
	s.Data = &v
	return s
}
