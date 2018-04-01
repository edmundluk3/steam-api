package api

import "github.com/edmundluk3/steam-api/model"

func GetPlayerSummaries(key model.SteamApiKey, ids []model.SteamUserId) ([]byte, error) {
	payload := map[string]interface{}{
		"key":      key,
		"steamids": ids,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_ENDPOINT,
		"/GetPlayerSummaries",
		"/v2",
		payload,
	)
}

func GetFriendList(key model.SteamApiKey, id model.SteamUserId, relationship string) ([]byte, error) {
	payload := map[string]interface{}{
		"key":          key,
		"steamid":      id,
		"relationship": relationship,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_ENDPOINT,
		"/GetFriendList",
		"/v1",
		payload,
	)
}

func GetPlayerBans(key model.SteamApiKey, ids []model.SteamUserId) ([]byte, error) {
	payload := map[string]interface{}{
		"key":      key,
		"steamids": ids,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_ENDPOINT,
		"/GetPlayerBans",
		"/v1",
		payload,
	)
}

func GetUserGroupList(key model.SteamApiKey, id model.SteamUserId) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_ENDPOINT,
		"/GetUserGroupList",
		"/v1",
		payload,
	)
}

func GetPlayerAchievements(key model.SteamApiKey, id model.SteamUserId,
	appId model.SteamAppID, language string) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
		"appid":   appId,
		"l":       language,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_STATS_ENDPOINT,
		"/GetPlayerAchievements",
		"/v0001",
		payload,
	)
}

func GetUserStatsForGame(key model.SteamApiKey, id model.SteamUserId,
	appId model.SteamAppID, language string) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
		"appid":   appId,
		"l":       language,
	}

	return ApiGet(
		STEAM_API_URL,
		USER_STATS_ENDPOINT,
		"/GetUserStatsForGame",
		"/v0002",
		payload,
	)
}

func GetRecentlyPlayedGame(key model.SteamApiKey, id model.SteamUserId,
	count uint) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
		"count":   count,
	}

	return ApiGet(
		STEAM_API_URL,
		PlAYER_ENDPOINT,
		"/GetRecentlyPlayedGames",
		"/v0001",
		payload,
	)
}

func GetOwnedGame(key model.SteamApiKey, id model.SteamUserId, includeAppInfo bool,
	includeFreeGame bool, ids []model.SteamAppID) ([]byte, error) {
	payload := map[string]interface{}{
		"key":                       key,
		"steamid":                   id,
		"include_appinfo":           includeAppInfo,
		"include_played_free_games": includeFreeGame,
		"appids_filter":             ids,
	}

	return ApiGet(
		STEAM_API_URL,
		PlAYER_ENDPOINT,
		"/GetOwnedGames",
		"/v0001",
		payload,
	)
}

func GetSteamLevel(key model.SteamApiKey, id model.SteamUserId) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
	}

	return ApiGet(
		STEAM_API_URL,
		PlAYER_ENDPOINT,
		"/GetSteamLevel",
		"/v1",
		payload,
	)
}

func GetBadges(key model.SteamApiKey, id model.SteamUserId) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
	}

	return ApiGet(
		STEAM_API_URL,
		PlAYER_ENDPOINT,
		"/GetBadges",
		"/v1",
		payload,
	)
}

func GetCommunityBadgesProgress(key model.SteamApiKey, id model.SteamUserId, badgeId uint64) ([]byte, error) {
	payload := map[string]interface{}{
		"key":     key,
		"steamid": id,
		"badgeid": badgeId,
	}

	return ApiGet(
		STEAM_API_URL,
		PlAYER_ENDPOINT,
		"/GetCommunityBadgeProgress",
		"/v1",
		payload,
	)
}
