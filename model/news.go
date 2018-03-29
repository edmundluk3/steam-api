package model

type (
	NewsList struct {
		AppNews AppNews `json:"appnews"`
	}
	AppNews struct {
		AppId     SteamAppID `json:"appid"`
		NewsItems []NewsItem `json:"newsitems"`
		Count     uint       `json:"count"`
	}

	NewsItem struct {
		Gid           string     `json:"gid"`
		Title         string     `json:"title"`
		Url           string     `json:"url"`
		IsExternalUrl bool       `json:"is_external_url"`
		Author        string     `json:"author"`
		Contents      string     `json:"contents"`
		FeedLabel     string     `json:"feedlabel"`
		Date          uint       `json:"date"`
		FeedName      string     `json:"feedname"`
		FeedType      uint8      `json:"feed_type"`
		Appid         SteamAppID `json:"appid"`
	}
)
