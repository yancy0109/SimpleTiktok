package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JWTResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserClaim struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

//过期时间
const TokenExpireDuration = time.Hour * 2

//密钥
var Secret = []byte("simpletiktok")

//生成token
func GenToken(userId int64) (string, error) {
	userClaim := UserClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "admin",                                    //签发人
		},
	}
	//使用指定签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	//使用指定Secret签名并获得完整编码后的字符串token
	return token.SignedString(Secret)
}

//token校验
func ParseToken(tokenString string) (int64, error) {
	var token, err = jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return -1, err
	}
	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims.UserId, nil
	}
	return -1, errors.New("invalid token")
}

//token校验中间件...随便扔个先 如果要用再修改吧
func JWTAuthMiddleware(context *gin.Context) {
	token, ok := context.GetQuery("token")
	if !ok {
		context.JSON(http.StatusOK, JWTResponse{
			StatusCode: -1,
			StatusMsg:  "未携带token",
		})
	}
	userId, err := ParseToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, JWTResponse{
			StatusCode: -1,
			StatusMsg:  "无效token",
		})
	}
	context.Set("userid", userId)
	//不调用next似乎也会继续方法链路 未测试
	context.Next()
}
