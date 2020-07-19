package utils

import (
	"github.com/workfoxes/gobase/pkg/schema/system"
	baseError "github.com/workfoxes/gobase/pkg/server/error"
	"strings"
)

func GenerateSuccessResponse() interface{} {
	return &system.SuccessResponse{Status: "Success"}
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleCustomError(err error, customerErr func() baseError.BaseError) {
	if err != nil {
		panic(customerErr())
	}
}

func EmailDomainExtractor(email string) (string, string) {
	result := strings.Split(string(email), "@")
	domain := strings.Split(string(result[1]), ".")
	return result[0], domain[0]
}
