package routes

import (
	controllers "../controller"
	authentication "../service/auth_service"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}