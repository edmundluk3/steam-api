package app

import (
	"github.com/edmundluk3/steam-api/api"
	steamErr "github.com/edmundluk3/steam-api/error"
	"encoding/json"
	"strconv"
	_ "log"
)

type App struct {
	Appid uint32 `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps struct {
		AppList []App `json:"apps"`
	} `json:"applist"`
}

type (
	AppDetailList map[string]AppData

	AppData struct {
		Success bool      `json:"success"`
		Data    AppDetail `json:"data"`
	}
	AppDetail struct {
		Type                string             `json:"type"`
		Name                string             `json:"name"`
		Appid               uint32             `json:"steam_appid"`
		RequiredAge         int                `json:"required_age"`
		IsFree              bool               `json:"is_free"`
		ControllerSupport   string             `json:"controller_support"`
		DLC                 []uint32           `json:"dlc"`
		DetailedDescription string             `json:"detailed_description"`
		AboutTheGame        string             `json:"about_the_game"`
		ShortDescription    string             `json:"short_description"`
		SupportedLanguages  string             `json:"supported_languages"`
		Reviews             string             `json:"reviews"`
		HeaderImage         string             `json:"header_image"`
		Website             string             `json:"website"`
		PCRequirements      map[string]string  `json:"pc_requirements"`
		MacRequirements     []interface{}      `json:"mac_requirements"`
		LinuxRequirements   []interface{}      `json:"linux_requirements"`
		LegalNotice         string             `json:"legal_notice"`
		Developers          []string           `json:"developers"`
		Publishers          []string           `json:"publishers"`
		PriceOverview       AppPrice           `json:"price_overview"`
		Packages            []interface{}      `json:"packages"`
		Metacritic          AppMetacritic      `json:"metacritic"`
		Categories          []AppCategory      `json:"categories"`
		Genres              []AppGenres        `json:"genres"`
		Screenshots         []AppScreenshots   `json:"screenshots"`
		Recommendations     AppRecommendations `json:"recommendations"`
		Achievements        AppAchievements    `json:"achievements"`
		// Platforms
		// package_groups
		ReleaseDate AppReleaseDate `json:"release_date"`
		SupportInfo AppSupportInfo `json:"support_info"`
		Background  string         `json:"background"`
	}
	AppPrice struct {
		Currency           string `json:"currency"`
		Initial            uint   `json:"initial"`
		Final              uint   `json:"final"`
		DiscountPercentage uint8  `json:"discount_percentage"`
	}

	AppPlatform struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	}
	AppMetacritic struct {
		Score uint8  `json:"score"`
		Url   string `json:"url"`
	}

	AppCategory struct {
		Id          uint   `json:"id"`
		Description string `json:"description"`
	}
	AppGenres struct {
		Id          string `json:"id"`
		Description string `json:"description"`
	}
	AppScreenshots struct {
		Id            uint   `json:"id"`
		PathThumbnail string `json:"path_thumbnail"`
		PathFull      string `json:"path_full"`
	}
	AppRecommendations struct {
		Total uint `json:"total"`
	}
	AppAchievementsItem struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}
	AppAchievements struct {
		Total       uint                  `json:"total"`
		Highlighted []AppAchievementsItem `json:"highlighted"`
	}
	AppReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	}
	AppSupportInfo struct {
		Url   string `json:"url"`
		Email string `json:"email"`
	}
)

func List() (*AppList, error) {
	resByte, err := api.GetAppList()

	if err != nil {
		return nil, err
	}

	var resp AppList

	if err := json.Unmarshal(resByte, &resp); err != nil {
		return nil, &steamErr.SteamAPIError{Code: 600, Msg: "json parse error"}
	}

	return &resp, nil
}

func Detail(id int, language string, cc string) (*AppDetail, error) {
	resByte, err := api.GetAppDetail(id, language, cc)
	if err != nil {
		return nil, err
	}

	var (
		list AppDetailList
		// data AppData
		resp AppDetail
	)

	if err := json.Unmarshal(resByte, &list); err != nil {
		return nil, err
	}

	found := false

	for k, v := range list {
		if i, err := strconv.Atoi(k); !found && err == nil && i == id {
			resp = v.Data
			found = true
		}
	}

	if (!found) {
		return &AppDetail{}, nil
	}

	return &resp, nil


}
