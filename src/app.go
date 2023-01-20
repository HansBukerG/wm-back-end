package app

import (
	"log"
	"net/http"

	"github.com/HansBukerG/wm-back-end/src/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func App_init() {
	host := "localhost"
	port := "8000"
	route := mux.NewRouter()

	routes.RegisterProductsRoutes(route)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET"},
	})

	handler := c.Handler(route)
	log.Println("ready to listen in: " + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, handler))
	http.Handle("/", route)

}
