package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jonathanwamsley/practice_apis/custom_pokemon_api/datasources/postgres/pokemon_db"
)

var router = mux.NewRouter()

func Start() {

	_, err := pokemon_db.NewDatabase()
	if err != nil {
		panic(err)
	}

	mapUrls()

	srv := &http.Server{
		Addr:         ":8081",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	fmt.Println("about to start application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
