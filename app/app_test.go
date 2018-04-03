package app

import (
	"testing"
	"log"
	"github.com/edmundluk3/steam-api/model"
)

func TestList(t *testing.T) {
	app, err := List()
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(app.Apps.AppList) <=0 {
		t.Errorf(err.Error())
	}
	log.Print(app.Apps.AppList[0])
}

func TestDetail(t *testing.T) {
	const ID model.SteamAppID = 292030

	detail, err := Detail(ID, "zh-cn", "cc")
	if err != nil {
		log.Fatal(err)
		return
	}

	if detail.Appid != ID {
		t.Errorf("%v", detail)
	}
	log.Printf("%d: %s" ,detail.Appid, detail.Name)
}

func TestPrice(t *testing.T) {
	ids :=  []model.SteamAppID{292030, 292939, 292025}

	price, err := Price(ids, "cc")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("%v", price)
}

func TestDetailBadData(t *testing.T) {
	ids := []model.SteamAppID{92, 70}

	for _, i := range ids {
		detail, err := Detail(i, "zh-cn", "cc")
		if err != nil {
			t.Errorf("%v", err)
		}

		log.Printf("%v", detail)
	}
}