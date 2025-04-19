package front

import (
	"html/template"
	"net/http"
	"slices"
	"store/app/api/license"
	"store/app/api/products"
	"store/types"
	"store/util"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type BuyData struct {
	ID   int
	Code string
}

func Buy(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(util.GetTeamplte("buy")...)
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

	prods := products.GetRawProds(r.URL.Path, csrf.Token(r))
	prod_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, prod := range prods {
		if prod.ID == prod_id {
			prods = append([]types.Product{}, prods[i])
		}
	}

	data := types.Payload{
		Title:     "License",
		SteamID:   steamID,
		IsAdmin:   isAdmin,
		Prods:     prods,
		License:   license.GetRawLicenses(),
		CsrfToken: csrf.Token(r),
		CsrfField: csrf.TemplateField(r),
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
