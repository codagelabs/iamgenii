package middleware

import (
	"net/http"

	log "github.com/iamgenii/logs"
)

// LoggingMiddleware log request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Do stuff here
		//	log.Println(r.RequestURI)

		req = req.WithContext(log.WithRqID(req.Context()))
		w.Header().Add("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req)
	})
}
