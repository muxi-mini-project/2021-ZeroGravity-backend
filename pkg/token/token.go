package token

import(
	
	"errors"

	"fmt"
	"log"


	"github.com/dgrijalva/jwt-go"
	"github.com/2021-ZeroGravity-backend/model"
)

JwtCliams...
type JwtClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

var Secret = "sault" //加盐

func GetToken(claims *model.JwtClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
		return ""
	}
	return signedToken
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

func VerifyToken(strToken string) (*JwtClaims, error) {
	//	strToken := c.Param("token")    //c.Param()可以获取单个参数,路径的也行

	//解析
	token, err := jwt.ParseWithClaims(strToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}

	return claims, nil //JwtClaims结构体类型
}