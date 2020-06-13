package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/go-redis/redis"
	"os"
)

var (
	RedisConnection *redis.Client
	SnsClient       *sns.SNS
	dialect         string
	host            string
	port            string
	user            string
	password        string
	err             error
	redisHost       string
	BaseDB          *Database
)

func init() {
	// Check and Load the Client Connection
	dialect = os.Getenv("_DB_DIALECT")
	host = os.Getenv("_DB_HOST")
	port = os.Getenv("_DB_PORT")
	user = os.Getenv("_DB_USER")
	password = os.Getenv("_DB_PASSWORD")
	redisHost = os.Getenv("_REDIS_URI")
	CreateRedisConnection()
}

func CreateRedisConnection() {
	if redisHost != "" {
		RedisConnection = redis.NewClient(&redis.Options{
			Addr: redisHost,
			DB:   0, // use default DB
		})
	}
}

func CreateAwsSession() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	SnsClient = sns.New(sess)
}
