package main

import (
	"godemo/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	routers.IndexRoutersInit(r)
	r.Run(":8080")

}
