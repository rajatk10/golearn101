package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//engine := gin.New()  no logger and panic prevention
	engine := gin.Default()
	//defining route with anonymous function handler
	engine.GET("/", func(c *gin.Context) {
		//gin.Context is a struct consists of HTTP request details, response writer, route parameters
		//query strings, headers, cookies, form values, JSON values, etc.

		//fmt.Println("gin framework endpoint '/' hit")
		c.String(200, "Hello World")
		//c is the context passed by gin
	})
	//Another way
	engine.Handle("HEAD", "/", func(c *gin.Context) {
		//fmt.Println("gin framework endpoint 'HEAD' hit")
		c.String(200, "")
	})

	engine.GET("/query/:id", queryHandler)
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	engine.Run(":8087") //starting on port 8087
	// or log.Fatal(engine.Run(":8087"))

}

func queryHandler(c *gin.Context) {
	//Get parameter from request

	userID := c.Param("id")
	//Query parameter
	//fmt.Printf("gin framework endpoint '/query/:id' hit %s", userID)
	userName := c.Query("name") //query string

	//Send Response
	c.JSON(200, gin.H{
		"user_id":   userID,
		"user_name": userName,
	})

}
