package interfaces

import "github.com/gin-gonic/gin"

type BaseController interface {
	SetRouter(*gin.RouterGroup)
	GetRelativePath() string
}
