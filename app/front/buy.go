package front

import "net/http"

type BuyData struct {
	ID   int
	Code string
}

func Buy(w http.ResponseWriter, r *http.Request) {

}
