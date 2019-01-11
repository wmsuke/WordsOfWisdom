package models

type Favorites struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	WordId int `xorm:"not null index INT(11)"`
	UserId int `xorm:"not null index INT(11)"`
}
