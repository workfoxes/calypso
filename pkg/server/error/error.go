package error

type BaseError interface {
	error
	Json() *Response
	SetRequestId(requestId string)
}

func ParseError(err error) BaseError {
	return New(err.Error())
}

type Response struct {
	RequestId    string
	ErrorCode    string
	ErrorMessage string
	LangMessage  string
}

func (err *Response) Error() string {
	return err.ErrorMessage
}

func (err *Response) Json() *Response {
	return err
}

func (err *Response) SetRequestId(requestId string) {
	err.RequestId = requestId
}
