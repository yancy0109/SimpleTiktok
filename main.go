package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/repository"
	"os"
)

func main() {
	//初始化数据库
	if err := repository.Init(); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()
	//路由初始化
	InitRouter(r)
	r.Run()
}
