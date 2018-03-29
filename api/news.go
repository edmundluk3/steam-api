package api

import "github.com/edmundluk3/steam-api/model"

func GetNewsList(id model.SteamAppID, count uint, maxLength uint) ([]byte, error) {
	payload := map[string]interface{}{
		"appid": id,
		"count": count,
		"maxlength": maxLength,
		"format": "json",
	}

	return ApiGet(
		STEAM_API_URL,
		NEWS_ENDPOINT,
		"/GetNewsForApp",
		"/v0002",
		payload,
	)
}