package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/limiter"
	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/metrics"
)

func getClientIP(r *http.Request) string {

	// If behind a reverse proxy/load balancer
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Take the first IP if multiple exist
		return strings.TrimSpace(strings.Split(forwarded, ",")[0])
	}

	// Some proxies use X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fallback for local development
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func RateLimiter(l limiter.Limiter, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		metrics.Requests.Inc()

		clientIP := getClientIP(r)

		allowed, err := l.Allow(clientIP)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !allowed {

			metrics.Throttled.Inc()

			w.Header().Set("Retry-After", "1")
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
