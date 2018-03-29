package api

import "github.com/edmundluk3/steam-api/model"

func GetGlobalAchievementsPercentages(id model.SteamAppID) ([]byte, error) {
	payload := map[string]interface{}{
		"gameid": id,
		"format": "json",
	}

	return ApiGet(
		STEAM_API_URL,
		USER_STATS_ENDPOINT,
		"/GetGlobalAchievementPercentagesForApp",
		"/v2",
		payload,
	)
}

func GetGlobalStatsForGame(id model.SteamAppID, count uint, name string) ([]byte, error) {
	payload := map[string]interface{}{
		"gameid": id,
		"count": count,
		"name": name,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_STATS_ENDPOINT,
		"/GetGlobalStatsForGame",
		"/v1",
		payload,
	)
}

func GetNumberOfCurrentPlayers(id model.SteamAppID) ([]byte, error) {
	return ApiGet(
		STEAM_API_URL,
		USER_STATS_ENDPOINT,
		"GetNumberOfCurrentPlayers",
		"/v1",
		map[string]interface{}{
			"appid": id,
		},
	)
}

func GetSchemaForGame(id model.SteamAppID, key string, language string) ([]byte, error) {
	payload := map[string]interface{}{
		"appid": id,
		"key": key,
		"l": language,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_STATS_ENDPOINT,
		"/GetSchemaForGame",
		"/v2",
		payload,
	)
}