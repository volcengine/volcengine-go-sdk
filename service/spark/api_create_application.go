// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateApplicationCommon = "createApplication"

// CreateApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateApplicationCommon operation. The "output" return
// value will be populated with the CreateApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateApplicationCommon Send returns without error.
//
// See CreateApplicationCommon for more information on using the CreateApplicationCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateApplicationCommonRequest method.
//	req, resp := client.CreateApplicationCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) CreateApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateApplicationCommon,
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

// CreateApplicationCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateApplicationCommon for usage and error information.
func (c *SPARK) CreateApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateApplicationCommonRequest(input)
	return out, req.Send()
}

// CreateApplicationCommonWithContext is the same as CreateApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateApplication = "createApplication"

// CreateApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateApplication operation. The "output" return
// value will be populated with the CreateApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateApplicationCommon Send returns without error.
//
// See CreateApplication for more information on using the CreateApplication
// API call, and error handling.
//
//	// Example sending a request using the CreateApplicationRequest method.
//	req, resp := client.CreateApplicationRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) CreateApplicationRequest(input *CreateApplicationInput) (req *request.Request, output *CreateApplicationOutput) {
	op := &request.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	output = &CreateApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateApplication API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateApplication for usage and error information.
func (c *SPARK) CreateApplication(input *CreateApplicationInput) (*CreateApplicationOutput, error) {
	req, out := c.CreateApplicationRequest(input)
	return out, req.Send()
}

// CreateApplicationWithContext is the same as CreateApplication with the addition of
// the ability to pass a context and additional request options.
//
// See CreateApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateApplicationWithContext(ctx volcengine.Context, input *CreateApplicationInput, opts ...request.Option) (*CreateApplicationOutput, error) {
	req, out := c.CreateApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateApplicationInput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	ApplicationName *string `type:"string"`

	ApplicationType *string `type:"string" enum:"EnumOfApplicationTypeForcreateApplicationInput"`

	Args *string `type:"string"`

	Conf map[string]*string `type:"map"`

	Dependency *DependencyForcreateApplicationInput `type:"structure"`

	DeployRequest *DeployRequestForcreateApplicationInput `type:"structure"`

	EngineVersion *string `type:"string" enum:"EnumOfEngineVersionForcreateApplicationInput"`

	Id *int64 `type:"int64"`

	Image *string `type:"string"`

	Jar *string `type:"string"`

	MainClass *string `type:"string"`

	ProjectId *string `type:"string"`

	SqlText *string `type:"string"`

	UniqueKey *string `type:"string"`

	UserId *string `type:"string"`
}

// String returns the string representation
func (s CreateApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateApplicationInput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CreateApplicationInput) SetAccountId(v string) *CreateApplicationInput {
	s.AccountId = &v
	return s
}

// SetApplicationName sets the ApplicationName field's value.
func (s *CreateApplicationInput) SetApplicationName(v string) *CreateApplicationInput {
	s.ApplicationName = &v
	return s
}

// SetApplicationType sets the ApplicationType field's value.
func (s *CreateApplicationInput) SetApplicationType(v string) *CreateApplicationInput {
	s.ApplicationType = &v
	return s
}

// SetArgs sets the Args field's value.
func (s *CreateApplicationInput) SetArgs(v string) *CreateApplicationInput {
	s.Args = &v
	return s
}

// SetConf sets the Conf field's value.
func (s *CreateApplicationInput) SetConf(v map[string]*string) *CreateApplicationInput {
	s.Conf = v
	return s
}

// SetDependency sets the Dependency field's value.
func (s *CreateApplicationInput) SetDependency(v *DependencyForcreateApplicationInput) *CreateApplicationInput {
	s.Dependency = v
	return s
}

// SetDeployRequest sets the DeployRequest field's value.
func (s *CreateApplicationInput) SetDeployRequest(v *DeployRequestForcreateApplicationInput) *CreateApplicationInput {
	s.DeployRequest = v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *CreateApplicationInput) SetEngineVersion(v string) *CreateApplicationInput {
	s.EngineVersion = &v
	return s
}

// SetId sets the Id field's value.
func (s *CreateApplicationInput) SetId(v int64) *CreateApplicationInput {
	s.Id = &v
	return s
}

// SetImage sets the Image field's value.
func (s *CreateApplicationInput) SetImage(v string) *CreateApplicationInput {
	s.Image = &v
	return s
}

// SetJar sets the Jar field's value.
func (s *CreateApplicationInput) SetJar(v string) *CreateApplicationInput {
	s.Jar = &v
	return s
}

