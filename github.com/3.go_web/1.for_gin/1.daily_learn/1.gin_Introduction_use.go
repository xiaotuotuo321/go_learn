package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 1.安装及使用
//func main() {
//	// 创建一个默认的路由引擎
//	r := gin.Default()
//	r.GET("/hello", func(c *gin.Context){
//		c.JSON(200, gin.H{
//			"message": "hello world!",
//		})
//	})
//	r.Run()
//}

// 2.restful api
/*
	1.get 获取资源
	2.post 新建资源
	3.put 更新资源
	4.delete 删除资源
*/

//func main() {
//	r := gin.Default()
//	r.GET("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "GET",
//		})
//	})
//
//	r.POST("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "POST",
//		})
//	})
//
//	r.PUT("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "PUT",
//		})
//	})
//
//	r.DELETE("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "DELETE",
//		})
//	})
//}

// 3.Gin 渲染
// 3.1.html渲染
//func main() {
//	r := gin.Default()
//	path, _ := os.Getwd()
//	fmt.Println(path)
//
//	r.GET("/hello", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "hello, world!",
//		})
//	})
//	// TODO: 加载文件可以模糊匹配
//	r.LoadHTMLFiles("github.com/3.go_web/1.for_gin/1.daily_learn/templates/posts/index.html", "github.com/3.go_web/1.for_gin/1.daily_learn/templates/users/index.html")
//
//	r.GET("/posts/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "posts/index.html", gin.H{
//			"title": "posts/index",
//		})
//	})
//
//	r.GET("/users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "users/index.html", gin.H{
//			"title": "users/index",
//		})
//	})
//	r.Run(":8080")
//}

// 3.2.自定义函数模块
//func main() {
//	r := gin.Default()
//	r.SetFuncMap(template.FuncMap{
//		"safe": func(str string) template.HTML{
//			return template.HTML(str)
//		},
//	})
//	r.LoadHTMLFiles("github.com/3.go_web/1.for_gin/1.daily_learn/templates/index.tmpl")
//	r.GET("index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
//	})
//	r.Run(":8080")
//}

// 3.3.静态文件处理
//func main() {
//	r := gin.Default()
//	r.Static("static", "./static")
//	r.LoadHTMLGlob("templates/**/*")
//	// ...
//	r.Run(":8080")
//}

// 3.4.使用模板继承
// gin框架默认都是使用单模板，如果需要使用block template时需要使用第三方的库
/*
templates
├── includes
│   ├── home.tmpl
│   └── index.tmpl
├── layouts
│   └── base.tmpl
└── scripts.tmpl
*/
//func loadTemplates(templatesDir string) multitemplate.Renderer{
//	r := multitemplate.NewRenderer()
//	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
//	if err != nil{
//		panic(err.Error())
//	}
//	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
//	if err != nil{
//		panic(err.Error())
//	}
//
//	// 为了layouts/ 和 includes/ 目录生成 templates map
//	for _, include := range includes{
//		layoutCopy := make([]string, len(layouts))
//		copy(layoutCopy, layouts)
//		files := append(layoutCopy, include)
//		r.AddFromFiles(filepath.Base(include), files...)
//	}
//	return r
//}
//
//func indexFunc(c *gin.Context){
//	c.HTML(http.StatusOK, "index.tmpl", nil)
//}
//
//func homeFunc(c *gin.Context){
//	c.HTML(http.StatusOK, "home.tmpl", nil)
//}
//
//func main() {
//	r := gin.Default()
//	r.HTMLRender = loadTemplates("./templates")
//	r.GET("/index", indexFunc)
//	r.GET("/home", homeFunc)
//	r.Run()
//}

// 3.5.补充文件处理
// 关于模块文件和静态文件的路径，
//func getCurrentPath() string{
//	if ex, err := os.Executable(); err == nil{
//		return filepath.Dir(ex)
//	}
//	return "./"
//}

// 3.6.json渲染
//func main() {
//	r := gin.Default()
//	// gin.H 是map[string]interface{}的缩写
//	r.GET("/someJSON", func(c *gin.Context) {
//		// 第一种，手动拼接
//		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
//	})
//
//	r.GET("/moreJSON", func(c *gin.Context) {
//		var msg struct{
//			Name string `json:"name"`
//			Message string
//			Age int
//		}
//		msg.Name = "小王子"
//		msg.Message = "Hello World!"
//		msg.Age = 18
//		c.JSON(http.StatusOK, msg)
//	})
//	r.Run(":8080")
//}

