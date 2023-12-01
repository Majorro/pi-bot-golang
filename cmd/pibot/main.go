package main

import (
	"github.com/majorro/pi-bot/internal/db"
	pibot "github.com/majorro/pi-bot/internal/tgbot"
)

func main() {
	db.InitAndConnect(pibot.Start)
}
