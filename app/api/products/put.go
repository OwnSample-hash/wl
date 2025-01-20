package products

import (
	"log"
	"net/http"
	"store/util"

	"github.com/gorilla/mux"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
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
	_, err = util.Db.Exec("UPDATE store_products SET name= ?, description = ?, price_per_month=?, price=?, one_month=?, life_time=?, discount=? WHERE id = ?", name, description, pricePerMonth, price, oneMonth, lifeTime, discount, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
