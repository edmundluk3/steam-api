package model

type (
	AppAchievementsPercentages struct {
		Achievements struct {
			AchievementsList AchievementsList`json:"achievements"`
		} `json:"achievementpercentages"`
	}

	AchievementsList []AppAchievementsPercentagesItem

	AppAchievementsPercentagesItem struct {
		Name       string  `json:"name"`
		Percentage float32 `json:"percentage"`
	}

	GameGlobalStats struct {
	}
)
