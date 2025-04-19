package coupons

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"store/types"
	"store/util"
)

func GetRawCopons() (coupons []types.Coupon) {
	rows, err := util.Db.Query("SELECT * FROM store_coupons")

	if errors.Is(err, sql.ErrNoRows) {
		log.Print("No rows")
		return []types.Coupon{}
	}
	if err != nil {
		panic(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	coupons = make([]types.Coupon, 0)
	for rows.Next() {
		coupon := types.Coupon{}
		if err := rows.Scan(&coupon.ID, &coupon.Code, &coupon.Discount, &coupon.ExpiresAt, &coupon.Status, &coupon.Remaining, &coupon.CreatedAt, &coupon.UpdatedAt); err != nil {
			panic(err)
		}
		coupon.FormattedCreatedAt = coupon.CreatedAt.Format("2006-01-02 15:04:05")
		coupon.FormattedUpdatedAt = coupon.UpdatedAt.Format("2006-01-02 15:04:05")
		coupon.FormattedExpiresAt = coupon.ExpiresAt.Format("2006-01-02 15:04:05")
		coupons = append(coupons, coupon)
	}

	return coupons
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(GetRawCopons()); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
