package client

import (
	"fmt"
	error2 "github.com/workfoxes/gobase/pkg/server/error"
	"github.com/workfoxes/gobase/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

type Database struct {
	Dialect  string
	userName string
	password string
	host     string
	port     string
	database string
	DB       *gorm.DB
}

func (db *Database) getArgs() string {
	if db.database == "" {
		db.database = "account"
	}
	switch db.Dialect {
	case "postgres":
		con := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.host, db.port,
			db.userName, db.password, db.database)
		return con
	default:
		return ""
	}
}

func (db *Database) CreateDB(database string) {
	_ = db.DB.Exec("CREATE DATABASE " + database)
	log.Print("Create new Database for account ")
}

func (db *Database) Create(v ...interface{}) {
	db.DB.Create(v)
	db.DB.Model(user).Update("CreatedAt", time.Now())
	db.DB.Model(user).Update("UpdatedAt", time.Now())
	log.Print("Create new Database for account ")
}

func LoadBaseDatabase() *Database {
	if BaseDB == nil {
		log.Print("Creating the Database connection from the Application server")
		BaseDB = LoadDatabase("account")
	}
	return BaseDB
}

func LoadDatabase(database string) *Database {
	_db := &Database{
		Dialect:  dialect,
		userName: user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
	var DbConnection *gorm.DB
	DbConnection, err = gorm.Open(dialect, _db.getArgs())
	_db.DB = DbConnection.Begin()
	utils.HandleCustomError(err, error2.TenantNotFound)
	return _db
}
