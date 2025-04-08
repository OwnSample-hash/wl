package auth

import (
	"net/http"
	"store/util"
)

func FakeLogin(w http.ResponseWriter, r *http.Request) {
	session, err := util.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  

	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
