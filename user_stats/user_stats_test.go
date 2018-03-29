package user_stats

import (
	"testing"
	"log"
)

func TestGlobalAchievementsPercentages(t *testing.T) {
	list, err := GlobalAchievementsPercentages(292030)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("%v", list)
}