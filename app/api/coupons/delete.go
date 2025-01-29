package coupons

import (
	"log"
	"net/http"
	"store/util"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	_, err := util.Db.Exec("DELETE FROM store_coupons WHERE id = ?", mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}
