# Go Market Data Router

A Go service seed for routing price ticks to subscribed clients using explicit subscription maps and deterministic recipient selection.

## Stack

Go, market data routing, subscriber fan-out

## Problem

Market data systems must route high-volume updates to the right subscribers without unnecessary fan-out or nondeterministic behavior.

## Architecture

- internal/book/orderbook.go contains the tick routing domain.
- cmd/router/main.go provides a small executable demonstration.
- The module is ready for benchmarks and gRPC streaming.

## Implemented Production Readiness

- CI compiles and tests all Go packages.
- Routing output is sorted for deterministic behavior.
- The code is split into command and internal packages.

## Run And Test

```powershell
go test ./...
go run ./cmd/router
```

## Quality Gates

- Project-specific GitHub Actions workflow included under .github/workflows/ci.yml.
- Generated build outputs and dependency folders are excluded through .gitignore.
- Tests and validation commands are intentionally small enough to run during code review.

## Production Extension Points

- Add benchmark tests for fan-out latency.
- Add Prometheus metrics.
- Add gRPC streaming transport.

## Repository Hygiene

This repository contains original portfolio code only. It does not include employer source code, private resumes, generated binaries, local credentials, or large media files.

