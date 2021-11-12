FROM golang:1.17.3
WORKDIR /go/src/github.com/workfoxes/calypso

COPY . .
RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...