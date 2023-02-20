package application

import (
	"crud-echo/config"
	"crud-echo/routes"
	gotenv "github.com/subosito/gotenv"
	"log"
	"net/http"
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
	log.Println("Your service is up and running at " + addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
