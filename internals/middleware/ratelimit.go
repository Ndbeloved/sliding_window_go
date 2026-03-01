package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type RateLimiter interface {
	Allow(ctx context.Context, key string) (bool, int, error)
}

func RateLimitMiddleware(limiter RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			ip := clientIp(r)
			if ip == "" {
				http.Error(w, "unable to determine IP", http.StatusBadRequest)
				return
			}

			allowed, remaining, err := limiter.Allow(ctx, ip)
			if err != nil {
				log.Printf("rate limit error: %v\n", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				w.Header().Set("Retry-After", "60")
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
			next.ServeHTTP(w, r)
		})
	}
}

func clientIp(r *http.Request) string {
	if f := r.Header.Get("X-Forwarded-For"); f != "" {
		ips := strings.Split(f, ",")
		return strings.TrimSpace(ips[0])
	}

	ipPort := strings.Split(r.RemoteAddr, ":")
	if len(ipPort) > 0 {
		return ipPort[0]
	}
	return ""
}
