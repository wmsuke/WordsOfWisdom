package repositores

import (
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/wmsuke/WordsOfWisdom/infrastructure/config"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		var c = config.NewConfig()
		connStr = c.DB.Connection
	}
	engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
