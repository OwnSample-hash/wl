package payments

import (
	"net/http"
	"store/types"
)

var Payments = []types.IPayment{}

func AddPayment(payment types.IPayment) {
	Payments = append(Payments, payment)
}

func PaymentGlobalHandler(w http.ResponseWriter, r *http.Request) {

}
