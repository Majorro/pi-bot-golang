package handlers

import (
	"github.com/go-pg/pg/v10"
	tele "gopkg.in/telebot.v3"
	"log"
)

type Handler interface {
	GetCommand() string
	Handle(ctx tele.Context, d *pg.DB) error
}

func AddAll(b *tele.Bot, pgDb *pg.DB) {
	handlers := []Handler{
		Grow{},
	}

	for _, h := range handlers {
		comm := h.GetCommand()
		b.Handle(comm, func(c tele.Context) error {
			err := h.Handle(c, pgDb)
			if err != nil {
				log.Printf("%s: %v", comm, err)
				return c.Send("ВСЕ В ДЕРЬМЕ")
			}

			return nil
		})
	}
}
