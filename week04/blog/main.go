package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK)
	})

	route.Run()
}
