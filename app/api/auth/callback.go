package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"store/app/api"
	"store/util"

	gosteamauth "github.com/TeddiO/GoSteamAuth/src"
)

func SteamCallback(w http.ResponseWriter, r *http.Request) {
	queryString, _ := url.ParseQuery(r.URL.RawQuery)
	queryMap := gosteamauth.ValuesToMap(queryString)
	steamID, _, err := gosteamauth.ValidateResponse(queryMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := util.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["steamID"] = steamID

	user := api.GetSteamProfile(steamID)
	var buffer bytes.Buffer
	err = json.NewEncoder(&buffer).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["user"] = buffer.String()

	if err = session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
