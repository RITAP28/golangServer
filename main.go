package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RITAP28/golangServer/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment variables")
	}

	dbUrl := os.Getenv("DATABASE_URL");
	if dbUrl == "" {
		log.Fatal("Database URL is not found in environment variables")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}));

	v1Router := chi.NewRouter();
	v1Router.Get("/ready", handler.HandlerErr);
	v1Router.Get("/err", handler.HandlerReadiness);

	router.Mount("/v1", v1Router);

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
