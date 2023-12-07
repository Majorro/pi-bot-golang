package handlers

import (
	"github.com/majorro/pi-bot/internal/db"
	"testing"
	"time"
)

func TestTryUpdateThing_GrowthWhenLastGrowthWasYesterday(t *testing.T) {
	u := &db.User{
		ThingSize:    10,
		LastGrowthAt: time.Now().AddDate(0, 0, -1),
	}

	growth, ok := tryUpdateThing(u)

	if !ok {
		t.Errorf("Expected growth, got: %d", growth)
	}
}

func TestTryUpdateThing_NoGrowthWhenLastGrowthWasToday(t *testing.T) {
	u := &db.User{
		ThingSize:    10,
		LastGrowthAt: time.Now(),
	}

	growth, ok := tryUpdateThing(u)

	if ok {
		t.Errorf("Expected no growth, got: %d", growth)
	}
}

func TestTryUpdateThing_NoGrowthWhenDifferentTimezones(t *testing.T) {
	var date time.Time
	tDate := time.Now()
	_, offset := tDate.Zone()

	if tDate.Hour() < 12+offset {
		date = tDate.In(time.FixedZone("UTC-12", -12*60*60))
	} else {
		date = tDate.In(time.FixedZone("UTC+12", 12*60*60))
	}

	u := &db.User{
		ThingSize:    10,
		LastGrowthAt: date,
	}

	growth, ok := tryUpdateThing(u)

	if ok {
		t.Errorf("Expected no growth, got: %d", growth)
	}
}
