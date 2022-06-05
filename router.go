package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/controller"
)

func InitRouter(r *gin.Engine) {
	//设置静态目录为public文件夹
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	//// basic apis
	//apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/feed", controller.VideoFeed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)
	//
	//// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/")
}
