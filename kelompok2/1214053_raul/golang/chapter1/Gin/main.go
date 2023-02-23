package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "maul",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	http.ListenAndServe(":7000", router)
}
