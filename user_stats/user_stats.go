package user_stats

import (
	"github.com/edmundluk3/steam-api/model"
	"github.com/edmundluk3/steam-api/api"
	"encoding/json"
)

func GlobalAchievementsPercentages(id model.SteamAppID) (*model.AchievementsList, error) {
	respByte, err := api.GetGlobalAchievementsPercentages(id)

	if err != nil {
		return nil, err
	}

	var resp model.AppAchievementsPercentages
	parseErr := json.Unmarshal(respByte, &resp)

	if parseErr != nil {
		return nil, parseErr
	}

	return &resp.Achievements.AchievementsList, nil
}

// todo
// func get_global_stats_for_game

func NumberOfCurrentPlayer(id model.SteamAppID) (*model.CurrentNumberOfPlayers, error) {
	respByte, err := api.GetNumberOfCurrentPlayers(id)

	if err != nil {
		return nil, err
	}

	var resp model.CurrentNumberOfPlayersResp
	parseErr := json.Unmarshal(respByte, &resp)

	if parseErr != nil {
		return nil, parseErr
	}
	return &resp.Response, nil
}

func SchemaForGame(id model.SteamAppID, key string, language string) (*model.GameSchema, error) {
	respByte, err := api.GetSchemaForGame(id, key, language)
	if err != nil {
		return nil, err
	}

	var resp model.GameSchemaResp

	parseErr := json.Unmarshal(respByte, &resp)

	if parseErr != nil {
		return nil, parseErr
	}

	return &resp.GameSchema, nil
}