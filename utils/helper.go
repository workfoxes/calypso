package utils

import "strings"

func GenerateSuccessResponse() interface{} {
	return SuccessResponse{Status: "Success"}
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func EmailDomainExtractor(email string) (string, string) {
	result := strings.Split(email, "@")
	domain := strings.Split(result[1], ".")
	return result[0], domain[0]
}
