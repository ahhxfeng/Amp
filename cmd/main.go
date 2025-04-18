package main

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.String(200, "Hello world!")
}

func main() {
	// simple go http server router use gin
	r := gin.Default()
	r.GET("/index", Index)
	r.GET("/", Index)

	// start the defualt http server 0.0.0.:8080
	r.Run()
}
