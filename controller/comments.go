package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/repository"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"time"
)

type CommentResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	Comment    Comment `json:"comment"`
}
type Comment struct {
	Content    string            `json:"content"`     // 评论内容
	CreateDate string            `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64             `json:"id"`          // 评论id
	User       repository.Author `json:"user"`        // 评论用户信息
}

//type User_rep struct {
//	FollowCount   int64  `json:"follow_count"`   // 关注总数
//	FollowerCount int64  `json:"follower_count"` // 粉丝总数
//	ID            int64  `json:"id"`             // 用户id
//	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
//	Name          string `json:"name"`           // 用户名称
//}

func CommentAction(context *gin.Context) {
	token := context.Query("token")
	//检查token
	userid, err := middleware.ParseToken(token)
	if err != nil {
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token无效",
		})
		return
	}
	//var userid int64
	//userid = 1001
	videoid := context.Query("video_id")
	actionType := context.Query("action_type")
	comment_text := context.Query("comment_text")
	var commentService service.CommentService
	if actionType == "1" {
		comment_Id, err := commentService.INSERT(userid, videoid, actionType, comment_text)
		users, err := repository.NewVideoListInstance().AuthorInformation(userid, userid)
		if err != nil {
			context.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error()})
		} else {
			context.JSON(http.StatusOK, CommentResponse{
				StatusCode: 0,
				StatusMsg:  "successfully",
				Comment: Comment{
					ID: comment_Id,
					User: repository.Author{
						Id:            userid,
						UserName:      users.UserName,
						FollowCount:   users.FollowerCount,
						FollowerCount: users.FollowerCount,
						IsFollow:      users.IsFollow,
					},
					Content:    comment_text,
					CreateDate: time.Now().Format("01/02"),
				}})
		}
	} else if actionType == "2" {
		comment_id := context.Query("comment_id")
		err := commentService.UPDATE(userid, videoid, actionType, comment_id)
		if err != nil {
			context.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error()})
		} else {
			context.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "successfully",
			})
		}
	}

}
