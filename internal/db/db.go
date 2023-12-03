package db

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

func InitAndConnect(callback func(db *pg.DB) error) error {
	db := pg.Connect(&pg.Options{
		Addr:     "db:5432",
		User:     "postgres",
		Password: "postgres",
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
