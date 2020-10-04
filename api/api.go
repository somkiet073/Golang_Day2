package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Env struct {
}

func (env *Env) GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "result": "OK"})
}

func (env *Env) GetUser(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "result": message})
}

func Post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func Put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func Del(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
