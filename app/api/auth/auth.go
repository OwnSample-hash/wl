package auth

import (
	"fmt"
	gosteamauth "github.com/TeddiO/GoSteamAuth/src"
	"net/http"
	"store/util"
)

func SteamAuth(w http.ResponseWriter, r *http.Request) {
	gosteamauth.RedirectClient(w, r, gosteamauth.BuildQueryString(
		fmt.Sprintf("http://%s:%d/api/auth/callback", util.Config.Http.Bind, util.Config.Http.Port)))
	return
}
