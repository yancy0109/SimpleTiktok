package controller

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sora-blue/SimpleTiktok/repository"
)

type UserRegisterResponse struct{
	Status_code int
	Status_msg string
	User_id int
	Token string
}

type UserLoginResponse struct{
	Status_code int
	Status_msg string
	User_id int
	Token string
}

func InitUser() {
	rand.Seed(time.Now().UnixNano())
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
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 2,
			Status_msg: "Register failed: token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	duration, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("failed to create token duration, err: %e\n", err)
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 3,
			Status_msg: "Register failed: time duration issue",
			User_id: 0,
			Token: "",
		})
	}
	token, err := maker.CreateToken(strconv.Itoa(mainid), duration)
	if err != nil{
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 2,
			Status_msg: "Register failed: token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	context.JSON(http.StatusOK, UserRegisterResponse{
		Status_code: 0,
		Status_msg: "Register success",
		User_id: mainid,
		Token: token,
	})
}

func Login(context *gin.Context){
	username := context.Query("username")
	password := context.Query("password")

	user, err := repository.FindUser(username)

	if err != nil {
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
		context.JSON(http.StatusOK, UserLoginResponse{
			Status_code: 2,
			Status_msg: "Login failed: token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}
	duration, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("failed to create token duration, err: %e\n", err)
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 3,
			Status_msg: "Login failed: time duration issue",
			User_id: 0,
			Token: "",
		})
	}
	token, err := maker.CreateToken(strconv.Itoa(int(user.ID)), duration)
	if err != nil{
		context.JSON(http.StatusOK, UserRegisterResponse{
			Status_code: 2,
			Status_msg: "Register failed: token gen issue",
			User_id: 0,
			Token: "",
		})
		return
	}

	context.JSON(http.StatusOK, UserLoginResponse{
		Status_code: 0,
		Status_msg: "Login success",
		User_id: int(user.ID),
		Token: token,
	})
}
