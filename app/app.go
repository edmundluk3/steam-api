package app

import (
	"github.com/edmundluk3/steam-api/api"
	"github.com/edmundluk3/steam-api/model"
	steamErr "github.com/edmundluk3/steam-api/error"
	steamConst "github.com/edmundluk3/steam-api/constant"
	"encoding/json"
	"strconv"
	_ "log"
)

func List() (*model.AppList, error) {
	resByte, err := api.GetAppList()

	if err != nil {
		return nil, err
	}

	var resp model.AppList

	if err := json.Unmarshal(resByte, &resp); err != nil {
		return nil, &steamErr.SteamAPIError{Code: 600, Msg: "json parse error"}
	}

	return &resp, nil
}

func Detail(id model.SteamAppID, language string, cc string) (*model.AppDetail, error) {
	resByte, err := api.GetAppDetail(id, steamConst.LanguageCode[language], cc)
	if err != nil {
		return nil, err
	}

	var (
		list model.AppDetailList
		// data AppData
		resp model.AppDetail
	)

	if err := json.Unmarshal(resByte, &list); err != nil {
		return nil, err
	}

	found := false

	for k, v := range list {
		if i, err := strconv.ParseInt(k, 10, 32); !found && err == nil && model.SteamAppID(i) == id {
			resp = v.Data
			found = true
		}
	}

	if !found {
		return &model.AppDetail{}, nil
	}

	return &resp, nil

}

func Price(ids []model.SteamAppID, cc string) (*model.AppPriceList, error) {
	resByte, err := api.GetAppPrice(ids, cc)

	if err != nil {
		return nil, err
	}

	var (
		list model.AppPriceList
	)

	if err := json.Unmarshal(resByte, &list); err != nil {
		return nil, err
	}

	for k, v := range list {
		if !v.Success {
			delete(list, k)
		}
	}

	return &list, nil
}
