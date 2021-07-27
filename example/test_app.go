package main

import "github.com/itsparser/gobase"

func main() {
	app := gobase.New(&gobase.ApplicationConfig{Name: "TestingApplication", Port: 8080})
	app.Start()
}
