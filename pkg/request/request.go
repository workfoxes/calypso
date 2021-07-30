package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Post(url string, body interface{}) (resp *http.Response, err error) {
	postBody, _ := json.Marshal(body)
	requestBody := bytes.NewBuffer(postBody)
	return http.Post(url, fiber.MIMEApplicationJSON, requestBody)
}

func Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
