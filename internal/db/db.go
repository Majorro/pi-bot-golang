package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

type User struct {
	Id int64 `pg:",pk"`

	Username  string
	ThingSize int
}

func (u User) String() string {
	return fmt.Sprintf("User %d: %s, %d", u.Id, u.Username, u.ThingSize)
}

func InitAndConnect(callback func(db *pg.DB)) {
	db := pg.Connect(&pg.Options{
		Addr:     "db:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "pibot",
	})
	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")

	err := createSchema(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Schema created")

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

func GetOrInsertUser(db *pg.DB, u *User) error {
	err := db.Model(u).WherePK().Select()
	if err != nil {
		log.Println(err, u)
		err = InsertUser(db, u)
		return err
	}

	return nil
}

func InsertUser(db *pg.DB, u *User) error {
	_, err := db.Model(&u).Insert()
	return err
}

func UpdateUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).Update()
	return err
}
