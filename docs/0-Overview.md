Overview[(中文)](0-Overview-zh.md) | [Credentials →](1-Credentials.md)

---

## SDK Integration

When calling APIs, it is recommended to integrate the SDK in your project. Using the SDK simplifies development, speeds up integration, and reduces long-term maintenance costs. Volcengine SDK integration typically includes three steps: importing the SDK, configuring access credentials, and writing API call code.

## Requirements

1. Go version **>= 1.14**.
2. If you use Ark service (`service/arkruntime`), Go version **>= 1.18** is required.
3. It is recommended to use `go mod` for dependency management.

## Table of Contents

1. [Credentials](1-Credentials.md) — AK/SK, STS, AssumeRole, OIDC, SAML, ECS Role, Default Chain
2. [Endpoint Configuration](2-Endpoint.md) — Custom Endpoint, RegionId, Automatic Resolution
3. [Transport](3-Transport.md) — HTTP Connection Pool, HTTPS Scheme, SSL Verification, TLS Version
4. [Proxy](4-Proxy.md) — HTTP(S) Proxy Configuration
5. [Timeout](5-Timeout.md) — Global Timeouts, Per-API Timeout
6. [Retry](6-Retry.md) — Retry Strategy
7. [Error Handling](7-ErrorHandling.md) — Exception Handling
8. [Debugging](8-Debugging.md) — Debug Mode, Log Output

---

English | [中文](0-Overview-zh.md)
