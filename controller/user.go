package controller

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sora-blue/SimpleTiktok"
	"github.com/sora-blue/SimpleTiktok/repository"
)

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

func Register(context *gin.Context){
	username := context.Query("username")
	password := context.Query("password")
	salt := strconv.FormatInt(salt_gen(username), 10)
	userid := username + strconv.Itoa(time.Now().Nanosecond())

	repository.create_user(userid, username, password, salt)

	repository.create_user()



}
