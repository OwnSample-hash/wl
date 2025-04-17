package products

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"store/types"
	"store/util"
)

func GetRawProds(path string, CsrkToken string) (prods []types.Product) {
	rows, err := util.Db.Query("SELECT * FROM store_products")
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
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.PricePerMonth, &prod.Price, &prod.OneMonth, &prod.LifeTime, &prod.Discount, &prod.IsActive, &prod.CreatedAt, &prod.UpdatedAt); err != nil {
			log.Println(err)
		}
		prod.Path = path
		prod.CsrfToken = CsrkToken
		prod.FormattedCreatedAt = prod.CreatedAt.Format("2006-01-02 15:04:05")
		prod.FormattedUpdatedAt = prod.UpdatedAt.Format("2006-01-02 15:04:05")
		if prod.Discount > 0 {
			prod.DiscountPrice = prod.Price - (prod.Price * (prod.Discount / 100))
			prod.DiscountedMonthlyPrice = prod.PricePerMonth - (prod.PricePerMonth * (prod.Discount / 100))
		} else {
			prod.DiscountPrice = prod.Price
			prod.DiscountedMonthlyPrice = prod.PricePerMonth
		}
		prods = append(prods, prod)
	}
	return prods
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(GetRawProds(r.URL.Path, "")); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
