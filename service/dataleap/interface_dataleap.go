// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dataleapiface provides an interface to enable mocking the DATALEAP service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dataleap

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// DATALEAPAPI provides an interface to enable mocking the
// dataleap.DATALEAP service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // DATALEAP.
//    func myFunc(svc DATALEAPAPI) bool {
//        // Make svc.DTSOpenDescribeResourceGroups request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := dataleap.New(sess)
//
//        myFunc(svc)
//    }
//
type DATALEAPAPI interface {
	DTSOpenDescribeResourceGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DTSOpenDescribeResourceGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DTSOpenDescribeResourceGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DTSOpenDescribeResourceGroups(*DTSOpenDescribeResourceGroupsInput) (*DTSOpenDescribeResourceGroupsOutput, error)
	DTSOpenDescribeResourceGroupsWithContext(volcengine.Context, *DTSOpenDescribeResourceGroupsInput, ...request.Option) (*DTSOpenDescribeResourceGroupsOutput, error)
	DTSOpenDescribeResourceGroupsRequest(*DTSOpenDescribeResourceGroupsInput) (*request.Request, *DTSOpenDescribeResourceGroupsOutput)

	DTSOpenListTagResourceGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DTSOpenListTagResourceGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DTSOpenListTagResourceGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DTSOpenListTagResourceGroups(*DTSOpenListTagResourceGroupsInput) (*DTSOpenListTagResourceGroupsOutput, error)
	DTSOpenListTagResourceGroupsWithContext(volcengine.Context, *DTSOpenListTagResourceGroupsInput, ...request.Option) (*DTSOpenListTagResourceGroupsOutput, error)
	DTSOpenListTagResourceGroupsRequest(*DTSOpenListTagResourceGroupsInput) (*request.Request, *DTSOpenListTagResourceGroupsOutput)

	DTSOpenTagResourceGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DTSOpenTagResourceGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DTSOpenTagResourceGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DTSOpenTagResourceGroups(*DTSOpenTagResourceGroupsInput) (*DTSOpenTagResourceGroupsOutput, error)
	DTSOpenTagResourceGroupsWithContext(volcengine.Context, *DTSOpenTagResourceGroupsInput, ...request.Option) (*DTSOpenTagResourceGroupsOutput, error)
	DTSOpenTagResourceGroupsRequest(*DTSOpenTagResourceGroupsInput) (*request.Request, *DTSOpenTagResourceGroupsOutput)

	DTSOpenUntagResourceGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DTSOpenUntagResourceGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DTSOpenUntagResourceGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DTSOpenUntagResourceGroups(*DTSOpenUntagResourceGroupsInput) (*DTSOpenUntagResourceGroupsOutput, error)
	DTSOpenUntagResourceGroupsWithContext(volcengine.Context, *DTSOpenUntagResourceGroupsInput, ...request.Option) (*DTSOpenUntagResourceGroupsOutput, error)
	DTSOpenUntagResourceGroupsRequest(*DTSOpenUntagResourceGroupsInput) (*request.Request, *DTSOpenUntagResourceGroupsOutput)
}

var _ DATALEAPAPI = (*DATALEAP)(nil)
