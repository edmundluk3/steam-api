package steam_api

import "testing"

func TestSteamApp_Detail(t *testing.T) {
	app := SteamApp{Id: 292030}
	detail, err := app.Detail("zh-cn", "cc")

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	t.Logf("%v", detail)

}

func TestSteamApp_News(t *testing.T) {
	app := SteamApp{Id: 292030}

	news, err := app.News(10, 100)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	t.Logf("%v", news)
}
