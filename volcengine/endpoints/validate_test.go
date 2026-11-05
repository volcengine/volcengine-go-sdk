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
	{name: "non ascii", value: "cn-北京", valid: false},

	// start and end rules
	{name: "single letter", value: "a", valid: true},
	{name: "end with digit", value: "ap-southeast-1", valid: true},
	{name: "underscore inside", value: "rds_mysql", valid: true},
	{name: "upper case", value: "Redis", valid: true},
	{name: "start with digit", value: "1cn-beijing", valid: false},
	{name: "single digit", value: "1", valid: false},
	{name: "start with hyphen", value: "-cn-beijing", valid: false},
	{name: "start with underscore", value: "_ecs", valid: false},
	{name: "end with hyphen", value: "cn-beijing-", valid: false},
	{name: "end with underscore", value: "rds_mysql_", valid: false},
	{name: "single hyphen", value: "-", valid: false},
	{name: "single underscore", value: "_", valid: false},
}

func TestValidateErrorMessage(t *testing.T) {
	cases := map[string]string{
		"cn@beijing": "illegal character at index 2",
		"1cn":        "must start with a letter",
		"cn-":        "must not end with '-' or '_'",
	}
	for v, want := range cases {
		err := ValidateRegion(v)
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("region %q: want error containing %q, got %v", v, want, err)
		}
	}
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

	// The resolver already rejects unknown regions/services on its own
	// (regionMatcher, ServiceNotFound), so asserting err != nil would pass even
	// without the injection check. Require the exact error of the label check.
	for _, c := range labelCases {
		if c.valid {
			continue
		}
		_, err := r.EndpointFor("ecs", c.value)
		assertSameError(t, c.name+": region", err, ValidateRegion(c.value))

		_, err = r.EndpointFor(c.value, "cn-beijing")
		assertSameError(t, c.name+": service", err, ValidateService(c.value))
	}
}

// assertSameError checks that got is the error produced by the label check.
func assertSameError(t *testing.T, name string, got, want error) {
	t.Helper()
	if want == nil {
		t.Fatalf("%s: test case is expected to be invalid", name)
	}
	if got == nil {
		t.Fatalf("%s: expect error %q, got nil", name, want)
	}
	if got.Error() != want.Error() {
		t.Fatalf("%s: expect error %q, got %q", name, want, got)
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

	// With ResolveUnknownService, a service missing from p.Services is used
	// as-is to fill the "{service}" placeholder, this is the path where an
	// injected service would really reach the host name.
	got, err = p.EndpointFor("vpc", "cn-beijing", ResolveUnknownServiceOption)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasSuffix(got.URL, "vpc.cn-beijing.volcengineapi.com") {
		t.Fatalf("unexpected url %s", got.URL)
	}

	for _, c := range labelCases {
		if c.valid {
			continue
		}
		_, err := p.EndpointFor("ecs", c.value)
		assertSameError(t, c.name+": region", err, ValidateRegion(c.value))

		_, err = p.EndpointFor(c.value, "cn-beijing", ResolveUnknownServiceOption)
		assertSameError(t, c.name+": unknown service", err, ValidateService(c.value))
	}
}
