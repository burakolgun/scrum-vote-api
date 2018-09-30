package routes

import (
	"../controller"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/register", controller.Register).Methods("POST")
	return router
}