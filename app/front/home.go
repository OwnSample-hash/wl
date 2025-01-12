package front

import (
	"html/template"
	"net/http"
	"store/app/api/products"
	"store/types"
	"store/util"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.gohtml")
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
	var data types.Payload

	if !ok {
		data = types.Payload{
			Title:   "Home",
			SteamID: "",
			Prods:   products.GetRawProds(),
		}
	} else {
		data = types.Payload{
			Title:   "Home",
			SteamID: steamID,
			Prods:   products.GetRawProds(),
		}
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
