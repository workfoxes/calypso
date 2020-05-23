package utils

import (
	error2 "github.com/airstrik/gobase/pkg/error"
	"github.com/airstrik/gobase/pkg/schema/system"
)

func GenerateSuccessResponse() interface{}  {
	return &system.SuccessResponse{ Status: "Success" }
}

func HandleError(err error)  {
	if err!=nil{
		panic(err)
	}
}

func HandleCustomError(err error, customerErr func() error2.BaseError)  {
	if err!=nil{
		panic(customerErr())
	}
}
