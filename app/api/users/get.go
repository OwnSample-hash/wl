package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"store/types"
	"store/util"
)

func GetRawUsers() (users []types.User) {
	rows, err := util.Db.Query("SELECT * FROM store_users")
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

	users = make([]types.User, 0)
	for rows.Next() {
		user := types.User{}
		if err := rows.Scan(&user.Id, &user.SecondaryId, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Println(err)
		}
		user.FormattedCreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
		user.FormattedUpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")
		users = append(users, user)
	}
	return users
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(GetRawUsers()); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
