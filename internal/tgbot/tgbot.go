package tgbot

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func initBot() (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  os.Getenv("PI_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	return tele.NewBot(pref)
}

func addHandlers(b *tele.Bot, pgDb *pg.DB) {
	b.Handle("/grow", func(c tele.Context) error {
		sender := c.Sender()
		log.Printf("/grow: %s-%d\n", sender.Username, sender.ID)

		u := &db.User{
			Id:       sender.ID,
			Username: sender.Username,
		}

		err := db.GetUser(pgDb, u)
		if err != nil {
			switch err {
			case pg.ErrNoRows:
				insertErr := db.InsertUser(pgDb, u)
				if insertErr != nil {
					log.Println(insertErr)
					return c.Send("ВСЕ В ДЕРЬМЕ")
				}
			default:
				log.Println(err)
				return c.Send("ВСЕ В ДЕРЬМЕ")
			}
		}
		log.Printf("/grow: got user from db - %v\n", u)

		growth := getThingGrowth()
		u.ThingSize += growth

		err = db.UpdateUser(pgDb, u)
		if err != nil {
			log.Println(err)
			return c.Send("ВСЕ В ДЕРЬМЕ")
		}
		log.Printf("/grow: updated user - %v\n", u)

		var msg string
		if growth > 0 {
			msg = `@%s, ваша штуковина выросла на %d см!!!
теперь её размер %d см!!!`
		} else {
			msg = `@%s, ваша штуковина уменьшилась на %d см!!!
теперь её размер %d см!!!`
		}
		return c.Send(fmt.Sprintf(msg, u.Username, growth, abs(u.ThingSize)))
	})
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

func Start(db *pg.DB) error {
	b, err := initBot()
	if err != nil {
		return err
	}
	log.Println("Bot created")

	addHandlers(b, db)
	log.Println("Handlers added")

	log.Println("Bot started")
	b.Start()

	return nil
}
