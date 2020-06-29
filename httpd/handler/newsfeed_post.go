package handler

import (
	"fmt"
	"net/http"

	"github.com/crazyoptimist/go-gin-rest-poc/services/newsfeed"
	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := newsfeedPostRequest{}
		bindError := context.Bind(&requestBody)
		if bindError != nil {
			fmt.Println(bindError)
		}

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		context.Status(http.StatusNoContent)
	}
}
