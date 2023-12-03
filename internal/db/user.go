package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"time"
)

type User struct {
	Id int64 `pg:",pk"`

	FullName     string `pg:",notnull"`
	Username     string `pg:",notnull"`
	ThingSize    int    `pg:",default:0" pg:",notnull"`
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

func GetOrderedUsers(db *pg.DB) (users []User, err error) {
	err = db.Model(&users).Order("thing_size desc").Limit(100).Select()
	if err != nil {
		return
	}
	return
}

func UpdateUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).WherePK().Update()
	return err
}
