package models

import (
	"time"
)

type Words struct {
	ID         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Word       string    `json:"word" xorm:"not null default '' VARCHAR(1000)"`
	Author     string    `json:"author" xorm:"not null default '' VARCHAR(100)"`
	CategoryID int       `json:"categoryid" xorm:"not null index INT(11)"`
	Nice       int       `json:"nice" xorm:"not null default 0 INT(11)"`
	Timestamp  time.Time `json:"timestamp" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
