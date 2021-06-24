package main


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



