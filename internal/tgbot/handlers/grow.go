package handlers

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	tele "gopkg.in/telebot.v3"
	"log"
	"math"
	"math/rand"
	"time"
)

type grow struct{}

func (h grow) getCommand() string {
	return "/grow"
}

func (h grow) handle(ctx tele.Context, d *pg.DB) error {
	u := ctx.Get("user").(*db.User)

	growth, ok := tryUpdateThing(u)
	if !ok {
		return fmt.Errorf("%s err: %w", h.getCommand(), ctx.Send(fmt.Sprintf("@%s, сегодня уже был рост штуковины!!!", u.Username)))
	}

	err := db.UpdateUser(d, u)
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}
	log.Printf("%s: updated user - %v\n", h.getCommand(), u)

	var msg string
	if growth >= 0 {
		msg = `@%s, ваша штуковина выросла на %d см!!! теперь её размер %d см!!!`
	} else {
		msg = `@%s, ваша штуковина уменьшилась на %d см!!! теперь её размер %d см!!!`
	}
	return fmt.Errorf("%s err: %w", h.getCommand(), ctx.Send(fmt.Sprintf(msg, u.Username, abs(growth), u.ThingSize)))
}

func tryUpdateThing(u *db.User) (int, bool) {
	if u.LastGrowthAt.UTC().YearDay() == time.Now().UTC().YearDay() {
		return 0, false
	}

	u.LastGrowthAt = time.Now().UTC()
	growth := getThingGrowth()
	u.ThingSize += growth

	return growth, true
}

func getThingGrowth() int {
	stdDev := 3.9
	mean := 3.0
	return int(math.Round(rand.NormFloat64()*stdDev + mean))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
