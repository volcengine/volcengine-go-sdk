package volcstack

import "context"

// BackgroundContext returns a context that will never be canceled, has no
// values, and no deadline. This context is used by the SDK to provide
// backwards compatibility with non-context API operations and functionality.
// See https://golang.org/pkg/context for more information on Contexts.
func BackgroundContext() Context {
	return context.Background()
}
