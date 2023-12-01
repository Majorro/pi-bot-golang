package tgbot

import (
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

func addHandlers(b *tele.Bot) {
	b.Handle("/grow", func(c tele.Context) error {
		return c.Send("Your thing has grown!")
	})
}

func Start() {
	b, err := initBot()
	if err != nil {
		log.Fatal(err)
		return
	}

	addHandlers(b)
	b.Start()
}
