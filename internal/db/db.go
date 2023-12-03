package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

type User struct {
	Id int64 `pg:",unique"`

	Username  string
	ThingSize int `pg:",default:0"`
}

func (u User) String() string {
	return fmt.Sprintf("User %d: %s, %d", u.Id, u.Username, u.ThingSize)
}

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

func GetUser(db *pg.DB, u *User) error {
	return db.Model(u).Where("id = ?id").Select()
}

func InsertUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).Insert()
	return err
}

func UpdateUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).WherePK().Update()
	return err
}
