package app

import (
	"os"

	"github.com/Fozzyack/badminton-tracker-backend/database"
	"github.com/Fozzyack/badminton-tracker-backend/internal/api"
	"github.com/Fozzyack/badminton-tracker-backend/internal/env"
	"github.com/Fozzyack/badminton-tracker-backend/internal/services"
	"github.com/Fozzyack/badminton-tracker-backend/internal/store"
	"github.com/Fozzyack/badminton-tracker-backend/migrations"
	"github.com/rs/zerolog"
)

type Application struct {
	Logger      zerolog.Logger
	AuthService *services.AuthService
	AuthHandler *api.AuthHandler
}

func NewApplication() (*Application, error) {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	if env.GetProduction() {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	pgDB, err := database.Open()
	if err != nil {
		return nil, err
	}

	err = database.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		return nil, err
	}

	// stores
	userStore := store.NewUserStore(pgDB)
	sessionStore := store.NewSessionStore(pgDB)

	// services
	txManager := services.NewSQLTxManager(pgDB)
	authService := services.NewAuthService(txManager, userStore, sessionStore)

	// handlers
	authHandler := api.NewAuthHandler(logger, authService)

	app := &Application{
		Logger:      logger,
		AuthService: authService,
		AuthHandler: authHandler,
	}

	return app, nil
}
