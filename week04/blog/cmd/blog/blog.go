package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/wangqj409/trainGo/week04/blog/routes"
)

func main() {
	gin.ForceConsoleColor()
	route := gin.Default()
	v1 := route.Group("/v1")
	{
		v1.GET("/category", routes.Bcategory)
		v1.GET("/category/:id", routes.BScategory)
		v1.POST("/category", routes.BPcategory)

		v1.GET("/article", routes.Barticles)
		v1.GET("/article/:id", routes.BSarticle)
		v1.POST("/article", routes.BParticle)

	}
	route.Run()
}
