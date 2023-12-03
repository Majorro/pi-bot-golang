package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

type User struct {
	Id int64 `pg:",unique"`

	Username  string
	ThingSize int `pg:",default:0"`
}

func (u User) String() string {
	return fmt.Sprintf("User %d: %s, %d", u.Id, u.Username, u.ThingSize)
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
