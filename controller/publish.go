package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"path/filepath"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Publish(context *gin.Context) {
	//解析context中 data token title
	title := context.PostForm("title")
	token := context.PostForm("token")
	//检验token获取user
	var userid int
	if token != "" {
		userid = 2123123
	}
	data, err := context.FormFile("data")
	//接收文件失败返回
	if err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	//对文件进行处理
	//1.获取文件名字
	filename := data.Filename
	//2.生成文件保存名
	finalname := "user11_" + filename
	//3.生成保存路径
	saveFile := filepath.Join("./public/", finalname)
	//开始保存
	if err := context.SaveUploadedFile(data, finalname); err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	//调取service层存储数据至数据库
	if err := service.SavePublish(userid, saveFile, "", title); err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "保存失败",
		})
	}
	context.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  filename + "uploaded successfully",
	})
}
