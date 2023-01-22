package app

import (
	"log"
	"net/http"

	"github.com/HansBukerG/wm-back-end/src/app_env"
	"github.com/HansBukerG/wm-back-end/src/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func App_init() {

	route := mux.NewRouter()

	routes.RegisterProductsRoutes(route)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET"},
	})

	handler := c.Handler(route)

	httpPort := app_env.GetEnv("HTTP_PORT")

	log.Println("ready to listen in port:" + httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, handler))
	http.Handle("/", route)

}
