package coupons

import (
	"net/http"
	"store/util"
)

func Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	code := r.FormValue("code")
	uses := r.FormValue("uses")
	discount := r.FormValue("discount")
	expiresAt := r.FormValue("expires_at")
	_, err = util.Db.Exec("INSERT INTO store_coupons (code, discount, exp, remaining) VALUES (?, ?, ?, ?)", code, discount, expiresAt, uses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
