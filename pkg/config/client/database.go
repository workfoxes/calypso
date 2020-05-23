package client

import (
	"bytes"
	"fmt"
	error2 "github.com/airstrik/gobase/pkg/error"
	"github.com/airstrik/gobase/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"os/exec"
)

var (
	dialect string
	host string
	port string
	user string
	password string
	err error
)

func init(){
	dialect = os.Getenv("_DB_DIALECT")
	host = os.Getenv("_DB_HOST")
	port= os.Getenv("_DB_PORT")
	user = os.Getenv("_DB_USER")
	password = os.Getenv("_DB_PASSWORD")
}

type Database struct {
	Dialect string
	userName string
	password string
	host string
	port string
	database string
	DB *gorm.DB
}

func (db *Database) getArgs() string {
	if db.database == "" {
		db.database = "account"
	}
	switch db.Dialect {
	case "postgres":
		con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.host, db.port,
			db.userName, db.password, db.database)
		return  con
	default:
		return ""
	}
}

func (db *Database) CreateDB(database string) {
	cmd := exec.Command("createdb", "-p", db.port, "-h", db.host, "-U", db.userName, "-e", database)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Output: %q\n", out.String())
}

func NewDatabase(database string) *Database {
	_db := &Database{
		Dialect: dialect,
		userName: user,
		password: password,
		host: host,
		port: port,
		database: database,
	}
	_db.DB, err = gorm.Open(dialect, _db.getArgs())
	utils.HandleCustomError(err, error2.TenantNotFound)
	return _db
}