// SetMainClass sets the MainClass field's value.
func (s *CreateApplicationInput) SetMainClass(v string) *CreateApplicationInput {
	s.MainClass = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *CreateApplicationInput) SetProjectId(v string) *CreateApplicationInput {
	s.ProjectId = &v
	return s
}

// SetSqlText sets the SqlText field's value.
func (s *CreateApplicationInput) SetSqlText(v string) *CreateApplicationInput {
	s.SqlText = &v
	return s
}

// SetUniqueKey sets the UniqueKey field's value.
func (s *CreateApplicationInput) SetUniqueKey(v string) *CreateApplicationInput {
	s.UniqueKey = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *CreateApplicationInput) SetUserId(v string) *CreateApplicationInput {
	s.UserId = &v
	return s
}

type CreateApplicationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateApplicationOutput) GoString() string {
	return s.String()
}

type DependencyForcreateApplicationInput struct {
	_ struct{} `type:"structure"`

	Archives []*string `type:"list"`

	Files []*string `type:"list"`

	Jars []*string `type:"list"`

	PyFiles []*string `type:"list"`
}

// String returns the string representation
func (s DependencyForcreateApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DependencyForcreateApplicationInput) GoString() string {
	return s.String()
}

// SetArchives sets the Archives field's value.
func (s *DependencyForcreateApplicationInput) SetArchives(v []*string) *DependencyForcreateApplicationInput {
	s.Archives = v
	return s
}

// SetFiles sets the Files field's value.
func (s *DependencyForcreateApplicationInput) SetFiles(v []*string) *DependencyForcreateApplicationInput {
	s.Files = v
	return s
}

// SetJars sets the Jars field's value.
func (s *DependencyForcreateApplicationInput) SetJars(v []*string) *DependencyForcreateApplicationInput {
	s.Jars = v
	return s
}

// SetPyFiles sets the PyFiles field's value.
func (s *DependencyForcreateApplicationInput) SetPyFiles(v []*string) *DependencyForcreateApplicationInput {
	s.PyFiles = v
	return s
}

type DeployRequestForcreateApplicationInput struct {
	_ struct{} `type:"structure"`

	Priority *string `type:"string"`

	ResourcePoolTrn *string `type:"string"`

	SchedulePolicy *string `type:"string"`

	ScheduleTimeout *string `type:"string"`
}

// String returns the string representation
func (s DeployRequestForcreateApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeployRequestForcreateApplicationInput) GoString() string {
	return s.String()
}

// SetPriority sets the Priority field's value.
func (s *DeployRequestForcreateApplicationInput) SetPriority(v string) *DeployRequestForcreateApplicationInput {
	s.Priority = &v
	return s
}

// SetResourcePoolTrn sets the ResourcePoolTrn field's value.
func (s *DeployRequestForcreateApplicationInput) SetResourcePoolTrn(v string) *DeployRequestForcreateApplicationInput {
	s.ResourcePoolTrn = &v
	return s
}

// SetSchedulePolicy sets the SchedulePolicy field's value.
func (s *DeployRequestForcreateApplicationInput) SetSchedulePolicy(v string) *DeployRequestForcreateApplicationInput {
	s.SchedulePolicy = &v
	return s
}

// SetScheduleTimeout sets the ScheduleTimeout field's value.
func (s *DeployRequestForcreateApplicationInput) SetScheduleTimeout(v string) *DeployRequestForcreateApplicationInput {
	s.ScheduleTimeout = &v
	return s
}

const (
	// EnumOfApplicationTypeForcreateApplicationInputSparkBatchJar is a EnumOfApplicationTypeForcreateApplicationInput enum value
	EnumOfApplicationTypeForcreateApplicationInputSparkBatchJar = "SPARK_BATCH_JAR"

	// EnumOfApplicationTypeForcreateApplicationInputSparkBatchSql is a EnumOfApplicationTypeForcreateApplicationInput enum value
	EnumOfApplicationTypeForcreateApplicationInputSparkBatchSql = "SPARK_BATCH_SQL"

	// EnumOfApplicationTypeForcreateApplicationInputSparkBatchPython is a EnumOfApplicationTypeForcreateApplicationInput enum value
	EnumOfApplicationTypeForcreateApplicationInputSparkBatchPython = "SPARK_BATCH_PYTHON"
)

const (
	// EnumOfEngineVersionForcreateApplicationInputSparkVersion301Os is a EnumOfEngineVersionForcreateApplicationInput enum value
	EnumOfEngineVersionForcreateApplicationInputSparkVersion301Os = "SPARK_VERSION_3_0_1_OS"

	// EnumOfEngineVersionForcreateApplicationInputSparkVersion322 is a EnumOfEngineVersionForcreateApplicationInput enum value
	EnumOfEngineVersionForcreateApplicationInputSparkVersion322 = "SPARK_VERSION_3_2_2"
)
