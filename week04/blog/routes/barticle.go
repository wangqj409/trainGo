package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wangqj409/trainGo/week04/blog/internal/model"
	"net/http"
	"strconv"
)

func Barticles(c *gin.Context) {
	articles := &model.Barticle{}
	offset, limit := c.DefaultQuery("offset", "0"), c.DefaultQuery("limit", "10")
	o, _ := strconv.Atoi(offset)
	l, _ := strconv.Atoi(limit)
	c.JSON(http.StatusOK, gin.H{
		"data": articles.List(o, l),
	})
}

func BSarticle(c *gin.Context) {
	id := c.Param("id")
	art := &model.Barticle{}
	c.JSON(http.StatusOK, gin.H{
		"data": art.FindOne(id),
	})

}

func BParticle(c *gin.Context) {
	cat_id := c.PostForm("cat_id")
	if cat_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing cat_id param",
		})
		return
	}

	art := &model.Barticle{}
	art.Title = c.PostForm("title")
	if len(art.Title) > 10 {
		art.SubTitle = art.Title[:10]
	} else {
		art.SubTitle = art.Title
	}
	art.Content = c.PostForm("content")
	cid, _ := strconv.Atoi(cat_id)
	art.CatId = uint(cid)
	art.Create()
	c.JSON(http.StatusOK, gin.H{
		"data": art,
	})

}
