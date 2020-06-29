FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d ./...

EXPOSE 8080
