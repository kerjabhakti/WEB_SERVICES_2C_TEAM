package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "HAYYY",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	http.ListenAndServe(":3000", router)
}
