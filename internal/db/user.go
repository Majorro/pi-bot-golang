package db

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"time"
)

type User struct {
	Id int64 `pg:",pk"`

	FullName     string `pg:",default:'',notnull"`
	Username     string `pg:",default:'',notnull"`
	ThingSize    int    `pg:",default:0,notnull"`
	LastGrowthAt time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User %d: %s, %d", u.Id, u.Username, u.ThingSize)
}

func GetUser(db *pg.DB, u *User) error {
	err := db.Model(u).WherePK().Select()
	if err != nil {
		return fmt.Errorf("error getting user %v: %w", u, err)
	}
	return nil
}

func InsertUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).Insert()
	if err != nil {
		return fmt.Errorf("error inserting user %v: %w", u, err)
	}
	return nil
}

func GetOrUpsertUser(db *pg.DB, u *User) (*User, error) {
	tempUser := &User{Id: u.Id}
	err := GetUser(db, tempUser)
	if err != nil {
		switch {
		case errors.Is(err, pg.ErrNoRows):
			insertErr := InsertUser(db, u)
			if insertErr != nil {
				return nil, fmt.Errorf("GetOrUpsertUser %v: %w", u, insertErr)
			}
		default:
			return nil, fmt.Errorf("GetOrUpsertUser %v: %w", u, err)
		}
	}

	u.ThingSize = tempUser.ThingSize
	u.LastGrowthAt = tempUser.LastGrowthAt
	err = UpdateUser(db, u)
	if err != nil {
		return nil, fmt.Errorf("GetOrUpsertUser %v: %w", u, err)
	}

	return u, nil
}

func GetOrderedUsers(db *pg.DB) (users []User, err error) {
	err = db.Model(&users).Order("thing_size desc").Limit(100).Select()
	if err != nil {
		return nil, fmt.Errorf("error getting ordered users: %w", err)
	}
	return
}

func UpdateUser(db *pg.DB, u *User) error {
	_, err := db.Model(u).WherePK().UpdateNotZero()
	if err != nil {
		return fmt.Errorf("error updating user %v: %w", u, err)
	}
	return nil
}

func GetRandomUsers(db *pg.DB, n int) (users []User, err error) {
	err = db.Model(&users).OrderExpr("RANDOM()").Limit(n).Select()
	if err != nil {
		return nil, fmt.Errorf("error getting random users: %w", err)
	}
	return
}
