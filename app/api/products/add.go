package products

import (
	"log"
	"net/http"
	"store/util"
)

func Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	pricePerMonth := r.FormValue("price_per_month")
	price := r.FormValue("price")
	discount := r.FormValue("discount")
	oneMonth := false
	lifeTime := false
	if r.FormValue("one_month") == "on" {
		oneMonth = true
	}
	if r.FormValue("life_time") == "on" {
		lifeTime = true
	}
	log.Printf("Inserting product: %s, %s, %s, %t, %t %s", name, description, price, oneMonth, lifeTime, discount)
	_, err = util.Db.Exec("INSERT INTO store_products (name, description, price_per_month, price, one_month, life_time, discount) VALUES (?, ?, ?, ?, ?, ?, ?)", name, description, pricePerMonth, price, oneMonth, lifeTime, discount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
