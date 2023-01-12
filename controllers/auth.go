package controllers

import (
	"errors"
	"house_system_backend/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	User model.UserInfo
	jwt.StandardClaims
}

const (
	TokenExpireDuration = time.Hour * 2
	M                   = time.Minute * 5
)

var MySecret = []byte("hssecret") // 生成签名的密钥

func jWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			// 无token直接拒绝
			c.Abort()
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 401,
				Msg:  "无登陆权限",
			})
			return
		}
		// 校验token
		claims, err := ParseToken(auth)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				// 若过期调用续签函数
				newToken, _ := RenewToken(claims)
				if newToken != "" {
					// 续签成功給返回头设置一个newtoken字段
					c.Header("newtoken", newToken)
					c.Request.Header.Set("Authorization", newToken)
					c.Next()
					return
				}
			}
			// Token验证失败或续签失败直接拒绝请求
			c.Abort()
			c.JSON(http.StatusBadRequest, model.ResponseData{
				Code: 401,
				Msg:  err.Error(),
			})
			return
		}
		// token未过期继续执行1其他中间件
		c.Next()
	}
}

func GenerateToken(userInfo model.UserInfo) (string, error) {
	expirationTime := time.Now().Add(M) // 两个小时有效期
	claims := &MyClaims{
		User: userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "yourname",
		},
	}
	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	if tokenString, err := token.SignedString(MySecret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}

}

func RenewToken(claims *MyClaims) (string, error) {
	// 若token过期不超过10分钟则给它续签
	if withinLimit(claims.ExpiresAt, 600) {
		return GenerateToken(claims.User)
	}
	return "", errors.New("登录已过期")
}

func ParseToken(tokenString string) (*MyClaims, error) {
	claims := &MyClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	// 若token只是过期claims是有数据的，若token无法解析claims无数据
	return claims, err
}

// 计算过期时间是否超过l
func withinLimit(s int64, l int64) bool {
	e := time.Now().Unix()
	// println(e - s)
	return e-s < l
}

func GetAuth(r *gin.Engine) {
	// jwt 已经封装成了中间件，可以放在任意请求中，但是由于本系统已设置了Traefik网关，鉴权程序不需要每次都走中间件后端程序，方便以后扩展业务，只需要走网关就可以实现token的鉴权以及转发
	// 如果token错误，在这里就会直接返回错误
	authorized := r.Group("/", jWTAuth())
	authorized.GET("/auth", func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		// 这里会触发解析token的解析算法
		claims, _ := ParseToken(auth)
		c.JSON(http.StatusOK, model.ResponseData{
			Code: 200,
			Msg:  "hello " + claims.User.UserName,
		})
	})
}
