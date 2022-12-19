package jwt_use

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 通过鉴权账户,并生成对应的token进行返回
func authHandler(c *gin.Context) {

	user := &UserInfo{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(200, gin.H{"code": 2001, "msg": "invalid params"})
		return
	}

	// 检查人员是否存在,并为其生成一个token,单点登陆友好选择jwt
	if user.UserName == "lx" && user.PassWord == "123qweasd" {
		tokenString, _ := GetToken(user.UserName)
		c.JSON(200, gin.H{"code": 0, "msg": "success", "data": gin.H{"token": tokenString}})
		return
	}
	c.JSON(200, gin.H{"code": 2002, "msg": "鉴权失败"})
	return
}

// 中间件,认证token合法性
func jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.Request.Header.Get("authorization")
		if authHandler == "" {
			c.JSON(200, gin.H{"code": 2003, "msg": "请求头部auth为空"})
			c.Abort()
			return
		}

		// 前两部门可以直接解析出来
		jwt := strings.Split(authHandler, ".")
		cnt := 0
		for _, val := range jwt {
			cnt++
			if cnt == 3 {
				break
			}
			msg, _ := base64.StdEncoding.DecodeString(val)
			fmt.Println("val ->", string(msg))
		}

		// 我们使用之前定义好的解析JWT的函数来解析它,并且在内部解析时判断了token是否过期
		mc, err := ParseToken(authHandler)
		if err != nil {
			fmt.Println("err = ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.UserName)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

//
func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}

// TSSHandler ssh加密验证
func TSSHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("start ->")
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			fmt.Println("err = ", err.Error())
			return
		}
		fmt.Println("+ err = ", err)
		c.Next()
		fmt.Println("success")
	}
}
func main() {
	r := gin.Default()
	// r.Use(TSSHandler(8080))
	r.POST("/auth", authHandler)
	r.GET("/home", jwtAuthMiddleware(), homeHandler)
	fmt.Println("ser is running")
	r.Run(":18243") // http
	// r.RunTLS(":8080","cert.pem", "key.pem") // https
}
