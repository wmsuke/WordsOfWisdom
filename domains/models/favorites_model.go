package models

type Favorites struct {
	Id     int `xorm:"not null pk autoincr INTEGER"`
	WordId int `xorm:"not null INTEGER"`
	UserId int `xorm:"not null INTEGER"`
}
