package services

import (
	"context"
	"database/sql"

	"github.com/Fozzyack/badminton-tracker-backend/internal/auth"
	"github.com/Fozzyack/badminton-tracker-backend/internal/models"
	"github.com/Fozzyack/badminton-tracker-backend/internal/store"
)

type AuthService struct {
	TxManager    TxManager
	UserStore    store.UserStore
	SessionStore store.SessionStore
}

func NewAuthService(TxManager TxManager, userStore store.UserStore, sessionStore store.SessionStore) *AuthService {
	return &AuthService{
		TxManager:    TxManager,
		UserStore:    userStore,
		SessionStore: sessionStore,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (*models.Session, error) {
	token, err := auth.GenerateToken()
	if err != nil {
		return nil, err
	}

	var session *models.Session
	err = as.TxManager.WithTx(ctx, func(tx *sql.Tx) error {
		user, err := as.UserStore.GetUserByEmail(ctx, email)
		if err != nil {
			return err
		}

		err = auth.ComparePassword(password, user.Password)
		if err != nil {
			return err
		}

		session, err = as.SessionStore.CreateSessionTx(ctx, tx, user.ID, token)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return session, nil
}
