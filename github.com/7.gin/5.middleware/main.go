package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
}
