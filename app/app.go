package app

import (
	"steam-api/api"
	steamErr "steam-api/error"
	"encoding/json"
)

type App struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps struct {
		AppList []App `json:"apps"`
	} `json:"applist"`
}

func List() (*AppList, error) {
	byte, err := api.GetAppList()

	if err != nil {
		return nil, err
	}

	var resp AppList

	if err := json.Unmarshal(byte, &resp); err != nil {
		return nil, &steamErr.SteamAPIError{Code: 600, Msg: "json parse error"}
	}

	return &resp, nil
}
