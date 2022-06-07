package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
)

type PublishResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Publish(context *gin.Context) {
	//解析context中 data token title
	title, exist := context.GetPostForm("title")
	if !exist {
		context.JSON(http.StatusOK, PublishResponse{
			StatusCode: 1,
			StatusMsg:  "缺少title",
		})
		return
	}
	token, exist := context.GetPostForm("token")
	if !exist {
		context.JSON(http.StatusOK, PublishResponse{
			StatusCode: 1,
			StatusMsg:  "缺少token",
		})
		return
	}
	//检验token 获取用户Id
	var authorId int64
	var err error
	authorId, err = middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, PublishResponse{
			StatusCode: 1,
			StatusMsg:  "token无效",
		})
		return
	}
	data, err := context.FormFile("data")
	//接收文件失败返回
	if err != nil {
		context.JSON(http.StatusOK, PublishResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//调取service层储存
	if err := service.PublishSave(authorId, title, data, context); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusOK, PublishResponse{
			StatusCode: 1,
			StatusMsg:  "保存失败",
		})
		return
	}
	context.JSON(http.StatusOK, PublishResponse{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}
