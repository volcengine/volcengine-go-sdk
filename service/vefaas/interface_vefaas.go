// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package vefaasiface provides an interface to enable mocking the VEFAAS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// VEFAASAPI provides an interface to enable mocking the
// vefaas.VEFAAS service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // VEFAAS.
//    func myFunc(svc VEFAASAPI) bool {
//        // Make svc.CreateFunction request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := vefaas.New(sess)
//
//        myFunc(svc)
//    }
//
type VEFAASAPI interface {
	CreateFunctionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateFunctionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateFunctionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateFunction(*CreateFunctionInput) (*CreateFunctionOutput, error)
	CreateFunctionWithContext(volcengine.Context, *CreateFunctionInput, ...request.Option) (*CreateFunctionOutput, error)
	CreateFunctionRequest(*CreateFunctionInput) (*request.Request, *CreateFunctionOutput)

	DeleteFunctionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteFunctionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteFunctionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteFunction(*DeleteFunctionInput) (*DeleteFunctionOutput, error)
	DeleteFunctionWithContext(volcengine.Context, *DeleteFunctionInput, ...request.Option) (*DeleteFunctionOutput, error)
	DeleteFunctionRequest(*DeleteFunctionInput) (*request.Request, *DeleteFunctionOutput)

	GetFunctionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetFunctionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetFunctionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetFunction(*GetFunctionInput) (*GetFunctionOutput, error)
	GetFunctionWithContext(volcengine.Context, *GetFunctionInput, ...request.Option) (*GetFunctionOutput, error)
	GetFunctionRequest(*GetFunctionInput) (*request.Request, *GetFunctionOutput)

	UpdateFunctionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateFunctionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateFunctionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateFunction(*UpdateFunctionInput) (*UpdateFunctionOutput, error)
	UpdateFunctionWithContext(volcengine.Context, *UpdateFunctionInput, ...request.Option) (*UpdateFunctionOutput, error)
	UpdateFunctionRequest(*UpdateFunctionInput) (*request.Request, *UpdateFunctionOutput)
}

var _ VEFAASAPI = (*VEFAAS)(nil)
