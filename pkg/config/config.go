package config

import (
	"github.com/airstrik/gobase/pkg/config/client"
	"github.com/go-chi/chi"
	"net/http"
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

func LoadContext(r *http.Request, userId string) *Context {
	_AccountId := chi.URLParam(r, "AccountId")
	var _db *client.Database
	if _AccountId == "" {
		_db = client.LoadBaseDatabase()
	} else {
		_db = client.LoadDatabase(_AccountId)
	}
	redisPrefix := _AccountId
	if userId != "" {
		redisPrefix += "::" + userId
	}
	return &Context{
		AccountId:	_AccountId,
		DB: 		_db,
		BaseDB: 	client.LoadBaseDatabase(),
		Cache: 		client.NewRedisClient(redisPrefix),
	}
}