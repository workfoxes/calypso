package utils

type ResponseObject struct {
	Id int
	Status string
	Message string
	ErrorCode string
	ArgMessage string
	RequestId string
}

type SuccessResponse struct {Status string}
