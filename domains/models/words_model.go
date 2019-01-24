package models

import (
	"time"
)

type Words struct {
	Id         int       `xorm:"not null pk autoincr INTEGER"`
	Word       string    `xorm:"not null TEXT"`
	Author     string    `xorm:"not null TEXT"`
	CategoryId int       `xorm:"not null INTEGER"`
	Timestamp  time.Time `xorm:"not null default 'now()' DATETIME"`
}

type Word struct {
	ID              int    `json:"id"`
	Word            string `json:"word"`
	Author          string `json:"author"`
	Category        string `json:"category"`
	Nice            int64  `json:"nice"`
	Favortite       int64  `json:"favortite"`
	NiceStatus      bool   `json:"nice_status"`
	FavortiteStatus bool   `json:"favortite_status"`
}
