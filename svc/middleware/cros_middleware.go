package middleware

import "net/http"

//CROSMiddleware resolves cross origin policies
func CROSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, content-type, Content-Type, Authorization, Content-Length, Accept-Encoding, accept, origin, Cache-Control, Accept, X-CSRF-DecodeJwtToken")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS, DELETE, PATCH")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, req)
	})
}
