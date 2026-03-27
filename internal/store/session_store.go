package store

import (
	"context"
	"database/sql"

	"github.com/Fozzyack/badminton-tracker-backend/internal/models"
)

type SessionStore interface {
	CreateSession(ctx context.Context, userID, token string) (*models.Session, error)
	CreateSessionTx(ctx context.Context, tx *sql.Tx, userID, token string) (*models.Session, error)
	GetSessionByToken(ctx context.Context, token string) (*models.Session, error)
}

func NewSessionStore(db *sql.DB) SessionStore {
	return &PostgresStore{db: db}
}

func (ps *PostgresStore) createSession(ctx context.Context, dbtx DBTX, userID, token string) (*models.Session, error) {
	session := &models.Session{}

	query := `
	INSERT INTO sessions (user_id, token)
	VALUES ($1, $2)
	RETURNING id, user_id, token, created_at, updated_at
	`

	err := dbtx.QueryRowContext(ctx, query, userID, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (ps *PostgresStore) CreateSession(ctx context.Context, userID, token string) (*models.Session, error) {
	return ps.createSession(ctx, ps.db, userID, token)
}

func (ps *PostgresStore) CreateSessionTx(ctx context.Context, tx *sql.Tx, userID, token string) (*models.Session, error) {
	return ps.createSession(ctx, tx, userID, token)
}

func (ps *PostgresStore) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	session := &models.Session{}

	query := `
	SELECT id, user_id, token, created_at, updated_at
	FROM sessions
	WHERE token = $1
	`

	err := ps.db.QueryRowContext(ctx, query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}
