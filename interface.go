package steam_api

import (
	"github.com/edmundluk3/steam-api/model"
	"github.com/edmundluk3/steam-api/app"
	"github.com/edmundluk3/steam-api/news"
	"github.com/edmundluk3/steam-api/user_stats"
)

type SteamApp struct {
	Id model.SteamAppID
}

func (a *SteamApp) Detail(language string, cc string) (*model.AppDetail, error) {
	return app.Detail(a.Id, language, cc)
}

// todo
func (a *SteamApp) Price(cc string) (*model.AppPriceList, error) {
	return app.Price([]model.SteamAppID{a.Id}, cc)
}

func (a *SteamApp) News(count uint, max uint) (*model.AppNews, error) {
	return news.List(a.Id, count, max)
}

func (a *SteamApp) GlobalAchievementsPercentages() (*model.AchievementsList, error) {
	return user_stats.GlobalAchievementsPercentages(a.Id)
}

func (a *SteamApp) NumberOfCurrentPlayer () (*model.CurrentNumberOfPlayers, error) {
	return user_stats.NumberOfCurrentPlayer(a.Id)
}

func (a *SteamApp) SchemaForGame(key string, language string) (*model.GameSchema, error) {
	return user_stats.SchemaForGame(a.Id, key, language)
}