// 3.7.xml渲染
//func main() {
//	r := gin.Default()
//	r.GET("someXML", func(c *gin.Context) {
//		c.XML(http.StatusOK, gin.H{"message": "hello world!"})
//	})
//
//	r.GET("/moreXML", func(c *gin.Context) {
//		type MessageRecord struct {
//			Name string
//			Message string
//			Age int
//		}
//		var msg MessageRecord
//		msg.Name = "小王子"
//		msg.Message = "hello world!"
//		msg.Age = 18
//		c.XML(http.StatusOK, msg)
//	})
//	r.Run(":8080")
//}

// 3.8.YMAL渲染
//func main() {
//	r := gin.Default()
//	r.GET("/someYAML", func(c *gin.Context) {
//		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
//	})
//	r.Run(":8080")
//}

// 3.9.protobuf渲染
//func main() {
//	r := gin.Default()
//	r.GET("/someProtoBuf", func(c *gin.Context) {
//		reps := []int64{int64(1), int64(2)}
//		label := "test"
//
//		data := &protoexample.Test{
//			Label: &label,
//			Reps: reps,
//		}
//
//		c.ProtoBuf(http.StatusOK, data)
//	})
//	r.Run(":8080")
//}

// 4.获取参数
// 4.1.querystring指的是URL中？后面携带的参数
//func main() {
// 	// default返回一个默认的路由引擎
//	r := gin.Default()
//	r.GET("/user/search", func(c *gin.Context) {
//		username := c.DefaultQuery("username", "小王子")
//		address := c.Query("address")
//
//		// 输出json结果给调用方
//		c.JSON(http.StatusOK, gin.H{
//			"message": "ok",
//			"username": username,
//			"address": address,
//		})
//	})
//	r.Run(":8080")
//}

// 4.2.获取form参数
// 当前端请求的数据通过form表单提交时，例如向/user/search发送一个post请求，
//func main() {
//	r := gin.Default()
//	r.POST("/user/search", func(c *gin.Context) {
//		username := c.PostForm("username")
//		address := c.PostForm("address")
//		// 输出json结果给调用方
//		c.JSON(http.StatusOK, gin.H{
//			"message": "ok",
//			"username": username,
//			"address": address,
//		})
//	})
//	r.Run(":8080")
//}

// 4.3.获取json参数
// 当前端请求的数据通过json提交时，例如向/json发送一个post请求
//func main() {
//	r := gin.Default()
//	r.POST("/json", func(c *gin.Context) {
//		b, _ := c.GetRawData()
//		var m map[string]interface{}
//		// 反序列化
//		_ = json.Unmarshal(b, &m)
//
//		c.JSON(http.StatusOK, m)
//	})
//	r.Run(":8080")
//}

// 4.4.获取path参数 请求参数通过url路径传递，
//func main() {
//	r := gin.Default()
//	r.GET("/user/search/:username/:address", func(c *gin.Context) {
//		username := c.Param("username")
//		address := c.Param("address")
//
//		c.JSON(http.StatusOK, gin.H{
//			"message": "ok",
//			"username": username,
//			"address": address,
//		})
//	})
//	r.Run(":8080")
//}

// 4.5.参数绑定
// 为了能够更方便的获取请求相关参数，提高开发效率，可以基于请求的content—type识别请求数据类型并利用反射机制，自动提取请求中的queryString、form
// 表单、json、xml等参数到结构体中。下面代码展示的是.ShouldBind(),它能够基于请求自动提取json、form表单和queryString类型的数据
//type Login struct {
//	User string `form:"user" json:"user" binding:"required"`
//	Passwd string `form:"passwd" json:"passwd" binding:"required"`
//}
//
//func main() {
//	r := gin.Default()
//
//	// 绑定json示例
//	r.POST("/json", func(c *gin.Context) {
//		var login Login
//		if err := c.ShouldBind(&login); err== nil{
//			fmt.Printf("login info：%#v", login)
//			c.JSON(http.StatusOK, gin.H{
//				"user": login.User,
//				"passwd": login.Passwd,
//			})
//		} else {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		}
//	})
//
//	//绑定form表单示例
//	r.POST("/form", func(c *gin.Context) {
//		var login Login
//		if err := c.ShouldBind(&login); err == nil{
//			c.JSON(http.StatusOK, gin.H{
//				"user": login.User,
//				"passwd": login.Passwd,
//			})
//		} else {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		}
//	})
//
//	// 绑定QueryString的示例
//	r.GET("/query", func(c *gin.Context) {
//		var login Login
//		if err := c.ShouldBind(&login); err == nil{
//			c.JSON(http.StatusOK, gin.H{
//				"user": login.User,
//				"passwd": login.Passwd,
//			})
//		} else {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		}
//	})
//
//	r.Run(":8080")
//}

// 5.文件上传
func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("/Users/whp/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	r.Run()
}