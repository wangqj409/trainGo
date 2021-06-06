package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wangqj409/trainGo/week04/blog/internal/model"
	"net/http"
)

// for list data
func Bcategory(c *gin.Context) {
	bcat := &model.Bcategory{}
	c.JSON(http.StatusOK, gin.H{
		"data": bcat.List(),
	})
	return
}

// for single data
func BScategory(c *gin.Context) {
	var bc model.Bcategory
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "required id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bc,
	})
	return
}

// for dealing post data
func BPcategory(c *gin.Context) {
	var cat model.Bcategory
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if cat.CatName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "cat_name should not be empty",
		})
		return
	}
}
