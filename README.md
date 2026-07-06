# Distributed Rate Limiter / API Gateway

A distributed API Gateway written in Go that performs request rate limiting using Redis as shared state.

The project supports two pluggable rate limiting algorithms:

- Token Bucket
- Sliding Window Log

The gateway sits in front of backend services and decides whether requests should be forwarded or rejected with HTTP **429 Too Many Requests**.

---

# Features

- Reverse Proxy API Gateway
- Redis-backed distributed rate limiting
- Token Bucket algorithm
- Sliding Window Log algorithm
- Configurable algorithm selection
- HTTP 429 responses with `Retry-After`
- Prometheus metrics
- Grafana dashboard
- Docker Compose deployment

---

# Architecture

```text
                 Client
                    │
                    ▼
          +------------------+
          |  API Gateway     |
          |      (Go)        |
          +------------------+
             │           │
             │           ▼
             │     Prometheus
             │           │
             ▼           ▼
           Redis      Grafana
             │
             ▼
      Backend Service
```

---

# Why Redis?

The gateway can run multiple instances.

If every instance stored rate limits in memory:

```
Gateway A
5 requests

Gateway B
5 requests
```

A client could bypass the limit simply by hitting another gateway.

Redis provides shared state so every gateway instance sees the same counters.

---

# Why Lua?

Redis commands like:

- INCR
- EXPIRE
- ZADD
- ZREMRANGEBYSCORE

must execute atomically.

Lua scripts guarantee that every rate-limit decision happens as one indivisible operation, preventing race conditions under concurrent requests.

---

# Algorithms

## Token Bucket

- Allows bursts
- Tokens refill over time
- Efficient
- Used by many API gateways

---

## Sliding Window

- Stores request timestamps
- Removes expired timestamps
- Counts requests within the current window
- Provides stricter rate limiting

---

# Metrics

Prometheus metrics exposed:

- gateway_requests_total
- gateway_throttled_total

Visualized using Grafana.

---

# Running

```bash
docker compose up --build
```

Gateway:

```
http://localhost:8080
```

Backend:

```
http://localhost:8081
```

Prometheus:

```
http://localhost:9090
```

Grafana:

```
http://localhost:3000
```

---

# Future Improvements

- Dynamic routing table
- Per-user API keys
- Configurable rate limits by subscription tier
- Distributed tracing
- Kubernetes deployment