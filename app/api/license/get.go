package license

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"store/types"
	"store/util"
)

func GetRawLicenses() (licenses []types.License) {
	rows, err := util.Db.Query("SELECT * FROM store_licenses")
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

	licenses = make([]types.License, 0)
	for rows.Next() {
		license := types.License{}
		if err := rows.Scan(&license.Id, &license.Uuid, &license.UserId, &license.ValidTill, &license.GracePeriod, &license.CreatedAt, &license.UpdatedAt); err != nil {
			log.Println(err)
		}
		licenses = append(licenses, license)
	}
	return licenses
}

func GetLicense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(GetRawLicenses()); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
