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