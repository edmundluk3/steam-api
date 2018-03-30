package model

type (
	AppAchievementsPercentages struct {
		Achievements struct {
			AchievementsList AchievementsList `json:"achievements"`
		} `json:"achievementpercentages"`
	}

	AchievementsList []AppAchievementsPercentagesItem

	AppAchievementsPercentagesItem struct {
		Name       string  `json:"name"`
		Percentage float32 `json:"percentage"`
	}

	GameGlobalStats struct {
	}

	CurrentNumberOfPlayersResp struct {
		Response CurrentNumberOfPlayers `json:"response"`
	}

	CurrentNumberOfPlayers struct {
		PlayerCount uint `json:"player_count"`
		Result      int  `json:"result"`
	}

	GameSchemaResp struct {
		GameSchema GameSchema `json:"game"`
	}

	GameSchema struct {
		Name    string `json:"gameName"`
		Version string `json:"gameVersion"`
		GameStats struct {
			Achievements []GameSchemaAchievements `json:"achievements"`
		} `json:"availableGameStats"`
	}

	GameSchemaAchievements struct {
		Name         string `json:"name"`
		DefaultValue int8   `json:"defaultvalue"`
		DisplayName  string `json:"displayName"`
		Hidden       int8   `json:"hidden"`
		Icon         string `json:"icon"`
		IconGray     string `json:"icongray"`
	}
)
