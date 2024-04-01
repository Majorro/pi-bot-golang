package handlers

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	"github.com/majorro/pi-bot/internal/tgbot/utils"
	tele "gopkg.in/telebot.v3"
)

type megaprikol struct{} // TODO: remove? 😼

func (h megaprikol) getCommand() string { return "/megaprikol" }

func (h megaprikol) handle(ctx tele.Context, d *pg.DB) error {
	mp, err := db.GetRandomMegaprikol(d)
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}

	err = db.IncrementMegaprikolUsageCount(d, mp)
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}

	var users []db.User
	users, err = db.GetRandomUsers(d, mp.PingsCount)
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}

	var pings []interface{}
	for _, u := range users {
		pings = append(pings, u.Username)
	}

	err = ctx.Send(utils.Format(mp.Content, pings...))
	if err != nil {
		return fmt.Errorf("%s err: %w", h.getCommand(), err)
	}
	return nil
}

//func (h megaprikol) handle(ctx tele.Context, d *pg.DB) error {
//	var mps []db.Megaprikol
//	d.Model(&mps).Select()
//
//	for _, mp := range mps {
//		users, err := db.GetRandomUsers(d, mp.PingsCount)
//		if err != nil {
//			return fmt.Errorf("%s err: %w", h.getCommand(), err)
//		}
//
//		var pings []interface{}
//		for _, u := range users {
//			pings = append(pings, u.Username)
//		}
//
//		err = ctx.Send(utils.Format(mp.Content, pings...))
//		if err != nil {
//			return fmt.Errorf("%s err: %w", h.getCommand(), err)
//		}
//	}
//	return nil
//}
