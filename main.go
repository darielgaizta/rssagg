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
	// Load environments
	godotenv.Load()

	if os.Getenv("PORT") == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// Router configuration
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// V1 Router
	// Create versions of your server for good practice!
	// Path /healthz is a standard practice to test and see if server is up and running.
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/errz", handlerErr)

	router.Mount("/v1", v1Router)

	// Server configuration
	server := &http.Server{
		Handler: router,
		Addr:    ":" + os.Getenv("PORT"),
	}

	// Start the server and keep it running forever
	// If there is any error, it will stop and send log.
	log.Printf("Starting server in PORT %v...", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
