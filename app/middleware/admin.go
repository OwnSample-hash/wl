package middleware

import (
	"net/http"
	"store/util"
	"strconv"
)

func CheckAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := util.Store.Get(r, "id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		steamID, ok := session.Values["steamID"].(string)
		if !ok {
			http.Error(w, "Unauthorized session not ok", http.StatusUnauthorized)
			return
		}
		i, err := strconv.ParseInt(steamID, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, id := range util.Config.AdminIds {
			if i == id {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Unauthorized after loop", http.StatusUnauthorized)
	})
}
