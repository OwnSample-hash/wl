package api

import (
	"fmt"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "GET /api/products")
	if err != nil {
		return
	}
	_, err = w.Write([]byte("Hello, World!"))
	if err != nil {
		return
	}
}
