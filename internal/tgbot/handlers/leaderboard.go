package handlers

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	tele "gopkg.in/telebot.v3"
	"strings"
)

type leaderboard struct{}

func (h leaderboard) getCommand() string {
	return "/leaderboard"
}

func (h leaderboard) handle(ctx tele.Context, d *pg.DB) error {
	u, err := handleFirstUserInteraction(h, ctx, d)
	if err != nil {
		return err
	}

	var allUsers []db.User
	allUsers, err = db.GetOrderedUsers(d)
	if err != nil {
		return err
	}

	rowTemplate := "%d. %s — _%d_\n"
	curUserRowTemplate := "%d. *%s* — _%d_\n"
	var builder strings.Builder
	builder.WriteString("*Топ штуковин*🤯\n\n")
	for i, usr := range allUsers {
		var t string
		if usr.Id == u.Id {
			t = curUserRowTemplate
		} else {
			t = rowTemplate
		}

		_, err := fmt.Fprintf(&builder, t, i+1, usr.FullName, usr.ThingSize)
		if err != nil {
			return err
		}
	}

	return ctx.Send(builder.String())
}
