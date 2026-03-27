[← Retry](5-Retry.md) | Error Handling[(中文)](6-ErrorHandling-zh.md) | [Debugging →](7-Debugging.md)

---

# Error Handling

Errors are categorized as:

| Type | Description | Error type |
|---|---|---|
| 1. Client error | Validation failures before reaching server | `volcengineerr.Error` or native `error` |
| 2. Server error | Request reaches server; business error returned | `volcengineerr.RequestFailure` |
| 3. Network/timeout | DNS or timeout | `volcengineerr.Error` |
| 4. Others | Other errors | `volcengineerr.Error` or native `error` |

(See the original Chinese guide for the full code sample.)

---

[← Retry](5-Retry.md) | Error Handling[(中文)](6-ErrorHandling-zh.md) | [Debugging →](7-Debugging.md)
