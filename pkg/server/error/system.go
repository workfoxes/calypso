package error

var err BaseError

func New(str string) BaseError {
	err = &Response{ErrorMessage: "Internal Server Error", LangMessage: "Internal Server Error", ErrorCode: "ERR00001"}
	return err
}

func TenantNotFound() BaseError {
	err = &Response{
		ErrorMessage: "Account Id is Invalid Not found",
		LangMessage: "Account Id is Invalid Not found in system",
		ErrorCode: "ERR00002",
	}
	return err
}

func JsonParseFailure() BaseError {
	err = &Response{
		ErrorMessage: "Failed to Accept the Request",
		LangMessage: "Failed to Accept the Request",
		ErrorCode: "ERR00003",
	}
	return err
}