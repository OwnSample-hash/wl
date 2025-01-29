package payments

import (
	"net/http"
)

func (payment *Payment) GetPath() string {
	return payment.Path
}

func (payment *Payment) PaymentHandler(w *http.ResponseWriter, r *http.Request) {
	// Process payment
}
