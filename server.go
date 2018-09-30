package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"./database/Redis"
	"./routes"
	"net/http"
)

func main() {
	redis.Init()
	r := routes.InitRoutes()

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3001", r)
}
