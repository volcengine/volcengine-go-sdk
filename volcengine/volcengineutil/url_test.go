package volcengineutil

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
)

// captureLogger records debug messages so a test can tell which branch of
// GetDefaultEndpointByServiceInfo produced the result.
type captureLogger struct {
	msgs []string
}

func (l *captureLogger) record(args ...interface{}) { l.msgs = append(l.msgs, fmt.Sprint(args...)) }

func (l *captureLogger) Log(args ...interface{})   { l.record(args...) }
func (l *captureLogger) Debug(args ...interface{}) { l.record(args...) }
func (l *captureLogger) DebugByLevel(_ volcengine.LogLevelType, args ...interface{}) {
	l.record(args...)
}
func (l *captureLogger) SetDebug(*bool)                            {}
func (l *captureLogger) SetDebugLogLevel(*volcengine.LogLevelType) {}
func (l *captureLogger) Warn(args ...interface{})                  { l.record(args...) }
func (l *captureLogger) Error(args ...interface{})                 { l.record(args...) }
func (l *captureLogger) SetIoWriter(io.Writer)                     {}

func (l *captureLogger) contains(sub string) bool {
	for _, m := range l.msgs {
		if strings.Contains(m, sub) {
			return true
		}
	}
	return false
}

func TestGetDefaultEndpointByServiceInfoRejectsInjection(t *testing.T) {
	// Pass DualStack explicitly, nil would read VOLC_ENABLE_DUALSTACK from the
	// environment and change the expected suffix.
	noDualStack := volcengine.Bool(false)

	bootstrap := map[string]struct{}{
		"cn-beijing@127.0.0.1": {},
		"cn-beijing:443":       {},
		"cn/beijing":           {},
		"cn?beijing":           {},
		"cn#beijing":           {},
		"cn beijing":           {},
		"cn\tbeijing":          {},
		"cn\nbeijing":          {},
		"cn\x00beijing":        {},
		"evil.com":             {},
		"cn-beijing":           {},
	}

	if got := *GetDefaultEndpointByServiceInfo("ecs", "cn-beijing", bootstrap, noDualStack, nil); got != "ecs.cn-beijing.volcengineapi.com" {
		t.Fatalf("unexpected endpoint %s", got)
	}

	for region := range bootstrap {
		if region == "cn-beijing" {
			continue
		}
		if got := *GetDefaultEndpointByServiceInfo("ecs", region, bootstrap, noDualStack, nil); got != endpoint {
			t.Fatalf("region %q should fall back to default endpoint, got %s", region, got)
		}
	}

	// An invalid service is never a key of defaultEndpoint, so it would also
	// fall back through the "service not found" branch. Check the log to make
	// sure it is the label check which rejected it.
	for _, svc := range []string{"ecs@127.0.0.1", "ecs.evil.com#", "ecs/x", "1ecs", "ecs-"} {
		logger := &captureLogger{}
		if got := *GetDefaultEndpointByServiceInfo(svc, "cn-beijing", bootstrap, noDualStack, logger); got != endpoint {
			t.Fatalf("service %q should fall back to default endpoint, got %s", svc, got)
		}
		if !logger.contains(fmt.Sprintf("invalid service %q", svc)) {
			t.Fatalf("service %q should be rejected by the label check, logs: %v", svc, logger.msgs)
		}
		if logger.contains("not found in default endpoint map") {
			t.Fatalf("service %q should be rejected before the map lookup, logs: %v", svc, logger.msgs)
		}
	}
}

// All built-in services and bootstrap regions must pass the endpoint label
// validation, otherwise they would silently fall back to the default endpoint.
func TestBuiltinServicesAndRegionsAreValidLabels(t *testing.T) {
	for name, info := range defaultEndpoint {
		if err := endpoints.ValidateService(name); err != nil {
			t.Errorf("built-in service key %q is invalid: %v", name, err)
		}
		if err := endpoints.ValidateService(info.Service); err != nil {
			t.Errorf("built-in service %q is invalid: %v", info.Service, err)
		}
		for region := range info.RegionEndpointMap {
			if err := endpoints.ValidateRegion(region); err != nil {
				t.Errorf("region %q of service %q is invalid: %v", region, name, err)
			}
		}
	}
	for region := range bootstrapRegion {
		if err := endpoints.ValidateRegion(region); err != nil {
			t.Errorf("bootstrap region %q is invalid: %v", region, err)
		}
	}
}
