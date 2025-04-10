// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetDXPInstanceCommon = "GetDXPInstance"

// GetDXPInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDXPInstanceCommon operation. The "output" return
// value will be populated with the GetDXPInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDXPInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDXPInstanceCommon Send returns without error.
//
// See GetDXPInstanceCommon for more information on using the GetDXPInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the GetDXPInstanceCommonRequest method.
//    req, resp := client.GetDXPInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) GetDXPInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetDXPInstanceCommon,
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

// GetDXPInstanceCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation GetDXPInstanceCommon for usage and error information.
func (c *EDX) GetDXPInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetDXPInstanceCommonRequest(input)
	return out, req.Send()
}

// GetDXPInstanceCommonWithContext is the same as GetDXPInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetDXPInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) GetDXPInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetDXPInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetDXPInstance = "GetDXPInstance"

// GetDXPInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDXPInstance operation. The "output" return
// value will be populated with the GetDXPInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDXPInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDXPInstanceCommon Send returns without error.
//
// See GetDXPInstance for more information on using the GetDXPInstance
// API call, and error handling.
//
//    // Example sending a request using the GetDXPInstanceRequest method.
//    req, resp := client.GetDXPInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) GetDXPInstanceRequest(input *GetDXPInstanceInput) (req *request.Request, output *GetDXPInstanceOutput) {
	op := &request.Operation{
		Name:       opGetDXPInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDXPInstanceInput{}
	}

	output = &GetDXPInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetDXPInstance API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation GetDXPInstance for usage and error information.
func (c *EDX) GetDXPInstance(input *GetDXPInstanceInput) (*GetDXPInstanceOutput, error) {
	req, out := c.GetDXPInstanceRequest(input)
	return out, req.Send()
}

// GetDXPInstanceWithContext is the same as GetDXPInstance with the addition of
// the ability to pass a context and additional request options.
//
// See GetDXPInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) GetDXPInstanceWithContext(ctx volcengine.Context, input *GetDXPInstanceInput, opts ...request.Option) (*GetDXPInstanceOutput, error) {
	req, out := c.GetDXPInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConnectionInfoForGetDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Bandwidth *string `type:"string" json:",omitempty"`

	ContactName *string `type:"string" json:",omitempty"`

	ContactPhone *string `type:"string" json:",omitempty"`

	FieldEngineer []*FieldEngineerForGetDXPInstanceOutput `type:"list" json:",omitempty"`

	IDCAddr *string `type:"string" json:",omitempty"`

	IDCCode *string `type:"string" json:",omitempty"`

	IDCLocation *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConnectionInfoForGetDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConnectionInfoForGetDXPInstanceOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetBandwidth(v string) *ConnectionInfoForGetDXPInstanceOutput {
	s.Bandwidth = &v
	return s
}

// SetContactName sets the ContactName field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetContactName(v string) *ConnectionInfoForGetDXPInstanceOutput {
	s.ContactName = &v
	return s
}

// SetContactPhone sets the ContactPhone field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetContactPhone(v string) *ConnectionInfoForGetDXPInstanceOutput {
	s.ContactPhone = &v
	return s
}

// SetFieldEngineer sets the FieldEngineer field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetFieldEngineer(v []*FieldEngineerForGetDXPInstanceOutput) *ConnectionInfoForGetDXPInstanceOutput {
	s.FieldEngineer = v
	return s
}

// SetIDCAddr sets the IDCAddr field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetIDCAddr(v string) *ConnectionInfoForGetDXPInstanceOutput {
	s.IDCAddr = &v
	return s
}

// SetIDCCode sets the IDCCode field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetIDCCode(v string) *ConnectionInfoForGetDXPInstanceOutput {
	s.IDCCode = &v
	return s
}

// SetIDCLocation sets the IDCLocation field's value.
func (s *ConnectionInfoForGetDXPInstanceOutput) SetIDCLocation(v string) *ConnectionInfoForGetDXPInstanceOutput {
	s.IDCLocation = &v
	return s
}

type ConstructionInfoForGetDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConstructionEndDate *string `type:"string" json:",omitempty"`

	ConstructionStartDate *string `type:"string" json:",omitempty"`

	ContactName *string `type:"string" json:",omitempty"`

	ContactPhone *string `type:"string" json:",omitempty"`

	DeviceInfo *string `type:"string" json:",omitempty"`

	IDCAddr *string `type:"string" json:",omitempty"`

	ODFInfo *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConstructionInfoForGetDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConstructionInfoForGetDXPInstanceOutput) GoString() string {
	return s.String()
}

// SetConstructionEndDate sets the ConstructionEndDate field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetConstructionEndDate(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.ConstructionEndDate = &v
	return s
}

