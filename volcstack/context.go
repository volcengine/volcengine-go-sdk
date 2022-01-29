package volcstack

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/aws/context.go

import "context"

// Context is an alias of the Go stdlib's context.Context interface.
// It can be used within the SDK's API operation "WithContext" methods.
//
// See https://golang.org/pkg/context on how to use contexts.
type Context = context.Context
