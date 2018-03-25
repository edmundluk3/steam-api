package api

import (
	"strconv"
	"fmt"
)

func GetAppList() ([]byte, error){
	return ApiGet(
		STEAM_API_URL,
		APP_ENDPOINT,
		"/GetAppList",
		"/v2",
		nil,
	)
}

func GetAppDetail(appid int, language string, cc string) ([]byte, error) {
	payload := map[string]string{
		"appids": strconv.Itoa(appid),
		"cc": cc,
		"l": language,
	}

	return ApiGet(
		"https://store.steampowered.com/api/appdetails",
		"",
		"",
		"",
		payload,
	)
}

func GetAppPrice(appids []int, cc string) ([]byte, error){
	var ids string
	for _, v := range appids {
		ids += fmt.Sprintf("%d", v)
	}

	payload := map[string]string{
		"appids": ids,
		"cc": cc,
		"filters": "price_overview",
	}


	return ApiGet(
		"https://store.steampowered.com/api/appdetails",
		"",
		"",
		"",
		payload,
	)
}