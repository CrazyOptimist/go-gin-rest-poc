package main

import (
	"fmt"

	"github.com/crazyoptimist/go-gin-rest-poc/httpd/handler"
	"github.com/crazyoptimist/go-gin-rest-poc/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	server := gin.Default()

	server.GET("/ping", handler.PingGet())
	server.GET("/newsfeed", handler.NewsfeedGet(feed))
	server.POST("/newsfeed", handler.NewsfeedPost(feed))

	error := server.Run()
	if error != nil {
		fmt.Println(error)
	}

}
