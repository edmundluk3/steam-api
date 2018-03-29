package api

import (
	_ "strconv"
	"fmt"
	"github.com/edmundluk3/steam-api/model"
)

func GetAppList() ([]byte, error){
	return ApiGet(
		STEAM_API_URL,
		APP_ENDPOINT,
		"/GetAppList",
		"/v2",
		nil,
	)
}

func GetAppDetail(appid model.SteamAppID, language string, cc string) ([]byte, error) {
	payload := map[string]interface{}{
		"appids": appid,
		"cc": cc,
		"l": language,
	}

	return ApiGet(
		"https://store.steampowered.com/api/appdetails",
		"",
		"",
		"",
		payload,
	)
}

func GetAppPrice(appids []model.SteamAppID, cc string) ([]byte, error){
	var ids string
	for _, v := range appids {
		ids += fmt.Sprintf("%d,", v)
	}

	payload := map[string]interface{}{
		"appids": ids,
		"cc": cc,
		"filters": "price_overview",
	}


	return ApiGet(
		"https://store.steampowered.com/api/appdetails",
		"",
		"",
		"",
		payload,
	)
}