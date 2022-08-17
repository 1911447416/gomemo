package main

import (
	"fmt"
	"gomemo/models"
	"gomemo/routers"
)

func main() {
	// 读取配置
	AppConfig := models.LoadConfig()
	fmt.Println(AppConfig)

	// 注册路由
	routers.SetupRouters(AppConfig)

}
