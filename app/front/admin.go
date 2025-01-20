package front

import (
	"html/template"
	"net/http"
	"store/app/api"
	"store/app/api/products"
	"store/types"
	"store/util"

	"github.com/gorilla/csrf"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(util.GetTeamplte("admin")...)
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
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	data := types.Payload{
		Title:     "Admin",
		SteamID:   steamID,
		Prods:     products.GetRawProds(r.URL.Path),
		Path:      r.URL.Path,
		User:      api.GetSteamProfile(steamID),
		IsAdmin:   true,
		CsrfField: csrf.TemplateField(r),
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		panic(err)
	}
}
