package api

/*
	来源：https://www.liwenzhou.com/posts/Go/jwt_in_gin/
*/
import (
	"errors"
	"kvm_backup/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	*model.AuthUser
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("念念不忘，必有回响")

// GenToken 生成JWT
func GenToken(u *model.AuthUser) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		u, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "kvm_backup",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func GetUserToken(c *gin.Context, u *model.AuthUser) {
	if v, err := c.Cookie(SID); err == nil {
		if claim, err := ParseToken(v); err == nil {

			u.ID = claim.AuthUser.ID
			u.Name = claim.AuthUser.Name
		}
	}
}

// 通过cookie来存储，客户端也可以通过配置文件存储,header传参
func SetUserToken(c *gin.Context, u *model.AuthUser, s sessions.Options) {
	if v, err := GenToken(u); err == nil {
		c.SetCookie(SID, v, s.MaxAge, s.Path, s.Domain, s.Secure, s.HttpOnly)
	}
}

// cookie清理
func DelUserToken(c *gin.Context, s *sessions.Options) {
	if _, err := c.Cookie(SID); err == nil {
		c.SetCookie(SID, "", -60, s.Path, s.Domain, s.Secure, s.HttpOnly)
	}
}
