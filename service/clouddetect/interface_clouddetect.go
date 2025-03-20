// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package clouddetectiface provides an interface to enable mocking the CLOUD_DETECT service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package clouddetect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// CLOUDDETECTAPI provides an interface to enable mocking the
// clouddetect.CLOUDDETECT service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // CLOUD_DETECT.
//    func myFunc(svc CLOUDDETECTAPI) bool {
//        // Make svc.CreateTask request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := clouddetect.New(sess)
//
//        myFunc(svc)
//    }
//
type CLOUDDETECTAPI interface {
	CreateTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateTask(*CreateTaskInput) (*CreateTaskOutput, error)
	CreateTaskWithContext(volcengine.Context, *CreateTaskInput, ...request.Option) (*CreateTaskOutput, error)
	CreateTaskRequest(*CreateTaskInput) (*request.Request, *CreateTaskOutput)

	DeleteTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteTask(*DeleteTaskInput) (*DeleteTaskOutput, error)
	DeleteTaskWithContext(volcengine.Context, *DeleteTaskInput, ...request.Option) (*DeleteTaskOutput, error)
	DeleteTaskRequest(*DeleteTaskInput) (*request.Request, *DeleteTaskOutput)

	GetTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetTask(*GetTaskInput) (*GetTaskOutput, error)
	GetTaskWithContext(volcengine.Context, *GetTaskInput, ...request.Option) (*GetTaskOutput, error)
	GetTaskRequest(*GetTaskInput) (*request.Request, *GetTaskOutput)

	GetTaskResultCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetTaskResultCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetTaskResultCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetTaskResult(*GetTaskResultInput) (*GetTaskResultOutput, error)
	GetTaskResultWithContext(volcengine.Context, *GetTaskResultInput, ...request.Option) (*GetTaskResultOutput, error)
	GetTaskResultRequest(*GetTaskResultInput) (*request.Request, *GetTaskResultOutput)

	ListNodesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListNodesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListNodesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListNodes(*ListNodesInput) (*ListNodesOutput, error)
	ListNodesWithContext(volcengine.Context, *ListNodesInput, ...request.Option) (*ListNodesOutput, error)
	ListNodesRequest(*ListNodesInput) (*request.Request, *ListNodesOutput)

	ListTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTask(*ListTaskInput) (*ListTaskOutput, error)
	ListTaskWithContext(volcengine.Context, *ListTaskInput, ...request.Option) (*ListTaskOutput, error)
	ListTaskRequest(*ListTaskInput) (*request.Request, *ListTaskOutput)

	ListTaskGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTaskGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTaskGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTaskGroups(*ListTaskGroupsInput) (*ListTaskGroupsOutput, error)
	ListTaskGroupsWithContext(volcengine.Context, *ListTaskGroupsInput, ...request.Option) (*ListTaskGroupsOutput, error)
	ListTaskGroupsRequest(*ListTaskGroupsInput) (*request.Request, *ListTaskGroupsOutput)

	RestartTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartTask(*RestartTaskInput) (*RestartTaskOutput, error)
	RestartTaskWithContext(volcengine.Context, *RestartTaskInput, ...request.Option) (*RestartTaskOutput, error)
	RestartTaskRequest(*RestartTaskInput) (*request.Request, *RestartTaskOutput)

	StopTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopTask(*StopTaskInput) (*StopTaskOutput, error)
	StopTaskWithContext(volcengine.Context, *StopTaskInput, ...request.Option) (*StopTaskOutput, error)
	StopTaskRequest(*StopTaskInput) (*request.Request, *StopTaskOutput)
}

var _ CLOUDDETECTAPI = (*CLOUDDETECT)(nil)
