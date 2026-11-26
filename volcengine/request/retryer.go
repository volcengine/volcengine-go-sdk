package request

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"errors"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// Retryer provides the interface drive the SDK's request retry behavior. The
// Retryer implementation is responsible for implementing exponential backoff,
// and determine if a request API error should be retried.
//
// client.DefaultRetryer is the SDK's default implementation of the Retryer. It
// uses the which uses the Request.IsErrorRetryable and Request.IsErrorThrottle
// methods to determine if the request is retried.
type Retryer interface {
	// RetryRules return the retry delay that should be used by the SDK before
	// making another request attempt for the failed request.
	RetryRules(*Request) time.Duration

	// ShouldRetry returns if the failed request is retryable.
	//
	// Implementations may consider request attempt count when determining if a
	// request is retryable, but the SDK will use MaxRetries to limit the
	// number of attempts a request are made.
	ShouldRetry(*Request) bool

	// MaxRetries is the number of times a request may be retried before
	// failing.
	MaxRetries() int
}

// WithRetryer sets a Retryer value to the given Config returning the Config
// value for chaining.
func WithRetryer(cfg *volcengine.Config, retryer Retryer) *volcengine.Config {
	cfg.Retryer = retryer
	return cfg
}

// retryableCodes is a collection of service response codes which are retry-able
// without any further action.
var retryableCodes = map[string]struct{}{
	"RequestError":            {},
	"RequestTimeout":          {},
	ErrCodeResponseTimeout:    {},
	"RequestTimeoutException": {}, // Glacier's flavor of RequestTimeout
}

var throttleCodes = map[string]struct{}{
	"ProvisionedThroughputExceededException": {},
	"Throttling":                             {},
	"ThrottlingException":                    {},
	"RequestLimitExceeded":                   {},
	"RequestThrottled":                       {},
	"RequestThrottledException":              {},
	"TooManyRequestsException":               {}, // Lambda functions
	"PriorRequestNotComplete":                {}, // Route53
	"TransactionInProgressException":         {},
}

// credsExpiredCodes is a collection of error codes which signify the credentials
// need to be refreshed. Expired tokens require refreshing of credentials, and
// resigning before the request can be retried.
var credsExpiredCodes = map[string]struct{}{
	"ExpiredToken":          {},
	"ExpiredTokenException": {},
	"RequestExpired":        {}, // EC2 Only
	"InvalidAccessKey":      {},
}

func isCodeThrottle(code string) bool {
	_, ok := throttleCodes[code]
	return ok
}

func isCodeRetryable(code string) bool {
	if _, ok := retryableCodes[code]; ok {
		return true
	}

	return isCodeExpiredCreds(code)
}

func isCodeExpiredCreds(code string) bool {
	_, ok := credsExpiredCodes[code]
	return ok
}

var validParentCodes = map[string]struct{}{
	ErrCodeSerialization: {},
	ErrCodeRead:          {},
}

// IsErrorRetryable returns whether the error is retryable, based on its Code.
// Returns false if error is nil.
func IsErrorRetryable(err error) bool {
	if err == nil {
		return false
	}
	return shouldRetryError(err)
}

type temporary interface {
	Temporary() bool
}

func shouldRetryError(origErr error) bool {
	return retryableError(origErr, true)
}

func retryableError(origErr error, retryUnknown bool) bool {
	return (&retryErrorWalk{}).retry(origErr, retryUnknown, 0)
}

type retryErrorWalk struct {
	nodes int
}

