package endpoints

import (
	"strings"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

var labelCases = []struct {
	name  string
	value string
	valid bool
}{
	{name: "normal", value: "cn-beijing", valid: true},
	{name: "empty deferred to caller", value: "", valid: true},
	{name: "at sign", value: "cn-beijing@127.0.0.1", valid: false},
	{name: "colon", value: "cn-beijing:443", valid: false},
	{name: "slash", value: "cn/beijing", valid: false},
	{name: "question mark", value: "cn?beijing", valid: false},
	{name: "fragment", value: "cn#beijing", valid: false},
	{name: "space", value: "cn beijing", valid: false},
	{name: "tab", value: "cn\tbeijing", valid: false},
	{name: "newline", value: "cn\nbeijing", valid: false},
	{name: "nul", value: "cn\x00beijing", valid: false},
	{name: "dot", value: "evil.com", valid: false},
}

func TestValidateRegion(t *testing.T) {
	for _, c := range labelCases {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateRegion(c.value)
			if c.valid && err != nil {
				t.Fatalf("expect valid, got err: %v", err)
			}
			if !c.valid {
				if err == nil {
					t.Fatalf("expect invalid region %q", c.value)
				}
				if e, ok := err.(volcengineerr.Error); !ok || e.Code() != ErrCodeInvalidRegion {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateService(t *testing.T) {
	for _, c := range labelCases {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateService(c.value)
			if c.valid && err != nil {
				t.Fatalf("expect valid, got err: %v", err)
			}
			if !c.valid {
				if err == nil {
					t.Fatalf("expect invalid service %q", c.value)
				}
				if e, ok := err.(volcengineerr.Error); !ok || e.Code() != ErrCodeInvalidService {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
	for _, s := range []string{"ecs", "cloud_detect", "Redis", "rds_mysql"} {
		if err := ValidateService(s); err != nil {
			t.Fatalf("service %q should be valid: %v", s, err)
		}
	}
}

func TestStandardEndpointResolverRejectsInjection(t *testing.T) {
	r := NewStandardEndpointResolver()

	got, err := r.EndpointFor("ecs", "cn-beijing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want := "ecs.cn-beijing.volcengineapi.com"; got.URL != want {
		t.Fatalf("want %s, got %s", want, got.URL)
	}

	for _, c := range labelCases {
		if c.valid {
			continue
		}
		if _, err := r.EndpointFor("ecs", c.value); err == nil {
			t.Fatalf("%s: region %q should be rejected", c.name, c.value)
		}
		if _, err := r.EndpointFor(c.value, "cn-beijing"); err == nil {
			t.Fatalf("%s: service %q should be rejected", c.name, c.value)
		}
	}
}

func TestPartitionEndpointForRejectsInjection(t *testing.T) {
	p := partition{
		ID:        "volcengine",
		DNSSuffix: "volcengineapi.com",
		Defaults:  endpoint{Hostname: "{service}.{region}.{dnsSuffix}"},
		Services:  services{"ecs": service{}},
	}

	got, err := p.EndpointFor("ecs", "cn-beijing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasSuffix(got.URL, "ecs.cn-beijing.volcengineapi.com") {
		t.Fatalf("unexpected url %s", got.URL)
	}

	for _, c := range labelCases {
		if c.valid {
			continue
		}
		if _, err := p.EndpointFor("ecs", c.value); err == nil {
			t.Fatalf("%s: region %q should be rejected", c.name, c.value)
		}
	}
}
