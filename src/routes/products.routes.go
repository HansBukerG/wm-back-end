package routes

import (
	"github.com/HansBukerG/wm-back-end/src/controller"
	"github.com/gorilla/mux"
)

var RegisterProductsRoutes = func(router *mux.Router) {
	router.HandleFunc("/search/{searchString}", controller.GetProductByString).Methods("GET")
	router.HandleFunc("/search/", controller.GetProductsDefault).Methods("GET")
}
