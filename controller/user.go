package controller

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yancy0109/SimpleTiktok/service"

	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/middleware"
	"github.com/yancy0109/SimpleTiktok/repository"
)

type UserLoginResponse struct {
	Status_code int    `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User_id     int64  `json:"user_id"`
	Token       string `json:"token"`
}

type UserInfoResponse struct {
	StatusCode int               `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	User       repository.Author `json:"user"`
}

func salt_gen(username string) int64 {
	var hx int64 = 0
	for _, ch := range username {
		hx = (hx << 8) ^ int64(ch)
	}
	hx = hx ^ time.Now().UnixNano()
	return hx
}

func pwHash(rawPw string) string {
	pwArr := sha256.Sum256([]byte(rawPw))

	var hashedPw string
	for _, pwEle := range pwArr {
		hashedPw += fmt.Sprintf("%x", pwEle)
	}
	return hashedPw
}

func Register(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	salt := strconv.FormatInt(salt_gen(username), 10)
	password = pwHash(password + salt)[0:50]

	_, err := repository.FindUser(username)

	if err == nil {
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: -1,
			Status_msg:  "用户名已经存在",
			User_id:     0,
			Token:       "",
		})
		return
	}

	user_id, err := repository.CreateUser(username, password, salt)
	if err != nil {
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: -1,
			Status_msg:  "创建用户失败",
			User_id:     0,
			Token:       "",
		})
		return
	}
	token, err := middleware.GenToken(user_id)
	if err != nil {
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: -1,
			Status_msg:  "token生成失败",
			User_id:     0,
			Token:       "",
		})
		return
	}
	context.JSON(http.StatusOK, UserLoginResponse{
		Status_code: 0,
		Status_msg:  "注册成功",
		User_id:     user_id,
		Token:       token,
	})
}

func Login(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	user, err := repository.FindUser(username)

	if err != nil {
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: -1,
			Status_msg:  "用户名无效",
			User_id:     0,
			Token:       "",
		})
		return
	}
	underVerify := pwHash(password + user.Salt)[0:50]

	if underVerify != user.Password {
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: -1,
			Status_msg:  "密码错误",
			User_id:     0,
			Token:       "",
		})
		return
	}
	token, err := middleware.GenToken(user.ID)
	if err != nil {
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: -1,
			Status_msg:  "token生成失败",
			User_id:     0,
			Token:       "",
		})
		return
	}
	context.JSON(http.StatusOK, UserLoginResponse{
		Status_code: 0,
		Status_msg:  "登录成功",
		User_id:     user.ID,
		Token:       token,
	})
}

func UserInfo(context *gin.Context) {
	var token string
	var userId int64
	var exist bool
	var err error
	if token, exist = context.GetQuery("token"); !exist {
		context.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "缺少token",
		})
		return
	}
	if userId, err = middleware.ParseToken(token); err != nil {
		context.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "token无效",
		})
		return
	}
	var user_id int64
	if user_id, err = strconv.ParseInt(context.Query("user_id"), 10, 64); err != nil {
		context.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "user_id解析错误",
		})
		return
	}

	if userId != user_id {
		context.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "token信息与user_id不符",
		})
		return
	}
	var usereInfo repository.Author
	if usereInfo, err = service.GetUserInfo(userId); err != nil {
		context.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败",
		})
		return
	}

	context.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取",
		User:       usereInfo,
	})
	return
}
