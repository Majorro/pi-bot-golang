package handlers

import (
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
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

	for _, h := range handlers {
		h := h
		comm := h.getCommand()
		b.Handle(comm, func(c tele.Context) error {
			err := h.handle(c, pgDb)
			if err != nil {
				log.Printf("%s: %v", comm, err)
				return c.Send("ВСЕ В ДЕРЬМЕ")
			}

			return nil
		})
	}
}

func handleFirstUserInteraction(h handler, ctx tele.Context, d *pg.DB) (u *db.User, err error) {
	sender := ctx.Sender()
	log.Printf("%s: %s-%d\n", h.getCommand(), sender.Username, sender.ID)

	var name string
	if sender.LastName == "" {
		name = sender.FirstName
	} else {
		name = sender.FirstName + " " + sender.LastName
	}
	u = &db.User{
		Id:        sender.ID,
		FullName:  name,
		Username:  sender.Username,
		ThingSize: 0,
	}

	u, err = db.GetOrInsertUser(d, u)
	if err != nil {
		return
	}
	log.Printf("%s: got user from db - %v\n", h.getCommand(), u)

	return
}
