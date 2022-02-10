package corehandlers

import (
	"os"
	"runtime"

	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
)

// SDKVersionUserAgentHandler is a request handler for adding the SDK Version
// to the user agent.
var SDKVersionUserAgentHandler = request.NamedHandler{
	Name: "core.SDKVersionUserAgentHandler",
	Fn: request.MakeAddToUserAgentHandler(volcstack.SDKName, volcstack.SDKVersion,
		runtime.Version(), runtime.GOOS, runtime.GOARCH),
}

const execEnvVar = `VOLCSTACK_EXECUTION_ENV`
const execEnvUAKey = `exec-env`

// AddHostExecEnvUserAgentHandler is a request handler appending the SDK's
// execution environment to the user agent.
//
// If the environment variable VOLCSTACK_EXECUTION_ENV is set, its value will be
// appended to the user agent string.
var AddHostExecEnvUserAgentHandler = request.NamedHandler{
	Name: "core.AddHostExecEnvUserAgentHandler",
	Fn: func(r *request.Request) {
		v := os.Getenv(execEnvVar)
		if len(v) == 0 {
			return
		}

		request.AddToUserAgent(r, execEnvUAKey+"/"+v)
	},
}
