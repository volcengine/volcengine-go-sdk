package session

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/defaults"
)

// setEnvForTest sets an environment variable and returns a function restoring
// its previous value. t.Setenv / t.Cleanup are not used so that the test keeps
// building with the oldest Go version declared in go.mod.
func setEnvForTest(t *testing.T, key, value string) func() {
	old, had := os.LookupEnv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatalf("setenv %s: %v", key, err)
	}
	return func() { restoreEnv(key, old, had) }
}

// unsetEnvForTest unsets an environment variable and returns a function
// restoring its previous value.
func unsetEnvForTest(key string) func() {
	old, had := os.LookupEnv(key)
	os.Unsetenv(key)
	return func() { restoreEnv(key, old, had) }
}

func restoreEnv(key, old string, had bool) {
	if had {
		os.Setenv(key, old)
	} else {
		os.Unsetenv(key)
	}
}

func newSessionWithRegion(region string) *Session {
	return &Session{
		Config:   defaults.Config().WithRegion(region).WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")),
		Handlers: defaults.Handlers(),
	}
}

// The region read from the environment is text input: surrounding white
// spaces are trimmed before the value is used.
func TestEnvConfigRegionIsTrimmed(t *testing.T) {
	defer unsetEnvForTest("VOLCSTACK_DEFAULT_REGION")()
	cases := map[string]string{
		"cn-beijing":          "cn-beijing",
		" cn-beijing ":        "cn-beijing",
		"\tcn-beijing\n":      "cn-beijing",
		"\r\n ap-southeast-1": "ap-southeast-1",
	}
	for raw, want := range cases {
		restore := setEnvForTest(t, "VOLCSTACK_REGION", raw)
		got := envConfigLoad(false).Region
		restore()
		if got != want {
			t.Errorf("VOLCSTACK_REGION=%q: got region %q, want %q", raw, got, want)
		}
	}
}

// A white space only region is treated as not set, so the next variable is
// used (or the region stays empty), as if the variable was not defined.
func TestEnvConfigBlankRegionIsNotSet(t *testing.T) {
	defer setEnvForTest(t, "VOLCSTACK_REGION", "  \t ")()
	defer unsetEnvForTest("VOLCSTACK_DEFAULT_REGION")()
	if got := envConfigLoad(false).Region; got != "" {
		t.Errorf("blank region: got %q, want empty", got)
	}

	defer setEnvForTest(t, "VOLCSTACK_DEFAULT_REGION", " ap-southeast-1 ")()
	if got := envConfigLoad(true).Region; got != "ap-southeast-1" {
		t.Errorf("blank region with default region: got %q, want %q", got, "ap-southeast-1")
	}
	// The default region is only read when the shared config is enabled.
	if got := envConfigLoad(false).Region; got != "" {
		t.Errorf("default region without shared config: got %q, want empty", got)
	}
}

// Trimming does not hide invalid characters: what remains after trimming is
// still validated before it is used to build the endpoint host.
func TestEnvConfigTrimmedRegionIsStillValidated(t *testing.T) {
	defer setEnvForTest(t, "VOLCSTACK_REGION", " cn-beijing@127.0.0.1 ")()
	defer unsetEnvForTest("VOLCSTACK_DEFAULT_REGION")()
	region := envConfigLoad(false).Region
	if region != "cn-beijing@127.0.0.1" {
		t.Fatalf("got region %q", region)
	}
	if _, err := newSessionWithRegion(region).clientConfigWithErr("ecs"); err == nil {
		t.Fatal("an invalid region read from the environment must be rejected")
	}
}

// Regions set in code are not trimmed: they are validated as is.
func TestCodeRegionIsNotTrimmed(t *testing.T) {
	if _, err := newSessionWithRegion("cn-beijing").clientConfigWithErr("ecs"); err != nil {
		t.Fatalf("valid region must pass: %v", err)
	}
	for _, region := range []string{" cn-beijing", "cn-beijing ", "\tcn-beijing", "cn-beijing\n", "   "} {
		if _, err := newSessionWithRegion(region).clientConfigWithErr("ecs"); err == nil {
			t.Errorf("region %q set in code must be rejected, not trimmed", region)
		}
	}
}

// defaults.Config reads the region from the environment and trims it as well.
func TestDefaultsConfigRegionIsTrimmed(t *testing.T) {
	restore := setEnvForTest(t, "VOLCSTACK_REGION", " cn-beijing\n")
	got := *defaults.Config().Region
	restore()
	if got != "cn-beijing" {
		t.Errorf("defaults.Config region: got %q, want %q", got, "cn-beijing")
	}

	defer setEnvForTest(t, "VOLCSTACK_REGION", "   ")()
	if got := *defaults.Config().Region; got != "" {
		t.Errorf("defaults.Config blank region: got %q, want empty", got)
	}
}

// writeSharedConfigForTest writes content to a temporary shared config file and
// returns its path with a function removing it.
func writeSharedConfigForTest(t *testing.T, content string) (string, func()) {
	dir, err := ioutil.TempDir("", "region-shared-config")
	if err != nil {
		t.Fatalf("temp dir: %v", err)
	}
	path := filepath.Join(dir, "config")
	if err := ioutil.WriteFile(path, []byte(content), 0600); err != nil {
		os.RemoveAll(dir)
		t.Fatalf("write shared config: %v", err)
	}
	return path, func() { os.RemoveAll(dir) }
}

// The shared config file is text input as well: surrounding white spaces of the
// region are trimmed, quoted or not, tabs included.
func TestSharedConfigRegionIsTrimmed(t *testing.T) {
	cases := map[string]string{
		"region = cn-beijing\n":         "cn-beijing",
		"region =   cn-beijing   \n":    "cn-beijing",
		"region = cn-beijing\t\n":       "cn-beijing",
		"region = \"  cn-beijing  \"\n": "cn-beijing",
		"region = \"\tcn-beijing\"\n":   "cn-beijing",
		"region = \"   \"\n":            "",
	}
	for line, want := range cases {
		path, cleanup := writeSharedConfigForTest(t, "[default]\n"+line)
		cfg, err := loadSharedConfig("default", []string{path}, true)
		cleanup()
		if err != nil {
			t.Fatalf("%q: load shared config: %v", line, err)
		}
		if cfg.Region != want {
			t.Errorf("%q: got region %q, want %q", line, cfg.Region, want)
		}
	}
}

// End to end: a white space only region environment variable does not hide the
// region of the shared config file, and the resulting client config is valid.
func TestBlankEnvRegionFallsBackToSharedConfig(t *testing.T) {
	path, cleanup := writeSharedConfigForTest(t, "[default]\nregion = \" cn-beijing \"\n")
	defer cleanup()
	defer setEnvForTest(t, "VOLCSTACK_REGION", "  \t ")()
	defer unsetEnvForTest("VOLCSTACK_DEFAULT_REGION")()
	defer unsetEnvForTest("VOLCSTACK_PROFILE")()
	defer unsetEnvForTest("VOLCSTACK_DEFAULT_PROFILE")()

	sess, err := NewSessionWithOptions(Options{
		SharedConfigState: SharedConfigEnable,
		SharedConfigFiles: []string{path},
		Config:            *defaults.Config().WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")),
	})
	if err != nil {
		t.Fatalf("new session: %v", err)
	}
	if got := *sess.Config.Region; got != "cn-beijing" {
		t.Fatalf("got region %q, want %q", got, "cn-beijing")
	}
	if _, err := sess.clientConfigWithErr("ecs"); err != nil {
		t.Fatalf("client config: %v", err)
	}
}
