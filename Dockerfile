FROM golang:1.14

WORKDIR $GOPATH/src/github.com/workfoxes/gobase
COPY . .
RUN go get -d -v ./...
