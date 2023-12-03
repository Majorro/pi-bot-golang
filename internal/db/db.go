package db

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"os"
)

func InitAndConnect(callback func(db *pg.DB) error) error {
	db := pg.Connect(&pg.Options{ // TODO: add custom database creation
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return err
	}
	log.Println("Connected to DB")

	err := createSchema(db)
	if err != nil {
		return err
	}
	log.Println("Schema created")

	return callback(db)
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, m := range models {
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
