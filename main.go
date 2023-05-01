package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", HelloWorld)
	router.GET("/:name", IndexHandler)
	router.Run()
}

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}
