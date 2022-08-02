package application

import (
	"net/http"
	"github.com/subosito/gotenv"
	"crud-echo/config"
	"crud-echo/routes"
	"log"
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


// 	addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
// 	conf := gin.Config{
// 		ListenAddr: addr,
// 		Handler:    router.Router(),
// 		OnStarting: func() {
// 			log.Info().Msg("Your service is up and running at " + addr)
// 		},
// 	}
//
// 	gin.Run(conf)



//     addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
//     e := echo.New()
//         e.GET("/", func(c echo.Context) error {
//             return c.String(http.StatusOK, "Hello ges")
//         })
//         e.Logger.Fatal(e.Start(":8888"))
}
