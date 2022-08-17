package controller

import (
	"fmt"
	"gomemo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/index.html", gin.H{
		"success":  true,
		"Memolist": models.GetAllMemo(),
	})
}

func (con IndexController) DoPost(c *gin.Context) {
	content := c.PostForm("content")
	if ok := models.CreateMemo(content); !ok {
		// c.JSON(http.StatusNotAcceptable, gin.H{
		// 	"success": false,
		// 	"msg":     "增加失败",
		// })
		fmt.Println("增加失败")
	} else {
		fmt.Println("增加成功")
	}
}

func (con IndexController) DoDelete(c *gin.Context) {
	if ok := models.DeleteMemo(c.Query("id")); !ok {
		fmt.Println("删除失败")
	}
}

func (con IndexController) DoDone(c *gin.Context) {
	if ok := models.DoneMemo(c.Query("id")); !ok {
		fmt.Println("修改失败")
	}
}

func (con IndexController) NoDone(c *gin.Context) {
	if ok := models.NoDoneMemo(c.Query("id")); !ok {
		fmt.Println("修改失败")
	}
}
