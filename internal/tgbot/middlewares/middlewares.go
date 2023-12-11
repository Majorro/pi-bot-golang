package middlewares

import (
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	tele "gopkg.in/telebot.v3"
	"log"
)

func ProvideUser(d *pg.DB) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			command := ctx.Text()

			sender := ctx.Sender()
			log.Printf("%s: %s\n", command, sender.Username)

			var name string
			if sender.LastName == "" {
				name = sender.FirstName
			} else {
				name = sender.FirstName + " " + sender.LastName
			}
			u := &db.User{
				Id:        sender.ID,
				FullName:  name,
				Username:  sender.Username,
				ThingSize: 0,
			}

			u, err := db.GetOrInsertUser(d, u)
			if err != nil {
				return err
			}
			log.Printf("%s: got user from db - %v\n", command, u)

			ctx.Set("user", u)

			return next(ctx)
		}
	}
}
