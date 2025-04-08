package products

import (
	"log"
	"net/http"
	"store/util"

	"github.com/gorilla/mux"
)

func FlipVisibility(w http.ResponseWriter, r *http.Request) {
	_, err := util.Db.Exec("UPDATE store_products SET is_active = NOT is_active WHERE id=?", mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}
