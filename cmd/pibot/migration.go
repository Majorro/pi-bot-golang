package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/majorro/pi-bot/internal/db"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
	"log"
	"os"
)

var migrationsDir = "internal/db/migrations"

// TODO: maybe use https://gorm.io/docs/ + https://atlasgo.io/guides/orms/gorm instead of go-pg
func main() {
	err := db.InitAndConnect(func(db *pg.DB) error {
		return migrations.Run(db, migrationsDir, os.Args)
	})
	if err != nil {
		log.Fatalln(err)
	}
}
