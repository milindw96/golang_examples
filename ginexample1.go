package main

// Gin Framework basic example
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", displayString)
	r.GET("/user/:name", userInfo)
	r.Run(":1313")

}

func userInfo(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"Name": name,
	})
}

func displayString(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
