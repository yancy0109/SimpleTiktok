package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/repository"
	"os"
)

func main() {
	//初始化数据库
	if err := repository.Init(); err != nil {
		os.Exit(-1)
	}
	fmt.Println("数据库初始化完成")
	r := gin.Default()
	//路由初始化
	InitRouter(r)
	fmt.Println("路由初始化完成")
	r.Run()
	fmt.Println("路由启动")
}
