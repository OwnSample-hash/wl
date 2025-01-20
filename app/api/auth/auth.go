package auth

import (
	"fmt"
	"net/http"
	"store/util"

	gosteamauth "github.com/TeddiO/GoSteamAuth/src"
)

func SteamAuth(w http.ResponseWriter, r *http.Request) {
	gosteamauth.RedirectClient(w, r, gosteamauth.BuildQueryString(
		fmt.Sprintf("http://%s:%d/api/auth/callback", util.Config.Http.Bind, util.Config.Http.Port)))
}
