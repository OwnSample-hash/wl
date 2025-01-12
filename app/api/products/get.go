package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"store/types"
	"store/util"
)

func GetRawProds() (prods []types.Product) {
	rows, err := util.Db.Query("SELECT * FROM products")
	if errors.Is(err, sql.ErrNoRows) {
		log.Fatal("No rows")
		return
	} else if err != nil {
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	prods = make([]types.Product, 0)
	for rows.Next() {
		prod := types.Product{}
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt); err != nil {
			log.Println(err)
		}
		prods = append(prods, prod)
	}
	return prods
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	prods := GetRawProds()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prods); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}

}
