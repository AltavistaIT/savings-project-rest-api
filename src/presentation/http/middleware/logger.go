package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s|UA:%s | IP: %s | Duration: %s", r.Proto, r.Method, r.RequestURI, r.UserAgent(), r.RemoteAddr, duration)
	})
}
