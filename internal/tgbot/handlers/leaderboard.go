package handlers

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	"github.com/majorro/pi-bot/internal/tgbot/utils"
	tele "gopkg.in/telebot.v3"
	"strings"
)

type leaderboard struct{}

func (h leaderboard) getCommand() string {
	return "/leaderboard"
}

func (h leaderboard) handle(ctx tele.Context, d *pg.DB) error {
	u := ctx.Get("user").(*db.User)

	allUsers, err := db.GetOrderedUsers(d)
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}

	rowTemplate := "%d. %s — <i>%d</i>\n"
	curUserRowTemplate := "%d. <b>%s</b> — <i>%d</i>\n"
	var builder strings.Builder
	builder.WriteString("<b>Топ штуковин</b>🤯\n\n")
	for i, usr := range allUsers {
		var t string
		if usr.Id == u.Id {
			t = curUserRowTemplate
		} else {
			t = rowTemplate
		}

		_, err := fmt.Fprintf(&builder, t, i+1, utils.EscapeHTML(usr.FullName), usr.ThingSize)
		if err != nil {
			return fmt.Errorf("%s err: %w", h.getCommand(), err)
		}
	}

	err = ctx.Send(builder.String(), &tele.SendOptions{ParseMode: tele.ModeHTML})
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}
	return nil
}
