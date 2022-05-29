package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/controller"
)

func InitRouter(r *gin.Engine) {
	//设置静态目录为public文件夹
	r.Static("/static", "./public")

	//视频发布
	publish := r.Group("/douyin/publish")
	publish.POST("/action", controller.Publish)
}
