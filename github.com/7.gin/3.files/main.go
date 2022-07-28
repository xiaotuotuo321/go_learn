package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func main() {
//	router := gin.Default()
//
//	router.MaxMultipartMemory = 8 << 20 // 8MiB
//	router.POST("/upload", func(c *gin.Context) {
//		file, _ := c.FormFile("file")
//		log.Println(file.Filename)
//
//		// 把文件上传到特定的地方 copy了一下
//		//c.SaveUploadedFile(file, dst)
//
//		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
//	})
//	router.Run()
//}

// 多个文件
func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 >> 20

	router.POST("/upload", func(c *gin.Context) {
		// multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			fmt.Println(file)
			//c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files upload", len(files)))
	})
	router.Run()
}
