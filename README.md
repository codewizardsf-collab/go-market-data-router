# Go Market Data Router

Go project for routing high-volume market ticks to subscribers with bounded queues, symbol subscriptions, and latency-aware fan-out.

## Resume Fit

- Go market data distribution.
- Goroutines/channels/concurrency patterns.
- Low-latency financial infrastructure.

## Toolchain

Go is not installed in this workspace, so this project is implementation-ready but not locally verified here.

## Production Next Steps

- Add Prometheus metrics.
- Add gRPC streaming transport.
- Add benchmark tests for fan-out latency.
