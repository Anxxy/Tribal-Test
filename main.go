package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
	"tribal/controllers"
	"tribal/controllers/interfaces"
	"tribal/server"
	"tribal/services"
)

func main() {
	log.SetLevel(log.DebugLevel)

	restClient := resty.New()

	service := services.NewService(restClient)

	srv := server.NewServer([]interfaces.BaseController{
		controllers.NewJokeController("/", service),
	})

	if err := srv.Listen(); err != nil {
		log.Panicln(err)
	}
}
