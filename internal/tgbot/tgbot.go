package tgbot

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	"log"
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
		u := &db.User{
			Id:       sender.ID,
			Username: sender.Username,
		}
		err := db.GetOrInsertUser(pgDb, u)
		if err != nil {
			log.Println(err)
			return c.Send("ВСЕ В ДЕРЬМЕ")
		}

		growth := rand.Intn(20) - 10
		u.ThingSize += growth

		err = db.UpdateUser(pgDb, u)
		if err != nil {
			log.Println(err)
			return c.Send("ВСЕ В ДЕРЬМЕ")
		}

		msg := `@%s, ваша штуковина выросла на %d см!!!
теперь её размер %d см!!!`
		return c.Send(fmt.Sprintf(msg, u.Username, growth, u.ThingSize))
	})
}

func Start(db *pg.DB) {
	b, err := initBot()
	if err != nil {
		log.Fatal(err)
	}

	addHandlers(b, db)
	b.Start()
	log.Println("Bot started")
}
