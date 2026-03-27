package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Fozzyack/badminton-tracker-backend/internal/app"
	"github.com/Fozzyack/badminton-tracker-backend/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRouter(app)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	app.Logger.Info().Str("port", port).Msg("Starting server")
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal().Err(err).Msg("Failed to start server")
	}

}
