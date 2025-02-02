package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs the details of each HTTP request
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Wrap the ResponseWriter to capture the status code
		wrappedResponseWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(wrappedResponseWriter, r)

		// Log the request details
		log.Printf(
			"%s %s %d %s",
			r.Method,             // HTTP Method
			r.RequestURI,         // Request URI
			wrappedResponseWriter.statusCode, // Response status code
			time.Since(startTime), // Time taken to process the request
		)
	})
}

// responseWriter wraps the http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and writes the header
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}