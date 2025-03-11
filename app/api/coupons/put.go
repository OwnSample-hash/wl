package coupons

import (
	"net/http"
	"store/util"

	"github.com/gorilla/mux"
)

func Put(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	code := r.FormValue("code")
	uses := r.FormValue("uses")
	discount := r.FormValue("discount")
	expiresAt := r.FormValue("expires_at")
	_, err = util.Db.Exec("UPDATE store_coupons SET code = ?, discount = ?, exp = ?, remaining = ? WHERE id = ?", code, discount, expiresAt, uses, mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
