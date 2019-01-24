package repositores

import (
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}
