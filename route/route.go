package route

import (
	"hellogoDay2/api"

	"github.com/gin-gonic/gin"
)

type Env struct {
	api.Env
}

func (e Env) Router() *gin.Engine {
	r := gin.Default()
	r.GET("/GetPost", e.GetPost)
	r.GET("/GetUser/:name/:action", e.GetUser)
	r.POST("/")
	r.PUT("/:id")
	r.DELETE("/:id")

	return r
}
