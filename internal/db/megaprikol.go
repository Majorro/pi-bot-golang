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

func IncrementMegaprikolUsageCount(db *pg.DB, m *Megaprikol) error {
	m.UsageCount++
	_, err := db.Model(m).WherePK().Column("usage_count").Update()
	if err != nil {
		return fmt.Errorf("error incrementing megaprikol usage count: %w", err)
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
