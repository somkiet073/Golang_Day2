package main

import (
	"hellogoDay2/workshop/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"username": name,
		"password": "1234",
		"fullname": "name name",
	})
}

func PostUser(c *gin.Context) {

	var users database.User
	c.Bind(&users)
	c.JSON(http.StatusOK, gin.H{
		"id":       users.Id,
		"username": users.Username,
		"password": users.Password,
		"fullname": users.Fullname,
	})
}

func main() {
	r := gin.Default()
	r.GET("/GetAllUser", GetAllUser)
	r.GET("/GetUser/:id", GetUser)
	r.POST("/", PostUser)

	r.Run(":8180")
}
