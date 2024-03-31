package middlewares

import (
	"fmt"
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
			log.Printf("%s sent %s\n", sender.Username, command)

			var name string
			if sender.LastName == "" {
				name = sender.FirstName
			} else {
				name = sender.FirstName + " " + sender.LastName
			}
			u := &db.User{
				Id:       sender.ID,
				FullName: name,
				Username: sender.Username,
			}

			u, err := db.GetOrUpsertUser(d, u)
			if err != nil {
				return fmt.Errorf("error providing user for context: %w", err)
			}
			log.Printf("%s: got user from db - %v\n", command, u)

			ctx.Set("user", u)

			return next(ctx)
		}
	}
}
