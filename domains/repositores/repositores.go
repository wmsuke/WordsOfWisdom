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
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://dev:secret@localhost/words-dev?sslmode=disable"
	}
	engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
