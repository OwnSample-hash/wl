package front

import (
	"encoding/json"
	"html/template"
	"net/http"
	"slices"
	"store/app/api"
	"store/app/api/coupons"
	"store/app/api/products"
	"store/types"
	"store/util"
	"strconv"

	"github.com/gorilla/csrf"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(util.GetTeamplte("index")...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	session, err := util.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	steamID, ok := session.Values["steamID"].(string)
	var data types.Payload
	isAdmin := false
	if ok {
		i, err := strconv.ParseInt(steamID, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		if slices.Contains(util.Config.AdminIds, i) {
			isAdmin = true
		}
	}
	var User types.SteamUser
	user, ok := session.Values["user"].(string)
	if !ok {
		User = api.GetSteamProfile(steamID)
	} else {
		err = json.Unmarshal([]byte(user), &User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}

	if !ok {
		data = types.Payload{
			Title:   "Home",
			SteamID: "",
			IsAdmin: isAdmin,
			Path:    r.URL.Path,
			User:    User,
			Prods:   products.GetRawProds(r.URL.Path, csrf.Token(r)),
			Coupons: coupons.GetRawProds(),
		}
	} else {
		data = types.Payload{
			Title:   "Home",
			SteamID: steamID,
			IsAdmin: isAdmin,
			Path:    r.URL.Path,
			User:    User,
			Prods:   products.GetRawProds(r.URL.Path, csrf.Token(r)),
			Coupons: coupons.GetRawProds(),
		}
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