// SetConstructionStartDate sets the ConstructionStartDate field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetConstructionStartDate(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.ConstructionStartDate = &v
	return s
}

// SetContactName sets the ContactName field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetContactName(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.ContactName = &v
	return s
}

// SetContactPhone sets the ContactPhone field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetContactPhone(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.ContactPhone = &v
	return s
}

// SetDeviceInfo sets the DeviceInfo field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetDeviceInfo(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.DeviceInfo = &v
	return s
}

// SetIDCAddr sets the IDCAddr field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetIDCAddr(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.IDCAddr = &v
	return s
}

// SetODFInfo sets the ODFInfo field's value.
func (s *ConstructionInfoForGetDXPInstanceOutput) SetODFInfo(v string) *ConstructionInfoForGetDXPInstanceOutput {
	s.ODFInfo = &v
	return s
}

type DXPInstanceForGetDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AP *string `type:"string" json:",omitempty"`

	Area *string `type:"string" json:",omitempty"`

	AutoRenew *string `type:"string" json:",omitempty"`

	Bandwidth *string `type:"string" json:",omitempty"`

	ConnectionInfo *ConnectionInfoForGetDXPInstanceOutput `type:"structure" json:",omitempty"`

	ConstructionInfo *ConstructionInfoForGetDXPInstanceOutput `type:"structure" json:",omitempty"`

	HealthStatus *bool `type:"boolean" json:",omitempty"`

	ISP *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceNo *string `type:"string" json:",omitempty"`

	ModuleType *string `type:"string" json:",omitempty"`

	PortType *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	VIFIdList []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DXPInstanceForGetDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DXPInstanceForGetDXPInstanceOutput) GoString() string {
	return s.String()
}

// SetAP sets the AP field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetAP(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.AP = &v
	return s
}

// SetArea sets the Area field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetArea(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.Area = &v
	return s
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetAutoRenew(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.AutoRenew = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetBandwidth(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.Bandwidth = &v
	return s
}

// SetConnectionInfo sets the ConnectionInfo field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetConnectionInfo(v *ConnectionInfoForGetDXPInstanceOutput) *DXPInstanceForGetDXPInstanceOutput {
	s.ConnectionInfo = v
	return s
}

// SetConstructionInfo sets the ConstructionInfo field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetConstructionInfo(v *ConstructionInfoForGetDXPInstanceOutput) *DXPInstanceForGetDXPInstanceOutput {
	s.ConstructionInfo = v
	return s
}

// SetHealthStatus sets the HealthStatus field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetHealthStatus(v bool) *DXPInstanceForGetDXPInstanceOutput {
	s.HealthStatus = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetISP(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.ISP = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetInstanceId(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetInstanceName(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetInstanceNo(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.InstanceNo = &v
	return s
}

// SetModuleType sets the ModuleType field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetModuleType(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.ModuleType = &v
	return s
}

// SetPortType sets the PortType field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetPortType(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.PortType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetStatus(v string) *DXPInstanceForGetDXPInstanceOutput {
	s.Status = &v
	return s
}

// SetVIFIdList sets the VIFIdList field's value.
func (s *DXPInstanceForGetDXPInstanceOutput) SetVIFIdList(v []*string) *DXPInstanceForGetDXPInstanceOutput {
	s.VIFIdList = v
	return s
}

type FieldEngineerForGetDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Phone *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FieldEngineerForGetDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FieldEngineerForGetDXPInstanceOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *FieldEngineerForGetDXPInstanceOutput) SetName(v string) *FieldEngineerForGetDXPInstanceOutput {
	s.Name = &v
	return s
}

// SetPhone sets the Phone field's value.
func (s *FieldEngineerForGetDXPInstanceOutput) SetPhone(v string) *FieldEngineerForGetDXPInstanceOutput {
	s.Phone = &v
	return s
}

type GetDXPInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Area is a required field
	Area *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetDXPInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDXPInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDXPInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetDXPInstanceInput"}
	if s.Area == nil {
		invalidParams.Add(request.NewErrParamRequired("Area"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetArea sets the Area field's value.
func (s *GetDXPInstanceInput) SetArea(v string) *GetDXPInstanceInput {
	s.Area = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetDXPInstanceInput) SetInstanceId(v string) *GetDXPInstanceInput {
	s.InstanceId = &v
	return s
}

type GetDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DXPInstance *DXPInstanceForGetDXPInstanceOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDXPInstanceOutput) GoString() string {
	return s.String()
}

// SetDXPInstance sets the DXPInstance field's value.
func (s *GetDXPInstanceOutput) SetDXPInstance(v *DXPInstanceForGetDXPInstanceOutput) *GetDXPInstanceOutput {
	s.DXPInstance = v
	return s
}
