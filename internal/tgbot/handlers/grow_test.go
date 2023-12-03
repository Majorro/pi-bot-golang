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

	if !ok || growth == 0 {
		t.Errorf("Expected growth, got: %d", growth)
	}
}

func TestTryUpdateThing_NoGrowthWhenLastGrowthWasToday(t *testing.T) {
	u := &db.User{
		ThingSize:    10,
		LastGrowthAt: time.Now(),
	}

	growth, ok := tryUpdateThing(u)

	if ok || growth != 0 {
		t.Errorf("Expected no growth, got: %d", growth)
	}
}
