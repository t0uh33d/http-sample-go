package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviroment variable
	godotenv.Load()

	// Getting the port from environment
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("NO PORT SPECIFIED")
	}

	// initializing the Chi router
	var router *chi.Mux = chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// specifying a http server
	var server *http.Server = &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	// for versioning
	var v1Router *chi.Mux = chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handleError)

	router.Mount("/v1", v1Router)

	log.Printf("STARTING THE SERVER ON PORT : %s", portStr)

	// listening to the requests on the server

	var serverError error = server.ListenAndServe()

	if serverError != nil {
		log.Fatal("FAILED TO START THE SERVER : \n", serverError)
	}
}
