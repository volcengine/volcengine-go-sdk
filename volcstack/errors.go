package volcstack

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/aws/errors.go

import "github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"

var (
	// ErrMissingRegion is an error that is returned if region configuration is
	// not found.
	ErrMissingRegion = volcstackerr.New("MissingRegion", "could not find region configuration", nil)

	// ErrMissingEndpoint is an error that is returned if an endpoint cannot be
	// resolved for a service.
	ErrMissingEndpoint = volcstackerr.New("MissingEndpoint", "'Endpoint' configuration is required for this service", nil)
)
