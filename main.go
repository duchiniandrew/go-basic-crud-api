package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id: 1, Username: "user1"},
	{Id: 2, Username: "user2"},
	{Id: 3, Username: "user3"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		println("Error: ", err)
	}

	for _, user := range users {
		if user.Id == userId {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
}

func addUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/albums", addUser)
	router.Run("localhost:3010")
}
