package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{

			"message": "GET users",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("POST user %s", user.Username),
		})
	})

	r.PUT("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("PUT user %s", username),
		})
	})

	r.DELETE("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("DELETE user %s", username),
		})
	})
	r.Run()
}
