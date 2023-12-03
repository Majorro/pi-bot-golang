package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"time"
)

type User struct {
	Id int64 `pg:",pk"`

	Username     string
	ThingSize    int `pg:",default:0" pg:",notnull"`
	LastGrowthAt time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User %d: %s, %d", u.Id, u.Username, u.ThingSize)
}

func GetUser(db *pg.DB, u *User) error {
	return db.Model(u).WherePK().Select()
}

func InsertUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).Insert()
	return err
}

func GetOrInsertUser(db *pg.DB, u *User) (*User, error) {
	err := GetUser(db, u)
	if err != nil {
		switch err {
		case pg.ErrNoRows:
			insertErr := InsertUser(db, u)
			if insertErr != nil {
				return u, insertErr
			}
		default:
			return u, err
		}
	}

	return u, nil
}

func UpdateUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).WherePK().Update()
	return err
}
