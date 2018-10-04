package routes

import (
	"../controller"
	authentication "../service/auth_service"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controller.Login).Methods("POST")

	router.HandleFunc("/register", controller.Register).Methods("POST")

	router.Handle("/refresh-token", negroni.New(negroni.HandlerFunc(controller.RefreshToken), )).Methods("GET")

	router.Handle(
		"/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controller.Logout),
		)).Methods("GET")
	return router
}
