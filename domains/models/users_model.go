package models

import (
	"time"
)

type Users struct {
	ID         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Key        string    `json:"key" xorm:"not null VARCHAR(384)"`
	AccessDate time.Time `json:"accessdate" xorm:"not null DATETIME"`
}
