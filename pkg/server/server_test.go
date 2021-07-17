package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test001CreateNewApplicationServer(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, 👋!")
		return c.SendString(msg) // => Hello john 👋!
	})
	app.Listen(":9000")
}

func TestCreateAppServer(t *testing.T) {
	type args struct {
		Name string
		Port int
	}
	tests := []struct {
		name string
		args args
		want *ApplicationServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAppServer(tt.args.Name, tt.args.Port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAppServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
