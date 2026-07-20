package middleware

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		// 1 petición por segundo, burst de 5
		limiter = rate.NewLimiter(1, 5)
		visitors[ip] = limiter
	}
	return limiter
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := getVisitor(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, "Demasiadas peticiones", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
