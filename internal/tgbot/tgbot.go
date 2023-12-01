package tgbot

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	"log"
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
		return c.Send("Your thing has grown!")
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
