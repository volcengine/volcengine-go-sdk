// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ntaiface provides an interface to enable mocking the NTA service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package nta

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// NTAAPI provides an interface to enable mocking the
// nta.NTA service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // NTA.
//    func myFunc(svc NTAAPI) bool {
//        // Make svc.CreateFileDetection request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := nta.New(sess)
//
//        myFunc(svc)
//    }
//
type NTAAPI interface {
	CreateFileDetectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateFileDetectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateFileDetectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateFileDetection(*CreateFileDetectionInput) (*CreateFileDetectionOutput, error)
	CreateFileDetectionWithContext(volcengine.Context, *CreateFileDetectionInput, ...request.Option) (*CreateFileDetectionOutput, error)
	CreateFileDetectionRequest(*CreateFileDetectionInput) (*request.Request, *CreateFileDetectionOutput)

	GetFileDetectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetFileDetectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetFileDetectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetFileDetection(*GetFileDetectionInput) (*GetFileDetectionOutput, error)
	GetFileDetectionWithContext(volcengine.Context, *GetFileDetectionInput, ...request.Option) (*GetFileDetectionOutput, error)
	GetFileDetectionRequest(*GetFileDetectionInput) (*request.Request, *GetFileDetectionOutput)
}

var _ NTAAPI = (*NTA)(nil)
