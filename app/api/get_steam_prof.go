package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"store/types"
	"store/util"
)

type SteamResp struct {
	Response struct {
		Players []types.SteamUser `json:"players"`
	} `json:"response"`
}

func GetSteamProfile(steamId string) types.SteamUser {
	resp, err := http.Get(fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", util.Config.SteamAPIKey, steamId))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response SteamResp
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	if len(response.Response.Players) == 0 {
		return types.SteamUser{}
	}
	return response.Response.Players[0]
}
