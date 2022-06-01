package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/service"
	"log"
	"net/http"
	"os/exec"
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
		return
	}
	//对文件进行处理
	//1.获取文件名字
	filename := data.Filename
	//2.生成文件保存名
	finalname := "user11_" + filename
	//3.生成保存路径
	saveFile := filepath.Join("./public/video/", finalname)
	//开始保存
	if err := context.SaveUploadedFile(data, saveFile); err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//save first frame
	saveImage := filepath.Join("./public/image/", finalname+".jpg")
	fmt.Println(saveFile)
	cmd := exec.Command("ffmpeg", "-i", "./"+saveFile, "-r", "1", "-s", "600x400", "-f", "singlejpeg", "-frames:v", "1", saveImage)
	//var stdout, stderr bytes.Buffer
	//cmd.Stderr = &stderr
	//cmd.Stdout = &stdout
	err = cmd.Run()
	//fmt.Printf("out:n%sn err:n%sn", string(stdout.Bytes()), string(stderr.Bytes()))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %sn", err)
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "保存失败",
		})
		return
	}
	//调取service层存储数据至数据库
	if err := service.SavePublish(userid, saveFile, saveImage, title); err != nil {
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
