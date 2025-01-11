package auth

import (
	"fmt"
	"net/http"
	"store/util"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	session, err := util.Store.Get(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1
	if err = session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusFound)
}
