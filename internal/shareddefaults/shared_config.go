package shareddefaults

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"os"
	"path/filepath"
	"runtime"
)

// SharedCredentialsFilename returns the SDK's default file path
// for the shared credentials file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.volcstack/credentials
//   - Windows: %USERPROFILE%\.volcstack\credentials
func SharedCredentialsFilename() string {
	return filepath.Join(UserHomeDir(), ".volcstack", "credentials")
}

// SharedConfigFilename returns the SDK's default file path for
// the shared config file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.volcstack/config
//   - Windows: %USERPROFILE%\.volcstack\config
func SharedConfigFilename() string {
	return filepath.Join(UserHomeDir(), ".volcstack", "config")
}

// UserHomeDir returns the home directory for the user the process is
// running under.
func UserHomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}

	// *nix
	return os.Getenv("HOME")
}
