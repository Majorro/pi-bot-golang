package handlers

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	tele "gopkg.in/telebot.v3"
	"log"
	"math"
	"math/rand"
)

type Grow struct{}

func (h Grow) GetCommand() string {
	return "/grow"
}

func (h Grow) Handle(ctx tele.Context, d *pg.DB) error {
	sender := ctx.Sender()
	log.Printf("%s: %s-%d\n", h.GetCommand(), sender.Username, sender.ID)

	u := &db.User{
		Id:       sender.ID,
		Username: sender.Username,
	}

	err := db.GetUser(d, u)
	if err != nil {
		switch err {
		case pg.ErrNoRows:
			insertErr := db.InsertUser(d, u)
			if insertErr != nil {
				return insertErr
			}
		default:
			return err
		}
	}
	log.Printf("%s: got user from db - %v\n", h.GetCommand(), u)

	growth := getThingGrowth()
	u.ThingSize += growth

	err = db.UpdateUser(d, u)
	if err != nil {
		return err
	}
	log.Printf("%s: updated user - %v\n", h.GetCommand(), u)

	var msg string
	if growth > 0 {
		msg = `@%s, ваша штуковина выросла на %d см!!!
теперь её размер %d см!!!`
	} else {
		msg = `@%s, ваша штуковина уменьшилась на %d см!!!
теперь её размер %d см!!!`
	}
	return ctx.Send(fmt.Sprintf(msg, u.Username, growth, abs(u.ThingSize)))
}

func getThingGrowth() int {
	stdDev := 4.0
	mean := 3.0
	return int(math.Round(rand.NormFloat64()*stdDev + mean))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
