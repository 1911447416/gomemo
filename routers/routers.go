package routers

import (
	"gomemo/models"

	"github.com/gin-gonic/gin"
)

func SetupRouters(AppConfig *models.AppConfig) {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	IndexRoutersInit(r)
	r.Run(":" + AppConfig.Port)
}
