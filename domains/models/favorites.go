package models

type Favorites struct {
	ID     int `json:"id" xorm:"not null pk autoincr INT(11)"`
	WordID int `json:"wordid" xorm:"not null index INT(11)"`
	UserID int `json:"userid" xorm:"not null index INT(11)"`
}
