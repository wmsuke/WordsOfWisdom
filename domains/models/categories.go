package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Categories struct {
	Id        int       `json:"id" xorm:"'id' not null pk autoincr INT(11)"`
	Category  string    `json:"category" xorm:"'category' not null default '' VARCHAR(100)"`
	Timestamp time.Time `json:"timestamp" xorm:"'timestamp' not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

var engine *xorm.Engine

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:admin@/WordsOfWisdom")
	if err != nil {
		panic(err)
	}
}

// CategoryRepository is
type CategoryRepository struct {
}

// NewCategoryRepository ...
func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{}
}

// GetByID ...
func (m CategoryRepository) GetByID(id int) *Categories {
	var category = Categories{Id: id}
	has, _ := engine.Get(&category)
	if has {
		return &category
	}

	return nil
}
