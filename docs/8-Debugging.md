[← Error Handling](7-ErrorHandling.md) | Debugging[(中文)](8-Debugging-zh.md) | [Overview →](0-Overview.md)

---

## Debugging

### Enable Debug Mode

Debug logs are disabled by default. Enable via `WithDebug(true)`.

```go
config := volcengine.NewConfig().
	WithRegion(region).
	WithDebug(true).
	WithCredentials(credentials.NewEnvCredentials())
```

### Print error details

After a call fails, call `volcengineerr.PrintError(os.Stderr, err)` directly to print captured status, headers, body and underlying causes without enabling Debug or a type assertion. `DetailedRequestFailure.ResponseDetails()` returns an HTTP snapshot copy, or nil when none was captured. The package-level `volcengineerr.GetErrorDetails(err)` returns complete error details. See [Full error details](7-ErrorHandling.md#full-error-details) for capture scope, legacy fallback and body completeness. Output occurs only when explicitly requested; the return value reports output failures.

### Log Output

By default, logs are written to `os.Stdout`. Use `WithLogWriter` to write to a file or other writer.

```go
file, _ := os.Create("sdk.log")
config := volcengine.NewConfig().
	WithRegion(region).
	WithDebug(true).
	WithLogWriter(file).
	WithCredentials(credentials.NewEnvCredentials())
```

---

[← Error Handling](7-ErrorHandling.md) | Debugging[(中文)](8-Debugging-zh.md) | [Overview →](0-Overview.md)
