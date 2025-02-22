// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cpiface provides an interface to enable mocking the CP service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// CPAPI provides an interface to enable mocking the
// cp.CP service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // CP.
//    func myFunc(svc CPAPI) bool {
//        // Make svc.CancelPipelineRun request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cp.New(sess)
//
//        myFunc(svc)
//    }
//
type CPAPI interface {
	CancelPipelineRunCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelPipelineRunCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelPipelineRunCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelPipelineRun(*CancelPipelineRunInput) (*CancelPipelineRunOutput, error)
	CancelPipelineRunWithContext(volcengine.Context, *CancelPipelineRunInput, ...request.Option) (*CancelPipelineRunOutput, error)
	CancelPipelineRunRequest(*CancelPipelineRunInput) (*request.Request, *CancelPipelineRunOutput)

	CreateComponentStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateComponentStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateComponentStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateComponentStep(*CreateComponentStepInput) (*CreateComponentStepOutput, error)
	CreateComponentStepWithContext(volcengine.Context, *CreateComponentStepInput, ...request.Option) (*CreateComponentStepOutput, error)
	CreateComponentStepRequest(*CreateComponentStepInput) (*request.Request, *CreateComponentStepOutput)

	CreatePipelineWebhookURLCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreatePipelineWebhookURLCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreatePipelineWebhookURLCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreatePipelineWebhookURL(*CreatePipelineWebhookURLInput) (*CreatePipelineWebhookURLOutput, error)
	CreatePipelineWebhookURLWithContext(volcengine.Context, *CreatePipelineWebhookURLInput, ...request.Option) (*CreatePipelineWebhookURLOutput, error)
	CreatePipelineWebhookURLRequest(*CreatePipelineWebhookURLInput) (*request.Request, *CreatePipelineWebhookURLOutput)

	CreateResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateResource(*CreateResourceInput) (*CreateResourceOutput, error)
	CreateResourceWithContext(volcengine.Context, *CreateResourceInput, ...request.Option) (*CreateResourceOutput, error)
	CreateResourceRequest(*CreateResourceInput) (*request.Request, *CreateResourceOutput)

	CreateTriggerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTriggerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTriggerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateTrigger(*CreateTriggerInput) (*CreateTriggerOutput, error)
	CreateTriggerWithContext(volcengine.Context, *CreateTriggerInput, ...request.Option) (*CreateTriggerOutput, error)
	CreateTriggerRequest(*CreateTriggerInput) (*request.Request, *CreateTriggerOutput)

	CreateWorkspaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateWorkspaceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateWorkspaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateWorkspace(*CreateWorkspaceInput) (*CreateWorkspaceOutput, error)
	CreateWorkspaceWithContext(volcengine.Context, *CreateWorkspaceInput, ...request.Option) (*CreateWorkspaceOutput, error)
	CreateWorkspaceRequest(*CreateWorkspaceInput) (*request.Request, *CreateWorkspaceOutput)

	DeleteComponentStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteComponentStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteComponentStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteComponentStep(*DeleteComponentStepInput) (*DeleteComponentStepOutput, error)
	DeleteComponentStepWithContext(volcengine.Context, *DeleteComponentStepInput, ...request.Option) (*DeleteComponentStepOutput, error)
	DeleteComponentStepRequest(*DeleteComponentStepInput) (*request.Request, *DeleteComponentStepOutput)

	DeletePipelineCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeletePipelineCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeletePipelineCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeletePipeline(*DeletePipelineInput) (*DeletePipelineOutput, error)
	DeletePipelineWithContext(volcengine.Context, *DeletePipelineInput, ...request.Option) (*DeletePipelineOutput, error)
	DeletePipelineRequest(*DeletePipelineInput) (*request.Request, *DeletePipelineOutput)

	DeleteResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteResource(*DeleteResourceInput) (*DeleteResourceOutput, error)
	DeleteResourceWithContext(volcengine.Context, *DeleteResourceInput, ...request.Option) (*DeleteResourceOutput, error)
	DeleteResourceRequest(*DeleteResourceInput) (*request.Request, *DeleteResourceOutput)

	DeleteTriggerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteTriggerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteTriggerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteTrigger(*DeleteTriggerInput) (*DeleteTriggerOutput, error)
	DeleteTriggerWithContext(volcengine.Context, *DeleteTriggerInput, ...request.Option) (*DeleteTriggerOutput, error)
	DeleteTriggerRequest(*DeleteTriggerInput) (*request.Request, *DeleteTriggerOutput)

	DeleteWorkspaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteWorkspaceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteWorkspaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteWorkspace(*DeleteWorkspaceInput) (*DeleteWorkspaceOutput, error)
	DeleteWorkspaceWithContext(volcengine.Context, *DeleteWorkspaceInput, ...request.Option) (*DeleteWorkspaceOutput, error)
	DeleteWorkspaceRequest(*DeleteWorkspaceInput) (*request.Request, *DeleteWorkspaceOutput)

	GetTaskRunLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetTaskRunLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetTaskRunLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetTaskRunLog(*GetTaskRunLogInput) (*GetTaskRunLogOutput, error)
	GetTaskRunLogWithContext(volcengine.Context, *GetTaskRunLogInput, ...request.Option) (*GetTaskRunLogOutput, error)
	GetTaskRunLogRequest(*GetTaskRunLogInput) (*request.Request, *GetTaskRunLogOutput)

	GetTaskRunLogDownloadURICommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetTaskRunLogDownloadURICommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetTaskRunLogDownloadURICommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetTaskRunLogDownloadURI(*GetTaskRunLogDownloadURIInput) (*GetTaskRunLogDownloadURIOutput, error)
	GetTaskRunLogDownloadURIWithContext(volcengine.Context, *GetTaskRunLogDownloadURIInput, ...request.Option) (*GetTaskRunLogDownloadURIOutput, error)
	GetTaskRunLogDownloadURIRequest(*GetTaskRunLogDownloadURIInput) (*request.Request, *GetTaskRunLogDownloadURIOutput)

	ListComponentStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListComponentStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListComponentStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListComponentStep(*ListComponentStepInput) (*ListComponentStepOutput, error)
	ListComponentStepWithContext(volcengine.Context, *ListComponentStepInput, ...request.Option) (*ListComponentStepOutput, error)
	ListComponentStepRequest(*ListComponentStepInput) (*request.Request, *ListComponentStepOutput)

	ListPipelineRunsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListPipelineRunsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListPipelineRunsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListPipelineRuns(*ListPipelineRunsInput) (*ListPipelineRunsOutput, error)
	ListPipelineRunsWithContext(volcengine.Context, *ListPipelineRunsInput, ...request.Option) (*ListPipelineRunsOutput, error)
	ListPipelineRunsRequest(*ListPipelineRunsInput) (*request.Request, *ListPipelineRunsOutput)

	ListPipelinesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListPipelinesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListPipelinesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListPipelines(*ListPipelinesInput) (*ListPipelinesOutput, error)
	ListPipelinesWithContext(volcengine.Context, *ListPipelinesInput, ...request.Option) (*ListPipelinesOutput, error)
	ListPipelinesRequest(*ListPipelinesInput) (*request.Request, *ListPipelinesOutput)

	ListResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListResourcesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListResources(*ListResourcesInput) (*ListResourcesOutput, error)
	ListResourcesWithContext(volcengine.Context, *ListResourcesInput, ...request.Option) (*ListResourcesOutput, error)
	ListResourcesRequest(*ListResourcesInput) (*request.Request, *ListResourcesOutput)

	ListTaskRunsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTaskRunsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTaskRunsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTaskRuns(*ListTaskRunsInput) (*ListTaskRunsOutput, error)
	ListTaskRunsWithContext(volcengine.Context, *ListTaskRunsInput, ...request.Option) (*ListTaskRunsOutput, error)
	ListTaskRunsRequest(*ListTaskRunsInput) (*request.Request, *ListTaskRunsOutput)

	ListTriggersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTriggersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTriggersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTriggers(*ListTriggersInput) (*ListTriggersOutput, error)
	ListTriggersWithContext(volcengine.Context, *ListTriggersInput, ...request.Option) (*ListTriggersOutput, error)
	ListTriggersRequest(*ListTriggersInput) (*request.Request, *ListTriggersOutput)

	ListWorkspacesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListWorkspacesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListWorkspacesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListWorkspaces(*ListWorkspacesInput) (*ListWorkspacesOutput, error)
	ListWorkspacesWithContext(volcengine.Context, *ListWorkspacesInput, ...request.Option) (*ListWorkspacesOutput, error)
	ListWorkspacesRequest(*ListWorkspacesInput) (*request.Request, *ListWorkspacesOutput)

	RunPipelineCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RunPipelineCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RunPipelineCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RunPipeline(*RunPipelineInput) (*RunPipelineOutput, error)
	RunPipelineWithContext(volcengine.Context, *RunPipelineInput, ...request.Option) (*RunPipelineOutput, error)
	RunPipelineRequest(*RunPipelineInput) (*request.Request, *RunPipelineOutput)

	UpdateComponentStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateComponentStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateComponentStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateComponentStep(*UpdateComponentStepInput) (*UpdateComponentStepOutput, error)
	UpdateComponentStepWithContext(volcengine.Context, *UpdateComponentStepInput, ...request.Option) (*UpdateComponentStepOutput, error)
	UpdateComponentStepRequest(*UpdateComponentStepInput) (*request.Request, *UpdateComponentStepOutput)

	UpdateResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateResource(*UpdateResourceInput) (*UpdateResourceOutput, error)
	UpdateResourceWithContext(volcengine.Context, *UpdateResourceInput, ...request.Option) (*UpdateResourceOutput, error)
	UpdateResourceRequest(*UpdateResourceInput) (*request.Request, *UpdateResourceOutput)

	UpdateTriggerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateTriggerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateTriggerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateTrigger(*UpdateTriggerInput) (*UpdateTriggerOutput, error)
	UpdateTriggerWithContext(volcengine.Context, *UpdateTriggerInput, ...request.Option) (*UpdateTriggerOutput, error)
	UpdateTriggerRequest(*UpdateTriggerInput) (*request.Request, *UpdateTriggerOutput)

	UpdateWorkspaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateWorkspaceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateWorkspaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateWorkspace(*UpdateWorkspaceInput) (*UpdateWorkspaceOutput, error)
	UpdateWorkspaceWithContext(volcengine.Context, *UpdateWorkspaceInput, ...request.Option) (*UpdateWorkspaceOutput, error)
	UpdateWorkspaceRequest(*UpdateWorkspaceInput) (*request.Request, *UpdateWorkspaceOutput)

	UpgradeComponentStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpgradeComponentStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpgradeComponentStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpgradeComponentStep(*UpgradeComponentStepInput) (*UpgradeComponentStepOutput, error)
	UpgradeComponentStepWithContext(volcengine.Context, *UpgradeComponentStepInput, ...request.Option) (*UpgradeComponentStepOutput, error)
	UpgradeComponentStepRequest(*UpgradeComponentStepInput) (*request.Request, *UpgradeComponentStepOutput)
}

var _ CPAPI = (*CP)(nil)
