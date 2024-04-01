package handlers

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	tele "gopkg.in/telebot.v3"
)

type megaprikolUnseen struct{} // TODO: remove? 😼

func (h megaprikolUnseen) getCommand() string { return "/megaprikol_unseen" }

func (h megaprikolUnseen) handle(ctx tele.Context, d *pg.DB) error {
	count, err := db.GetUnseenMegaprikolsCount(d)
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}

	err = ctx.Send(fmt.Sprintf("ВЫ ЕЩЕ СТОЛЬКО МЕГАПРИКОЛОВ НЕ ВИДЕЛИ: %d", count))
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}
	return nil
}
