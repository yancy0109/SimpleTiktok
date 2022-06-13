package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	StatusCode  int32                 `json:"status_code"`
	StatusMsg   string                `json:"status_msg"`
	CommentList []service.CommentInfo `json:"comment_list"`
}

func CommentList(context *gin.Context) {
	var token string
	var videoId string
	var userId int64
	var exist bool
	var err error
	if token, exist = context.GetQuery("token"); !exist {
		context.JSON(http.StatusOK, CommentListResponse{
			StatusCode: -1,
			StatusMsg:  "缺少token",
		})
		return
	}
	userId, err = middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, PublishResponse{
			StatusCode: -1,
			StatusMsg:  "token无效",
		})
		return
	}
	if videoId, exist = context.GetQuery("video_id"); !exist {
		context.JSON(http.StatusOK, CommentListResponse{
			StatusCode: -1,
			StatusMsg:  "缺少video_id",
		})
		return
	}
	var result_videoId int64
	if result_videoId, err = strconv.ParseInt(videoId, 10, 64); err != nil {
		context.JSON(http.StatusOK, CommentListResponse{
			StatusCode: -1,
			StatusMsg:  "解析video_id错误",
		})
		return
	}
	var commentList []service.CommentInfo
	if commentList, err = service.GetCommentList(result_videoId, userId); err != nil {
		context.JSON(http.StatusOK, CommentListResponse{
			StatusCode: -1,
			StatusMsg:  "获取评论列表错误",
		})
		return
	}
	context.JSON(http.StatusOK, CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取评论列表成功",
		CommentList: commentList,
	})
}
