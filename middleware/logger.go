package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		log.Println("Ami logger middleware ami age print hbo .....")

		next.ServeHTTP(w, r)

		log.Println("Ami logger middleware ami seshe print hbo .....")

		duration := time.Since(start)

		log.Println(r.Method, r.URL.Path, duration)
	})
}
