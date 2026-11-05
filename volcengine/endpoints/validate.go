package endpoints

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// ErrCodeInvalidRegion is the error code returned when a region is not a
// valid host name label.
const ErrCodeInvalidRegion = "InvalidRegion"

// ErrCodeInvalidService is the error code returned when a service is not a
// valid host name label.
const ErrCodeInvalidService = "InvalidService"

// endpointLabelRule describes the format accepted for region and service, it
// is used in error messages.
const endpointLabelRule = "must start with a letter, end with a letter or digit, " +
	"and contain only letters, digits, '-' and '_'"

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// isEndpointLabelChar reports whether c can be used inside a single host name
// label built from the region or service. Only ASCII letters, digits, '-' and
// '_' are allowed. Anything else, such as '@', ':', '/', '?', '#', '.', white
// spaces or control characters, could change the structure of the resolved
// host and is therefore rejected to prevent SSRF through endpoint injection.
func isEndpointLabelChar(c byte) bool {
	return isLetter(c) || isDigit(c) || c == '-' || c == '_'
}

// validateEndpointLabel checks that s is a valid endpoint label, equivalent to
// the regular expression ^[A-Za-z]([A-Za-z0-9_-]*[A-Za-z0-9])?$ :
//   - it only contains ASCII letters, digits, '-' and '_'
//   - it starts with a letter
//   - it does not end with '-' or '_'
//
// It returns an empty string if s is valid, or the reason why it is not. An
// empty string is considered valid, the caller decides how to handle a
// missing value.
func validateEndpointLabel(s string) string {
	if len(s) == 0 {
		return ""
	}
	// Check the character set first, it is the one that prevents SSRF.
	for i := 0; i < len(s); i++ {
		if !isEndpointLabelChar(s[i]) {
			return fmt.Sprintf("illegal character at index %d", i)
		}
	}
	if !isLetter(s[0]) {
		return "must start with a letter"
	}
	if last := s[len(s)-1]; last == '-' || last == '_' {
		return "must not end with '-' or '_'"
	}
	return ""
}

// ValidateRegion checks that region is safe to be used when building an
// endpoint host name. An empty region is allowed and is left to the caller.
func ValidateRegion(region string) error {
	if reason := validateEndpointLabel(region); reason != "" {
		return volcengineerr.New(ErrCodeInvalidRegion,
			fmt.Sprintf("invalid region %q: %s, region %s", region, reason, endpointLabelRule), nil)
	}
	return nil
}

// ValidateService checks that service is safe to be used when building an
// endpoint host name. An empty service is allowed and is left to the caller.
func ValidateService(service string) error {
	if reason := validateEndpointLabel(service); reason != "" {
		return volcengineerr.New(ErrCodeInvalidService,
			fmt.Sprintf("invalid service %q: %s, service %s", service, reason, endpointLabelRule), nil)
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
