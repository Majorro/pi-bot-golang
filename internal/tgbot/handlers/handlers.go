package handlers

import (
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/tgbot/middlewares"
	tele "gopkg.in/telebot.v3"
	"log"
)

type handler interface {
	getCommand() string
	handle(ctx tele.Context, d *pg.DB) error
}

func AddAll(b *tele.Bot, pgDb *pg.DB) {
	handlers := []handler{
		grow{},
		leaderboard{},
	}

	b.Use(middlewares.ProvideUser(pgDb)) // TODO: restrict to db handlers only

	for _, h := range handlers {
		h := h
		comm := h.getCommand()
		b.Handle(comm, func(c tele.Context) error {
			err := h.handle(c, pgDb)
			if err != nil {
				log.Printf("%s: %v", comm, err)
				return c.Send("ВСЕ В ДЕРЬМЕ @majorro228")
			}

			return nil
		})
	}
}
