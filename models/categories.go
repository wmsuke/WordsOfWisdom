package models

import (
	"time"
)

type Categories struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Category  string    `xorm:"not null default '' VARCHAR(100)"`
	Timestamp time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
