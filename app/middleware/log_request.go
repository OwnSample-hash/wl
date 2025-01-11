package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LoggingHandler(os.Stdout, next).ServeHTTP(w, r)
	})
}
