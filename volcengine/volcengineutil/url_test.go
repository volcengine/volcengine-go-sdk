package volcengineutil

import "testing"

func TestGetDefaultEndpointByServiceInfoRejectsInjection(t *testing.T) {
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

	if got := *GetDefaultEndpointByServiceInfo("ecs", "cn-beijing", bootstrap, nil, nil); got != "ecs.cn-beijing.volcengineapi.com" {
		t.Fatalf("unexpected endpoint %s", got)
	}

	for region := range bootstrap {
		if region == "cn-beijing" {
			continue
		}
		if got := *GetDefaultEndpointByServiceInfo("ecs", region, bootstrap, nil, nil); got != endpoint {
			t.Fatalf("region %q should fall back to default endpoint, got %s", region, got)
		}
	}

	if got := *GetDefaultEndpointByServiceInfo("ecs@127.0.0.1", "cn-beijing", bootstrap, nil, nil); got != endpoint {
		t.Fatalf("invalid service should fall back to default endpoint, got %s", got)
	}
}
