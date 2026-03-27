package api

import (
	"net/http"

	"github.com/Fozzyack/badminton-tracker-backend/internal/services"
	"github.com/rs/zerolog"
)

type AuthHandler struct {
	Logger      zerolog.Logger
	AuthService *services.AuthService
}

func NewAuthHandler(logger zerolog.Logger, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Logger:      logger,
		AuthService: authService,
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest
	err := DecodeJSON(r, &loginRequest)
	if err != nil {
		ah.Logger.Error().Err(err).Msg("Failed to decode request")
		SendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	session, err := ah.AuthService.Login(r.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		ah.Logger.Error().Err(err).Msg("Failed to login")
		SendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if session == nil {
		ah.Logger.Error().Msg("Failed to login")
		SendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	SendJSON(w, session.Token)

}
