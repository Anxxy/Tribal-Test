package controllers

import (
	"github.com/gin-gonic/gin"
	interfaces2 "tribal/controllers/interfaces"
	"tribal/services/interfaces"
)

type JokesController struct {
	Router  *gin.RouterGroup
	Service interfaces.Service
	Path    string
}

func NewJokeController(path string, service interfaces.Service) interfaces2.BaseController {
	return JokesController{
		Router:  nil,
		Service: service,
		Path:    path,
	}
}

func (c JokesController) SetRouter(router *gin.RouterGroup) {
	c.Router = router

	c.Router.GET("/chistes", c.GetJokes)
}

func (c JokesController) GetRelativePath() string {
	return c.Path
}

func (c JokesController) GetJokes(ctx *gin.Context) {
	jokes, err := c.Service.GetNMusicJokes(15)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, jokes)
}
