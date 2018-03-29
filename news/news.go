package news

import (
	"github.com/edmundluk3/steam-api/model"
	"github.com/edmundluk3/steam-api/api"
	"encoding/json"
)

func List (id model.SteamAppID, count uint, max uint) (*model.AppNews, error) {
	resByte, err := api.GetNewsList(id, count, max)

	if err != nil {
		return nil, err
	}

	var resp model.NewsList

	parseErr := json.Unmarshal(resByte, &resp)

	if err != nil {
		return nil, parseErr
	}

	return &resp.AppNews, nil
}