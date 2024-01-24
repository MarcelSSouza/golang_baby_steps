package main

import "github.com/gin-gonic/gin"

func mainB() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})
	
    r.GET("/bora_bahia", func(c *gin.Context) {
		c.String(200, "minha porra")	
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
