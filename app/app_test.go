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
