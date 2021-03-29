package application

import (
	"net/http"

	"github.com/jonathanwamsley/practice_apis/custom_pokemon_api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
