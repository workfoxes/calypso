package error


func InValidCredential() BaseError {
	var error BaseError
	error = &Response{ErrorCode: "ERR001001"}
	return error
}