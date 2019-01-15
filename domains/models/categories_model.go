package models

import (
	"time"
)

type Categories struct {
	ID        int       `json:"id" xorm:"'id' not null pk autoincr INT(11)"`
	Category  string    `json:"category" xorm:"'category' not null default '' VARCHAR(100)"`
	Timestamp time.Time `json:"timestamp" xorm:"'timestamp' not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
