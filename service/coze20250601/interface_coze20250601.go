// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package coze20250601iface provides an interface to enable mocking the COZE20250601 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package coze20250601

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// COZE20250601API provides an interface to enable mocking the
// coze20250601.COZE20250601 service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // COZE20250601.
//    func myFunc(svc COZE20250601API) bool {
//        // Make svc.AuthorizeCozeToUser request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := coze20250601.New(sess)
//
//        myFunc(svc)
//    }
//
type COZE20250601API interface {
	AuthorizeCozeToUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AuthorizeCozeToUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AuthorizeCozeToUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AuthorizeCozeToUser(*AuthorizeCozeToUserInput) (*AuthorizeCozeToUserOutput, error)
	AuthorizeCozeToUserWithContext(volcengine.Context, *AuthorizeCozeToUserInput, ...request.Option) (*AuthorizeCozeToUserOutput, error)
	AuthorizeCozeToUserRequest(*AuthorizeCozeToUserInput) (*request.Request, *AuthorizeCozeToUserOutput)

	AuthorizeVolcToUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AuthorizeVolcToUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AuthorizeVolcToUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AuthorizeVolcToUser(*AuthorizeVolcToUserInput) (*AuthorizeVolcToUserOutput, error)
	AuthorizeVolcToUserWithContext(volcengine.Context, *AuthorizeVolcToUserInput, ...request.Option) (*AuthorizeVolcToUserOutput, error)
	AuthorizeVolcToUserRequest(*AuthorizeVolcToUserInput) (*request.Request, *AuthorizeVolcToUserOutput)

	CreateUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateUser(*CreateUserInput) (*CreateUserOutput, error)
	CreateUserWithContext(volcengine.Context, *CreateUserInput, ...request.Option) (*CreateUserOutput, error)
	CreateUserRequest(*CreateUserInput) (*request.Request, *CreateUserOutput)

	ListCozeUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListCozeUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListCozeUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListCozeUser(*ListCozeUserInput) (*ListCozeUserOutput, error)
	ListCozeUserWithContext(volcengine.Context, *ListCozeUserInput, ...request.Option) (*ListCozeUserOutput, error)
	ListCozeUserRequest(*ListCozeUserInput) (*request.Request, *ListCozeUserOutput)
}

var _ COZE20250601API = (*COZE20250601)(nil)
