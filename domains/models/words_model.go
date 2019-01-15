package models

import (
	"time"
)

type Words struct {
	ID         int       `xorm:"not null pk autoincr INT(11)"`
	Word       string    `xorm:"not null default '' VARCHAR(1000)"`
	Author     string    `xorm:"not null default '' VARCHAR(100)"`
	CategoryID int       `xorm:"not null index INT(11)"`
	Nice       int       `xorm:"not null default 0 INT(11)"`
	Timestamp  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type Word struct {
	ID        int    `json:"id"`
	Word      string `json:"word"`
	Author    string `json:"author"`
	Category  string `json:"category"`
	Nice      int    `json:"nice"`
	Favortite int    `json:"favortite"`
}

func FindOne(id int) *Words {
	// var category = models.Categories{ID: id}
	var word = Words{ID: id}
	has, _ := engine.Get(&word)
	if has {
		return &word
	}

	return nil
}
