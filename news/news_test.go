package news

import (
	"testing"
	"log"
)

func TestList(t *testing.T) {
	const (
		count = 10
		max = 200
	)
	list, err := List(292030, count, max)

	if err != nil {
		log.Fatal(err)
	}

	if len(list.AppNews.NewsItems) != count {
		t.Errorf("count: %d not equal with %d", len(list.AppNews.NewsItems), count)
	}
}