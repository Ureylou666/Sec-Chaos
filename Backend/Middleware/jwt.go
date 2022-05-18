package Middleware

import (
	"Backend/Utils"
	"Backend/Utils/ErrMsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(Utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string, role_id int) (string, int) {
	expireTime := time.Now().Add(3 * time.Hour)
	SetClaims := MyClaims{
		username,
		role_id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Sec_Chaos",
		},
	}

	reqToken := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqToken.SignedString(JwtKey)
	if err != nil {
		return "", ErrMsg.ERROR
	} else {
		return token, ErrMsg.SUCCESS
	}
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if claims, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return claims, ErrMsg.SUCCESS
	} else {
		return nil, ErrMsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	var code int
	var key *MyClaims
	return func(c *gin.Context) {
		// 判断请求是否带有jwt
		tokenHeader := c.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			code = ErrMsg.ERROR_JWT_MISSING
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    code,
				"message": ErrMsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// 判断jwt格式
		reqToken := strings.SplitN(tokenHeader, " ", 2)
		if len(reqToken) != 2 && reqToken[0] != "Bearer" {
			code = ErrMsg.ERROR_JWT_TYPEWRONG
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    code,
				"message": ErrMsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// 验证jwt
		key, code = CheckToken(reqToken[1])
		// 验证token是否Sec_Chaos签发
		if code == ErrMsg.ERROR {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    ErrMsg.ERROR_JWT_NOTVALID,
				"message": ErrMsg.GetErrMsg(ErrMsg.ERROR_JWT_NOTVALID),
			})
			c.Abort()
			return
		}
		// 验证token有效期
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    ErrMsg.ERROR_JWT_TIMEEXPIRED,
				"message": ErrMsg.GetErrMsg(ErrMsg.ERROR_JWT_TIMEEXPIRED),
			})
			c.Abort()
			return
		}
		// middleware验证JWT通过
		c.Set("username", key.Username)
		c.Next()
	}
}
