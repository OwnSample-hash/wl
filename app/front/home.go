package front

import (
	"html/template"
	"log"
	"net/http"
	"store/util"
)

type Payload struct {
	Title   string
	SteamID string
}

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

	var data Payload

	if !ok {
		log.Println("steamID not found")
		data = Payload{
			Title:   "Home",
			SteamID: "",
		}
	} else {
		log.Println("steamID found", steamID)
		data = Payload{
			Title:   "Home",
			SteamID: steamID,
		}
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
