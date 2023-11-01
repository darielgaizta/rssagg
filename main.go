package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/darielgaizta/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environments
	godotenv.Load()

	// Environment validation
	if os.Getenv("PORT") == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB URL is not found in the environment")
	}

	// Open connection to Database
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	// API config will be used to store DB connection
	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	// Server configuration
	server := &http.Server{
		Handler: router,
		Addr:    ":" + os.Getenv("PORT"),
	}

	// Start the server and keep it running forever
	// If there is any error, it will stop and send log.
	log.Printf("Starting server in PORT %v...", os.Getenv("PORT"))
	server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
