package main

import (
	"github.com/majorro/pi-bot/internal/db"
	pibot "github.com/majorro/pi-bot/internal/tgbot"
	"log"
)

func main() {
	err := db.InitAndConnect(pibot.Start)
	if err != nil {
		log.Fatal(err)
	}
}
