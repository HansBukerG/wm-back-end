package app

import (
	"log"
	"net/http"

	"github.com/HansBukerG/wm-back-end/src/routes"
	"github.com/gorilla/mux"
)

func App_init() {
	route := mux.NewRouter()

	routes.RegisterProductsRoutes(route)

	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8000", route))

}
