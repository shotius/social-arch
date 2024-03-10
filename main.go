package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID uint `json:id`
	Name string `json:name`
	LastName string `json:lastName`
	PhotoUrl string `json:photoUrl`
}

var user = User{ID: 0, Name: "Shota", LastName: "Archem", PhotoUrl: ""}

func getCurrentUser(c *gin.Context) {
c.IndentedJSON(http.StatusOK, user)
}

func main() {
	router := gin.Default()
	router.GET("/user", getCurrentUser)

	router.Run("0.0.0.0:8080")
}