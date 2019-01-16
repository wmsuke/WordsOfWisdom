package models

import (
	"time"
)

type Words struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Word       string    `xorm:"not null default '' VARCHAR(1000)"`
	Author     string    `xorm:"not null default '' VARCHAR(100)"`
	CategoryId int       `xorm:"not null index INT(11)"`
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
