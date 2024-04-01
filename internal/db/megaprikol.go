package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"strings"
)

type Megaprikol struct {
	Id int64 `pg:",pk"`

	Content    string `pg:",default:'',notnull"`
	PingsCount int    `pg:",default:0,notnull"`
	UsageCount int    `pg:",default:0,notnull"`
}

func (m Megaprikol) String() string {
	return fmt.Sprintf("%s: %d", strings.ReplaceAll(m.Content, "\n", "  ")[:20], m.UsageCount)
}

func GetRandomMegaprikol(db *pg.DB) (*Megaprikol, error) {
	var ms []Megaprikol
	err := db.Model(&ms).OrderExpr("RANDOM()").Limit(1).Select()
	if err != nil {
		return nil, fmt.Errorf("error getting random megaprikol: %w", err)
	}
	return &ms[0], nil
}

func UpdateMegaprikol(db *pg.DB, m *Megaprikol) error {
	_, err := db.Model(m).WherePK().UpdateNotZero()
	if err != nil {
		return fmt.Errorf("error updating megaprikol: %w", err)
	}
	return nil
}

func GetUnusedMegaprikols(db *pg.DB) (ms []Megaprikol, err error) {
	err = db.Model(&ms).Where("usage_count = 0").Select()
	if err != nil {
		return nil, fmt.Errorf("error getting unused megaprikols: %w", err)
	}
	return
}

func GetUnseenMegaprikolsCount(db *pg.DB) (int, error) {
	count, err := db.Model((*Megaprikol)(nil)).Where("usage_count = 0").Count()
	if err != nil {
		return 0, fmt.Errorf("error getting unseen megaprikols count: %w", err)
	}
	return count, nil
}
