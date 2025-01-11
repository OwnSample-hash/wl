package auth

import (
	"fmt"
	gosteamauth "github.com/TeddiO/GoSteamAuth/src"
	"log"
	"net/http"
	"net/url"
	"store/util"
)

func SteamCallback(w http.ResponseWriter, r *http.Request) {
	queryString, _ := url.ParseQuery(r.URL.RawQuery)
	queryMap := gosteamauth.ValuesToMap(queryString)
	steamID, isValid, err := gosteamauth.ValidateResponse(queryMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := util.Store.Get(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["steamID"] = steamID

	if err = session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("%s, %t", steamID, isValid)
	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusFound)
}
