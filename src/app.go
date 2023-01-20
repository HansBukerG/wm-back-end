package app

import (
	"log"
	"net/http"

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
	log.Fatal(http.ListenAndServe("localhost:8000", handler))
	http.Handle("/", route)

}
