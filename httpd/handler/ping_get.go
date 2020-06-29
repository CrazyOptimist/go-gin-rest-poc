package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{
			"hello": "found me",
		})
	}
}
