package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:admin@/WordsOfWisdom")
	if err != nil {
		panic(err)
	}
}
