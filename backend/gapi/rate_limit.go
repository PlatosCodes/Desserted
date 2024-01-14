package gapi

import (
	"net/http"

	"golang.org/x/time/rate"
)

// RateLimitMiddleware creates a new HTTP middleware for rate limiting.
func RateLimitMiddleware(r rate.Limit, b int) func(http.Handler) http.Handler {
	limiter := rate.NewLimiter(r, b)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if !limiter.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
