package main

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("夏天夏天悄悄过去")

// GenToken 生成JWT
func GenToken(username,password string) (string, error) {
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer: "mt-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	token,err := jwt.ParseWithClaims(tokenString,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret,nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims,nil
	}
	return nil, errors.New("invalid token")
}

func main() {
	r := gin.Default()
	
	r.POST("/auth",authHandler)
	
	r.GET("/home", JWTAuthMiddleware(), homeHandler)

	r.Run("localhost:8080")
}

func homeHandler(c *gin.Context) {
	get, exists := c.Get("username")
	if !exists {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username":get},
	})
}

func authHandler(c *gin.Context) {
	var user MyClaims
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "q1mi" && user.Password == "q1mi123"{
		tokenStr,_:=GenToken(user.Username,user.Password)
		c.JSON(http.StatusOK,gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenStr},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}else{
			claim,err := ParseToken(token)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2005,
					"msg":  "无效的Token",
				})
				c.Abort()
				return
			}else if time.Now().Unix() > claim.ExpiresAt{
				c.JSON(http.StatusOK, gin.H{
					"code": 2006,
					"msg":  "超时",
				})
				c.Abort()
				return
			}
			// 将当前请求的username信息保存到请求的上下文c上
			c.Set("username", claim.Username)
		}
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}