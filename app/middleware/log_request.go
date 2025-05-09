package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LoggingHandler(os.Stdout, next).ServeHTTP(w, r)
	})
}
