package controller

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yancy0109/SimpleTiktok/repository"
)

type UserRegisterResponse struct{
	Status_code int `json:"status_code"`
	Status_msg string `json:"status_msg"`
	User_id int `json:"user_id"`
	Token string `json:"token"`
}

type UserLoginResponse struct{
	Status_code int `json:"status_code"`
	Status_msg string `json:"status_msg"`
	User_id int `json:"user_id"`
	Token string `json:"token"`
}

type User_rep struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Follow_count int `json:"follow_count"`
	Follower_count int `json:"follower_count"`
	Is_follow bool `json:"is_follow"` 
}

type UserInfoResponse struct {
	Status_code int `json:"status_code"`
	Status_msg string `json:"status_msg"`
	User []User_rep `json:"user"`
}

func salt_gen(username string) int64{
	var hx int64 = 0
	for _, ch := range username {
		hx = (hx << 8) ^ int64(ch);
	}
	hx = hx ^ time.Now().UnixNano()
	return hx
}

func pwHash(rawPw string) string {
	pwarr := sha256.Sum256([]byte(rawPw))

	var hashedPw string;
	for _, pwele := range pwarr {
		hashedPw += fmt.Sprintf("%x", pwele);
	}
	return hashedPw
}

func Register(context *gin.Context){
	username := context.Query("username")
	password := context.Query("password")

	salt := strconv.FormatInt(salt_gen(username), 10)
	password = pwHash(password+salt)[0:50];

//	userid_int := time.Now().UnixNano()
//	userid := strconv.FormatInt(userid_int, 10)

	mainid, err := repository.CreateUser(username, password, salt)
	if err != nil {
		fmt.Printf("failed to register, err: %e\n", err)
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 1,
			Status_msg: "Register failed: repository issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	maker := GetGlobalMaker()
	if maker == nil {
		fmt.Printf("failed to get global maker\n")
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 2,
			Status_msg: "Register failed: Token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	duration, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("failed to create Token duration, err: %e\n", err)
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 3,
			Status_msg: "Register failed: time duration issue",
			User_id: 0,
			Token: "",
		})
	}
	Token, err := maker.CreateToken(strconv.Itoa(mainid), duration)
	if err != nil{
		fmt.Printf("failed to create Token, err: %e\n", err)
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 2,
			Status_msg: "Register failed: Token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	fmt.Printf("success, User_id: %v, Token: %v\n", mainid, Token)
	context.JSON(http.StatusOK, UserRegisterResponse{
		Status_code: 0,
		Status_msg: "Register success",
		User_id: mainid,
		Token: Token,
	})
}

func Login(context *gin.Context){
	username := context.Query("username")
	password := context.Query("password")

	user, err := repository.FindUser(username)

	if err != nil {
		fmt.Printf("failed to find user\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 1,
			Status_msg: "Login failed: cannot find user",
			User_id: 0,
			Token: "",
		})
		return
	}
	underVerify := pwHash(password + user.Salt)[0:50]

	if underVerify != user.Password {
		fmt.Printf("password not matched\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 4,
			Status_msg: "Login failed: password not matched",
			User_id: 0,
			Token: "",
		})
		return
	}
	maker := GetGlobalMaker()
	if maker == nil {
		fmt.Printf("failed to get maker\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 2,
			Status_msg: "Login failed: Token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	duration, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("failed to create Token duration, err: %e\n", err)
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 3,
			Status_msg: "Login failed: time duration issue",
			User_id: 0,
			Token: "",
		})
	}
	Token, err := maker.CreateToken(strconv.Itoa(int(user.ID)), duration)
	if err != nil{
		fmt.Printf("failed to create Token\n")
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 2,
			Status_msg: "Login failed: Token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}

	fmt.Printf("success, User_id: %v, Token: %v\n", user.ID, Token)
	context.JSON(http.StatusOK, UserLoginResponse{
		Status_code: 0,
		Status_msg: "Login success",
		User_id: int(user.ID),
		Token: Token,
	})
}

func UserInfo(context *gin.Context){
	users := make([]User_rep, 2)
	users[0] = User_rep{
		Id: 0,
		Name: "Sheep Sherry",
		Follow_count: 0,
		Follower_count: 0,
		Is_follow: false,
	}
	context.JSON(http.StatusOK, UserInfoResponse{
		Status_code: 0,
		Status_msg: "OK",
		User: users,
	})
}
