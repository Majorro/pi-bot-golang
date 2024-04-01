package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"os"
)

func InitAndConnect(callback func(db *pg.DB) error) error {
	db := pg.Connect(&pg.Options{ // TODO: add custom database creation
		Addr:     os.Getenv("PI_BOT_DB_HOST") + ":" + os.Getenv("PI_BOT_DB_PORT"),
		User:     os.Getenv("PI_BOT_DB_USER"),
		Password: os.Getenv("PI_BOT_DB_PASSWORD"),
	})
	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return errors.New("can't connect to db")
	}
	log.Println("Connected to DB")

	err := createSchema(db)
	if err != nil {
		return fmt.Errorf("db init err: %w", err)
	}
	log.Println("Schema created")

	return callback(db)
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Megaprikol)(nil),
	}

	for _, m := range models {
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	// TODO: refactor
	_, err := db.Exec("CREATE INDEX IF NOT EXISTS users_thing_size_btree ON users USING btree (thing_size DESC);")
	if err != nil {
		return err
	}

	return nil
}
