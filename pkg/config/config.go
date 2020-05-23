package config

import (
	"github.com/airstrik/gobase/pkg/config/client"
)

var (
	conn string
)

type Context struct {
	AccountId string
	DB *client.Database
	BaseDB *client.Database
	Cache *client.RedisClient
}

func NewContext(account string) *Context {
	baseDB := client.NewDatabase(account)
	return &Context{
		AccountId:	account,
		DB: 		baseDB,
		BaseDB: 	client.NewDatabase(""),
		//Cache: 		client.NewRedisClient(account),
	}
}