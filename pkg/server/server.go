package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// ApplicationServer : Application server will hold the service object for the application
type ApplicationServer struct {
	Server *fiber.App
	Name   string
	Port   int
}

// CreateAppServer : func to create Application server object to Manage the application server
func CreateAppServer(Name string, Port int) *ApplicationServer {
	app := fiber.New()
	_server := &ApplicationServer{Name: Name, Port: Port, Server: app}
	return _server
}

// LoadDefaultMiddleware : this function will load all the middleware that are need for application
func (app *ApplicationServer) LoadDefaultMiddleware() {
	app.Use(logger.New())
	app.Use(limiter.New())
	app.Use(etag.New())
	app.Use(csrf.New())
	app.Use(pprof.New())
	app.Use(requestid.New())
	app.Use(compress.New(compress.Config{Level: compress.LevelBestCompression}))

}

// Use : This function will allow us to add the middleware into the web application
func (app *ApplicationServer) Use(args ...interface{}) {
	app.Server.Use(args...)
}

// Start : Will Start the Application service for the Thermite
func (app *ApplicationServer) Start() {
	_port := strconv.Itoa(app.Port)
	log.Fatal(app.Server.Listen(":" + _port))
}
