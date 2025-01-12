package front

import (
	"html/template"
	"net/http"
	"store/app/api/products"
	"store/types"
	"store/util"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/admin.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := util.Store.Get(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	steamID, ok := session.Values["steamID"].(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	data := types.Payload{
		Title:   "Admin",
		SteamID: steamID,
		Prods:   products.GetRawProds(),
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
