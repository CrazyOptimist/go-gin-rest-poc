package handler

import (
	"net/http"

	"github.com/crazyoptimist/go-gin-rest-poc/services/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(context *gin.Context) {
		results := feed.GetAll()
		context.JSON(http.StatusOK, results)
	}
}
