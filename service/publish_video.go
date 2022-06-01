package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/repository"
	"mime/multipart"
	"path/filepath"
)

func PublishSave(authorId int64, title string, data *multipart.FileHeader, context *gin.Context) error {
	return NewPublishVideo(authorId, title, data, context).Do()
}

func NewPublishVideo(authorId int64, title string, data *multipart.FileHeader, context *gin.Context) *PublishVideoFlow {
	return &PublishVideoFlow{
		AuthorId: authorId,
		Title:    title,
		Data:     data,
		Context:  context,
	}
}

type PublishVideoFlow struct {
	AuthorId      int64  `json:"author_id"`
	Title         string `json:"title"`
	Data          *multipart.FileHeader
	SaveVedioPath string
	SaveImagePath string
	Context       *gin.Context
}

func (f *PublishVideoFlow) Do() error {
	if err := f.CheckTitle(); err != nil {
		return err
	}
	var finalPath string
	var err error
	if finalPath, err = f.SaveVideo(); err != nil {
		return err
	}
	if err := f.SaveImage(finalPath); err != nil {
		return err
	}
	if err := f.Publish(); err != nil {
		return err
	}
	return nil
}
func (f *PublishVideoFlow) CheckTitle() error {
	if len(f.Title) > 50 {
		return errors.New("Title Too Long")
	}
	return nil
}
func (f *PublishVideoFlow) SaveVideo() (string, error) {
	//对文件进行处理
	//1.获取文件名字
	filename := f.Data.Filename
	//2.生成文件保存名
	finalname := fmt.Sprint(f.AuthorId) + "_" + filename
	//3.生成保存路径
	saveVedioPath := filepath.Join("./public/video/", finalname)
	f.SaveVedioPath = saveVedioPath
	//开始保存
	if err := f.Context.SaveUploadedFile(f.Data, saveVedioPath); err != nil {
		return "", err
	}
	return finalname, nil
}

func (f *PublishVideoFlow) SaveImage(finalname string) error {
	//videoPath := "./public/vedio/" + finalname
	//保存视频流首帧作为封面
	saveCovel := filepath.Join("./public/cover/", finalname+".jpg")
	f.SaveImagePath = saveCovel
	//cmd := exec.Command("ffmpeg", "-i", "./"+videoPath, "-r", "1", "-s", "600x400", "-f", "singlejpeg", "-frames:v", "1", saveCovel)
	//var stdout, stderr bytes.Buffer
	//cmd.Stderr = &stderr
	//cmd.Stdout = &stdout
	//err = cmd.Run()
	//fmt.Printf("out:n%sn err:n%sn", string(stdout.Bytes()), string(stderr.Bytes()))
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %sn", err)
	//	context.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  "保存失败",
	//	})
	//	return
	//}
	return nil
}
func (f *PublishVideoFlow) Publish() error {
	vedio := &repository.Video{
		AuthorId: f.AuthorId,
		Title:    f.Title,
		PlayUrl:  f.SaveVedioPath,
		CoverUrl: f.SaveImagePath,
	}
	if err := repository.NewVedioDaoInstance().PublishVideo(vedio); err != nil {
		return err
	}
	return nil
}