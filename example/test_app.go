package main

import "github.com/workfoxes/gobase"

func main() {
	app := gobase.New(&gobase.ApplicationConfig{Name: "TestingApplication", Port: 8080})
	app.Start()
}
