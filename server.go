package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"./database/redis"
	"./routes"
	"./settings"
	"net/http"
)

func main() {
	redis.Init()
	settings.Init()
	r := routes.InitRoutes()

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3001", r)
}
