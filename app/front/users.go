package front

import (
	"html/template"
	"net/http"
	"slices"
	"store/app/api/users"
	"store/types"
	"store/util"
	"strconv"

	"github.com/gorilla/csrf"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(util.GetTeamplte("users")...)
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
		if slices.Contains(util.Config.AdminIds, i) {
			isAdmin = true
		}
	}

	data := types.Payload{
		Title:     "Users",
		SteamID:   steamID,
		IsAdmin:   isAdmin,
		Users:     users.GetRawUsers(),
		CsrfToken: csrf.Token(r),
		CsrfField: csrf.TemplateField(r),
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
