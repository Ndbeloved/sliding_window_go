# Rate Limited API - Sliding Window Alogorithm

## Overview

This is a production-shaped, backend API project built in Go. It demonstrates a clean architecture for building REST APIs with rate limiting using the Sliding Window algorithm backed by Redis. The project is modular, testable, and designed with maintainability in mind.

Key features:

- CLI flag-based configuration.
- HTTP server lifecycle management with graceful shutdown.
- Rate limiting using Sliding Window algorithm.
- Middleware abstraction for rate limiting.
- Clean separation of concerns between server, router, middleware, handlers, and algorithm logic.
- Ready for extension: logging, auth, metrics, and more.

## Project Structure

```
cmd/
  server/main.go          # Entry point, wiring only

internal/config/
  config.go               # CLI flag parsing & validation

internal/server/
  server.go               # Server orchestration & graceful shutdown

internal/router/
  router.go               # Route definitions and middleware wiring

internal/handlers/
  health.go               # Health endpoint
  protected.go            # Protected endpoint

internal/middleware/
  ratelimit.go            # HTTP middleware wrapping rate limiter

internal/ratelimit/
  sliding_window.go       # Sliding window algorithm implementation

pkg/response/
  json.go                 # JSON response helpers
```

## Configuration

Configuration is handled via CLI flags. Example:

```bash
./server -port=8080 -limit=10 -window=60 -redis=localhost:6379
```

| Flag       | Description                          | Default          |
|------------|--------------------------------------|----------------|
| `-port`    | Server listening port                 | 8080           |
| `-limit`   | Allowed requests per window           | 10             |
| `-window`  | Rate limit window in seconds          | 60             |
| `-redis`   | Redis server address                  | localhost:6379 |

## Features

### Rate Limiting

- Implemented using Sliding Window algorithm.
- Redis sorted sets (ZSET) store timestamps for each client.
- Accurate, avoids bursts allowed by fixed window algorithms.
- Middleware applies limits transparently to handlers.

### Graceful Shutdown

- Listens to SIGINT/SIGTERM signals.
- Allows in-flight requests to finish before closing.

### Middleware & Handlers

- Middleware layer separates HTTP concerns from business logic.
- Handlers remain pure and only focus on responding.

### Extensible Architecture

- Easy to add logging, metrics, JWT auth, or additional middleware.
- Router is isolated; swapping `http.ServeMux` with a framework like `chi` or `gorilla/mux` is simple.
- Server orchestrates components without knowing their internals.

## Running the Project

1. Make sure Redis is running and accessible.
2. Build the project:

```bash
go build -o server ./cmd/server
```

3. Run with flags:

```bash
./server -port=8080 -limit=10 -window=60 -redis=localhost:6379
```

4. Test endpoints:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/protected
```

## Future Improvements

- JWT authentication and protected routes.
- Centralized logging middleware.
- Metrics for rate limiting (requests, blocked counts, etc.).
- Unit tests for rate limiter and middleware.
- Dockerization for deployment.
- Horizontal scaling considerations.

## License

MIT License

