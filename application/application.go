package application

import (
	"crud-echo/config"
	"crud-echo/routes"
	"net/http"

	log "github.com/sirupsen/logrus"
	gotenv "github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}
func StartApp() {
	addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
	s := http.Server{
		Addr:    addr,
		Handler: routes.Routes(),
	}
	log.Info("Your service is up and running at ", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Error(err)
	}
}
