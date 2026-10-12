package endpoints

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// ErrCodeInvalidRegion is the error code returned when a region contains
// characters that are not allowed to be part of a host name label.
const ErrCodeInvalidRegion = "InvalidRegion"

// ErrCodeInvalidService is the error code returned when a service contains
// characters that are not allowed to be part of a host name label.
const ErrCodeInvalidService = "InvalidService"

// isEndpointLabelChar reports whether c can be used inside a single host name
// label built from the region or service. Only ASCII letters, digits, '-' and
// '_' are allowed. Anything else, such as '@', ':', '/', '?', '#', '.', white
// spaces or control characters, could change the structure of the resolved
// host and is therefore rejected to prevent SSRF through endpoint injection.
func isEndpointLabelChar(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z':
		return true
	case c >= 'A' && c <= 'Z':
		return true
	case c >= '0' && c <= '9':
		return true
	case c == '-' || c == '_':
		return true
	}
	return false
}

// validateEndpointLabel returns the offset of the first illegal byte in s, or
// -1 if s is valid. An empty string is considered valid, the caller decides
// how to handle a missing value.
func validateEndpointLabel(s string) int {
	for i := 0; i < len(s); i++ {
		if !isEndpointLabelChar(s[i]) {
			return i
		}
	}
	return -1
}

// ValidateRegion checks that region is safe to be used when building an
// endpoint host name. An empty region is allowed and is left to the caller.
func ValidateRegion(region string) error {
	if i := validateEndpointLabel(region); i >= 0 {
		return volcengineerr.New(ErrCodeInvalidRegion,
			fmt.Sprintf("region %q contains illegal character at index %d, "+
				"only letters, digits, '-' and '_' are allowed", region, i), nil)
	}
	return nil
}

// ValidateService checks that service is safe to be used when building an
// endpoint host name. An empty service is allowed and is left to the caller.
func ValidateService(service string) error {
	if i := validateEndpointLabel(service); i >= 0 {
		return volcengineerr.New(ErrCodeInvalidService,
			fmt.Sprintf("service %q contains illegal character at index %d, "+
				"only letters, digits, '-' and '_' are allowed", service, i), nil)
	}
	return nil
}

// ValidateServiceAndRegion validates both service and region.
func ValidateServiceAndRegion(service, region string) error {
	if err := ValidateService(service); err != nil {
		return err
	}
	return ValidateRegion(region)
}
