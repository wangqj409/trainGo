package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wangqj409/trainGo/week04/blog/internal/model"
	"net/http"
	"strconv"
)

// for list data
func Bcategory(c *gin.Context) {
	bcat := &model.Bcategory{}
	c.JSON(http.StatusOK, gin.H{
		"data": bcat.List(),
	})
}

// for single data
func BScategory(c *gin.Context) {
	bc := model.Bcategory{}
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "required id",
		})
		return
	}
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": bc.FindOne(id),
	})
}

// for dealing post data
func BPcategory(c *gin.Context) {
	cat := &model.Bcategory{}
	cat.CatName = c.PostForm("cat_name")
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	cat.Pid = uint32(pid)
	if cat.CatName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "cat_name should not be empty",
		})
		return
	}
	cat.Create()
	c.JSON(http.StatusOK, gin.H{
		"data": cat,
	})
}
