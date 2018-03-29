package app

import (
	"testing"
	"log"
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
	const ID uint32 = 292030

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