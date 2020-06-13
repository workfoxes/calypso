package db

import (
	"github.com/itsparser/greeny-be/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var (
	dialect string
)

func init() {
	dialect = os.Getenv("_DB_DIALECT")

}

type Database struct {
	Dialect string
	DB      *gorm.DB
}

func getArgs(dialect string) string {
	switch dialect {
	case "postgres":
		return "user:postgres@dockerpassword/postgres"
	default:
		return ""
	}
}

func NewDatabase() *Database {
	db, err := gorm.Open(dialect, getArgs(dialect))
	utils.HandleError(err)
	return &Database{
		Dialect: dialect,
		DB:      db,
	}
}
