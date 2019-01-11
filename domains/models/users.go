package models

import (
	"time"
)

type Users struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Key        string    `xorm:"not null VARCHAR(384)"`
	AccessDate time.Time `xorm:"not null DATE"`
}
