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
	price := r.FormValue("price")
	var oneMonth bool = false
	var lifeTime bool = false
	if r.FormValue("one_month") == "on" {
		oneMonth = true
	}
	if r.FormValue("life_time") == "on" {
		lifeTime = true
	}
	log.Printf("Inserting product: %s, %s, %s, %s, %s", name, description, price, oneMonth, lifeTime)
	_, err = util.Db.Exec("INSERT INTO products (name, description, price, one_month, life_time) VALUES (?, ?, ?, ?, ?)", name, description, price, oneMonth, lifeTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
