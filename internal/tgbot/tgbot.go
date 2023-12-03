package tgbot

import (
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/tgbot/handlers"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func initBot() (*tele.Bot, error) {
	pref := tele.Settings{
		Token:     os.Getenv("PI_BOT_TOKEN"),
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	}

	return tele.NewBot(pref)
}

func Start(db *pg.DB) error {
	b, err := initBot()
	if err != nil {
		return err
	}
	log.Println("Bot created")

	handlers.AddAll(b, db)
	log.Println("Handlers added")

	log.Println("Bot started")
	b.Start()

	return nil
}
