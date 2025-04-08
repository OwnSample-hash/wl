package front

import (
	"html/template"
	"net/http"
	"store/types"
	"store/util"
	"strconv"

	"github.com/gorilla/csrf"
)

func LicenseHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(util.GetTeamplte("license")...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := util.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	steamID, ok := session.Values["steamID"].(string)
	isAdmin := false
	if ok {
		i, err := strconv.ParseInt(steamID, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, id := range util.Config.AdminIds {
			if i == id {
				isAdmin = true
				break
			}
		}
	}

	data := types.Payload{
		Title:     "License",
		SteamID:   steamID,
		IsAdmin:   isAdmin,
		CsrfToken: csrf.Token(r),
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
