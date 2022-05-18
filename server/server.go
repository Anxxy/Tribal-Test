package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tribal/controllers/interfaces"
)

type Server struct {
	App *gin.Engine
}

func NewServer(controllers []interfaces.BaseController) *Server {
	s := &Server{App: gin.Default()}

	s.InitializeMiddlewares()
	s.InitializeControllers(controllers)
	return s
}

func (s *Server) InitializeMiddlewares() {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	s.App.Use(cors.New(corsConf))
}

func (s *Server) InitializeControllers(controllers []interfaces.BaseController) {
	for _, controller := range controllers {
		controller.SetRouter(s.App.Group(controller.GetRelativePath()))
	}
}

func (s *Server) Listen() error {
	return s.App.Run("127.0.0.1:8080")
}
