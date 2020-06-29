FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

EXPOSE 8080
