package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//func main() {
//	router := gin.Default()
//
//	router.GET("/user/:name", func(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "hello %s", name)
//	})
//
//	router.GET("/user/:name/*action", func(c *gin.Context) {
//		name := c.Param("name")
//		action := c.Param("action")
//		message := name + " is " + action
//		c.String(http.StatusOK, message)
//	})
//
//	router.POST("/user/:name/*action", func(c *gin.Context) {
//		b := c.FullPath() == "/user/:name/*action"
//		c.String(http.StatusOK, "%t", b)
//	})
//
//	router.GET("/user/groups", func(c *gin.Context) {
//		c.String(http.StatusOK, "The available groups are [...]")
//	})
//
//	router.Run()
//}

/*
query string
*/
//func main() {
//	router := gin.Default()
//
//	router.GET("/welcome", func(c *gin.Context) {
//		firstname := c.DefaultQuery("firstname", "Guest")
//		lastName := c.Query("lastname")
//		c.String(http.StatusOK, "hello %s %s", firstname, lastName)
//	})
//	router.Run()
//}

/*
urlencided form
*/
//func main() {
//	router := gin.Default()
//
//	router.POST("form_post", func(c *gin.Context) {
//		message := c.PostForm("message")
//		nick := c.DefaultPostForm("nike", "haha")
//
//		c.JSON(http.StatusOK, gin.H{
//			"status":  "posted",
//			"message": message,
//			"nike":    nick,
//		})
//	})
//	router.Run()
//}

//func main() {
//	router := gin.Default()
//
//	router.POST("/post", func(c *gin.Context) {
//		id := c.Query("id")
//		page := c.DefaultQuery("page", "0")
//		name := c.PostForm("name")
//		message := c.PostForm("message")
//
//		fmt.Printf("id: %s; page: %s; name: %s; message: %s;", id, page, name, message)
//
//		c.JSON(http.StatusOK, gin.H{
//			"id":      id,
//			"page":    page,
//			"name":    name,
//			"message": message,
//		})
//	})
//
//	router.Run()
//}

/*
query map
*/

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	router.Run()
}
