package models

import (
	"time"
)

type Users struct {
	Id         int       `xorm:"not null pk autoincr INTEGER"`
	Key        string    `xorm:"not null TEXT"`
	AccessDate time.Time `xorm:"not null DATETIME"`
}
