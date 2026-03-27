[← Error Handling](6-ErrorHandling.md) | Debugging[(中文)](7-Debugging-zh.md) | [SDK Integration →](../SDK_Integration.md)

---

# Debugging

## Enable Debug Mode

Debug logs are disabled by default. Enable via `WithDebug(true)`.

```go
config := volcengine.NewConfig().
    WithRegion(region).
    WithDebug(true).
    WithCredentials(credentials.NewEnvCredentials())
```

# Log Output

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

[← Error Handling](6-ErrorHandling.md) | Debugging[(中文)](7-Debugging-zh.md) | [SDK Integration →](../SDK_Integration.md)
