// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestoreTableCommon = "RestoreTable"

// RestoreTableCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreTableCommon operation. The "output" return
// value will be populated with the RestoreTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreTableCommon Send returns without error.
//
// See RestoreTableCommon for more information on using the RestoreTableCommon
// API call, and error handling.
//
//	// Example sending a request using the RestoreTableCommonRequest method.
//	req, resp := client.RestoreTableCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VEDBM) RestoreTableCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestoreTableCommon,
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

// RestoreTableCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation RestoreTableCommon for usage and error information.
func (c *VEDBM) RestoreTableCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestoreTableCommonRequest(input)
	return out, req.Send()
}

// RestoreTableCommonWithContext is the same as RestoreTableCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreTableCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) RestoreTableCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestoreTableCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestoreTable = "RestoreTable"

// RestoreTableRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreTable operation. The "output" return
// value will be populated with the RestoreTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreTableCommon Send returns without error.
//
// See RestoreTable for more information on using the RestoreTable
// API call, and error handling.
//
//	// Example sending a request using the RestoreTableRequest method.
//	req, resp := client.RestoreTableRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VEDBM) RestoreTableRequest(input *RestoreTableInput) (req *request.Request, output *RestoreTableOutput) {
	op := &request.Operation{
		Name:       opRestoreTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreTableInput{}
	}

	output = &RestoreTableOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestoreTable API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation RestoreTable for usage and error information.
func (c *VEDBM) RestoreTable(input *RestoreTableInput) (*RestoreTableOutput, error) {
	req, out := c.RestoreTableRequest(input)
	return out, req.Send()
}

// RestoreTableWithContext is the same as RestoreTable with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreTable for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) RestoreTableWithContext(ctx volcengine.Context, input *RestoreTableInput, opts ...request.Option) (*RestoreTableOutput, error) {
	req, out := c.RestoreTableRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RestoreTableInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	RestoreTime *string `type:"string" json:",omitempty"`

	TableMeta []*TableMetaForRestoreTableInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RestoreTableInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreTableInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreTableInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestoreTableInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupId sets the BackupId field's value.
func (s *RestoreTableInput) SetBackupId(v string) *RestoreTableInput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestoreTableInput) SetInstanceId(v string) *RestoreTableInput {
	s.InstanceId = &v
	return s
}

// SetRestoreTime sets the RestoreTime field's value.
func (s *RestoreTableInput) SetRestoreTime(v string) *RestoreTableInput {
	s.RestoreTime = &v
	return s
}

// SetTableMeta sets the TableMeta field's value.
func (s *RestoreTableInput) SetTableMeta(v []*TableMetaForRestoreTableInput) *RestoreTableInput {
	s.TableMeta = v
	return s
}

type RestoreTableOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RestoreTableOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreTableOutput) GoString() string {
	return s.String()
}

type TableForRestoreTableInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NewTableName *string `type:"string" json:",omitempty"`

	TableName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TableForRestoreTableInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TableForRestoreTableInput) GoString() string {
	return s.String()
}

// SetNewTableName sets the NewTableName field's value.
func (s *TableForRestoreTableInput) SetNewTableName(v string) *TableForRestoreTableInput {
	s.NewTableName = &v
	return s
}

// SetTableName sets the TableName field's value.
func (s *TableForRestoreTableInput) SetTableName(v string) *TableForRestoreTableInput {
	s.TableName = &v
	return s
}

type TableMetaForRestoreTableInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DBName *string `type:"string" json:",omitempty"`

	NewDBName *string `type:"string" json:",omitempty"`

	Table []*TableForRestoreTableInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TableMetaForRestoreTableInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TableMetaForRestoreTableInput) GoString() string {
	return s.String()
}

// SetDBName sets the DBName field's value.
func (s *TableMetaForRestoreTableInput) SetDBName(v string) *TableMetaForRestoreTableInput {
	s.DBName = &v
	return s
}

// SetNewDBName sets the NewDBName field's value.
func (s *TableMetaForRestoreTableInput) SetNewDBName(v string) *TableMetaForRestoreTableInput {
	s.NewDBName = &v
	return s
}

// SetTable sets the Table field's value.
func (s *TableMetaForRestoreTableInput) SetTable(v []*TableForRestoreTableInput) *TableMetaForRestoreTableInput {
	s.Table = v
	return s
}
