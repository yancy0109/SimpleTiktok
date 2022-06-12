package controller

import (
	"crypto/sha256"
	"fmt"
	"github.com/yancy0109/SimpleTiktok/service"
	"net/http"
	"strconv"
	"time"

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

type User_rep struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Follow_count   int64  `json:"follow_count"`
	Follower_count int64  `json:"follower_count"`
	Is_follow      bool   `json:"is_follow"`
}

type UserInfoResponse struct {
	Status_code int          `json:"status_code"`
	Status_msg  string       `json:"status_msg"`
	User        service.User `json:"user"`
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
	pwarr := sha256.Sum256([]byte(rawPw))

	var hashedPw string
	for _, pwele := range pwarr {
		hashedPw += fmt.Sprintf("%x", pwele)
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
		fmt.Printf("user already exists\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 4,
			Status_msg:  "Register failed: user already exists",
			User_id:     0,
			Token:       "",
		})
		return
	}

	user_id, err := repository.CreateUser(username, password, salt)
	if err != nil {
		fmt.Printf("failed to register, err: %e\n", err)
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 1,
			Status_msg:  "Register failed: repository issue",
			User_id:     0,
			Token:       "",
		})
		return
	}
	token, err := middleware.GenToken(user_id)
	if err != nil {
		fmt.Printf("failed to create Token, err: %e\n", err)
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 2,
			Status_msg:  "Register failed: Token gen issue",
			User_id:     0,
			Token:       "",
		})
		return
	}
	fmt.Printf("success, User_id: %v, Token: %v\n", user_id, token)
	context.JSON(http.StatusOK, UserLoginResponse{
		Status_code: 0,
		Status_msg:  "Register success",
		User_id:     user_id,
		Token:       token,
	})
}

func Login(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	user, err := repository.FindUser(username)

	if err != nil {
		fmt.Printf("failed to find user\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 1,
			Status_msg:  "Login failed: cannot find user",
			User_id:     0,
			Token:       "",
		})
		return
	}
	underVerify := pwHash(password + user.Salt)[0:50]

	if underVerify != user.Password {
		fmt.Printf("password not matched\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 4,
			Status_msg:  "Login failed: password not matched",
			User_id:     0,
			Token:       "",
		})
		return
	}
	token, err := middleware.GenToken(user.ID)
	if err != nil {
		fmt.Printf("failed to create Token\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 2,
			Status_msg:  "Login failed: Token gen issue",
			User_id:     0,
			Token:       "",
		})
		return
	}

	fmt.Printf("success, User_id: %v, Token: %v\n", user.ID, token)
	context.JSON(http.StatusOK, UserLoginResponse{
		Status_code: 0,
		Status_msg:  "Login success",
		User_id:     user.ID,
		Token:       token,
	})
}

func UserInfo(context *gin.Context) {
	var userService service.UserService
	token := context.Query("token")
	userIdStr := context.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	tokenUserId, err := middleware.ParseToken(token)
	if err != nil {
		msg := "token无效"
		context.JSON(http.StatusOK, UserInfoResponse{
			Status_code: 0,
			Status_msg:  msg,
			User: service.User{
				FollowCount:   0,
				FollowerCount: 0,
				ID:            -1,
				IsFollow:      false,
				Name:          "sheep",
			},
		})
		return
	}

	user, err := userService.QueryUserInfo(userId, tokenUserId)
	context.JSON(http.StatusOK, UserInfoResponse{
		Status_code: 0,
		Status_msg:  "OK",
		User:        user,
	})
}
