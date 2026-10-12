# tinylb

`tinylb` is a lightweight Layer 7 (L7) HTTP load balancer built from scratch in Go. Designed as a foundational systems engineering project, it distributes incoming client traffic across multiple backend servers using a round-robin algorithm while continuously monitoring backend health.

## Architecture

```
         ┌─ backend1 :9001
client → tinylb :8080 ─┼─ backend2 :9002
         │            └─ backend3 :9003
         ├─ health checker (goroutine, every 5s)
         └─ /metrics → Prometheus
```

## Core Mechanics

- **L7 Reverse Proxy:** Intercepts incoming HTTP requests on port `8080` and forwards them to upstream backend servers using Go's `net/http/httputil` reverse proxy package.
- **Round-Robin Routing:** Uses thread-safe atomic counters (`sync/atomic`) to cycle through an active slice of backend targets evenly.
- **Background Health Checking:** A dedicated background `goroutine` paired with a `time.Ticker` pings backend health endpoints every 5 seconds. Automatically marks dead nodes down and reroutes traffic around them.
- **Observability:** Exposes a `/metrics` endpoint formatted for Prometheus via `prometheus/client_golang` to track request counts, latencies, and backend health states.
- **Graceful Shutdown:** Traps OS interruption signals (`SIGINT`, `SIGTERM`) using `signal.NotifyContext` to safely drain active connections before exiting.
- **Containerization & Deployment:** Packaged into minimal distroless Docker images via multi-stage builds, orchestrated with Docker Compose, and deployable to cloud infrastructure (EC2).