func (w *retryErrorWalk) retry(origErr error, retryUnknown bool, depth int) bool {
	if depth >= 64 || w.nodes >= 256 {
		return false
	}
	w.nodes++

	switch err := origErr.(type) {
	case volcengineerr.Error:
		if err.Code() == CanceledErrorCode {
			return false
		}
		_, nested := validParentCodes[err.Code()]
		if nested {
			// Read and serialization errors require known retryable causes.
			retryUnknown = false
		} else if !isCodeRetryable(err.Code()) {
			// An ordinary cause must not make a terminal SDK code retryable.
			// Keep known transient causes and the RequestError fallback.
			retryUnknown = false
		}

		var shouldRetry bool
		var hasOrigErr bool
		if batch, ok := err.(volcengineerr.BatchedErrors); ok {
			// OrigErr can manufacture another batch containing the same list.
			// Inspect the actual causes instead, requiring each cause to retry.
			for _, child := range batch.OrigErrs() {
				if w.nodes >= 256 {
					hasOrigErr, shouldRetry = true, false
					break
				}
				if child == nil {
					w.nodes++
					continue
				}
				hasOrigErr = true
				shouldRetry = w.retry(child, retryUnknown, depth+1)
				if !shouldRetry {
					break
				}
			}
		} else if child := err.OrigErr(); child != nil {
			hasOrigErr = true
			shouldRetry = w.retry(child, retryUnknown, depth+1)
		}
		if nested {
			return shouldRetry
		}
		if err.Code() == "RequestError" && hasOrigErr && !shouldRetry {
			return false
		}
		if isCodeRetryable(err.Code()) {
			return true
		}
		return shouldRetry

	case *url.Error:
		if strings.Contains(err.Error(), "connection refused") {
			// Refused connections should be retried as the service may not yet
			// be running on the port. Go TCP dial considers refused
			// connections as not temporary.
			return true
		}
		// *url.Error only implements Temporary after golang 1.6 but since
		// url.Error only wraps the error:
		return w.retry(err.Err, retryUnknown, depth+1)

	case temporary:
		if netErr, ok := err.(*net.OpError); ok && netErr.Op == "dial" {
			return true
		}
		// If the error is temporary, we want to allow continuation of the
		// retry process
		return err.Temporary() || isErrConnectionReset(origErr)

	case nil:
		return retryUnknown

	default:
		switch err.Error() {
		case "net/http: request canceled",
			"net/http: request canceled while waiting for connection":
			// known 1.5 error case when an http request is cancelled
			return false
		}
		if unwrapped := errors.Unwrap(origErr); unwrapped != nil {
			return w.retry(unwrapped, retryUnknown, depth+1)
		}
		// Preserve the fallback for native errors and retryable SDK codes.
		return retryUnknown || isErrConnectionReset(origErr)
	}
}

// IsErrorThrottle returns whether the error is to be throttled based on its code.
// Returns false if error is nil.
func IsErrorThrottle(err error) bool {
	var aerr volcengineerr.Error
	if errors.As(err, &aerr) && aerr != nil {
		return isCodeThrottle(aerr.Code())
	}
	return false
}

// IsErrorExpiredCreds returns whether the error code is a credential expiry
// error. Returns false if error is nil.
func IsErrorExpiredCreds(err error) bool {
	var aerr volcengineerr.Error
	if errors.As(err, &aerr) && aerr != nil {
		return isCodeExpiredCreds(aerr.Code())
	}
	return false
}

// IsErrorRetryable returns whether the error is retryable, based on its Code.
// Returns false if the request has no Error set.
//
// Alias for the utility function IsErrorRetryable
func (r *Request) IsErrorRetryable() bool {
	logger := r.Config.Logger
	if isErrCode(r.Error, r.RetryErrorCodes) {
		logger.DebugByLevel(volcengine.LogDebugWithRequestRetries, "[Retry] retryable error code is contained in RetryErrorCodes list", r.Error)
		return true
	}

	// HTTP response status code 501 should not be retried.
	// 501 represents Not Implemented which means the request method is not
	// supported by the server and cannot be handled.
	if r.HTTPResponse != nil {
		// HTTP response status code 500 represents internal server error and
		// should be retried without any throttle.
		if r.HTTPResponse.StatusCode == 500 {
			return true
		}
	}

	if IsErrorRetryable(r.Error) {
		logger.DebugByLevel(volcengine.LogDebugWithRequestRetries, "[Retry] error is retryable", r.Error)
		return true
	}

	return false
}

// IsErrorThrottle returns whether the error is to be throttled based on its
// code. Returns false if the request has no Error set.
//
// Alias for the utility function IsErrorThrottle
func (r *Request) IsErrorThrottle() bool {
	if isErrCode(r.Error, r.ThrottleErrorCodes) {
		return true
	}

	if r.HTTPResponse != nil {
		switch r.HTTPResponse.StatusCode {
		case
			429, // error caused due to too many requests
			502, // Bad Gateway error should be throttled
			503, // caused when service is unavailable
			504: // error occurred due to gateway timeout
			return true
		}
	}

	return IsErrorThrottle(r.Error)
}

func isErrCode(err error, codes []string) bool {
	var aerr volcengineerr.Error
	if errors.As(err, &aerr) && aerr != nil {
		for _, code := range codes {
			if code == aerr.Code() {
				return true
			}
		}
	}

	return false
}

// IsErrorExpired returns whether the error code is a credential expiry error.
// Returns false if the request has no Error set.
//
// Alias for the utility function IsErrorExpiredCreds
func (r *Request) IsErrorExpired() bool {
	return IsErrorExpiredCreds(r.Error)
}
