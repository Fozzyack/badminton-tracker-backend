package store

import (
	"context"
	"database/sql"

	"github.com/Fozzyack/badminton-tracker-backend/internal/models"
)

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserStore(db *sql.DB) UserStore {
	return &PostgresStore{db: db}
}

func (ps *PostgresStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `
	SELECT id, email, password_hash, created_at, updated_at
	FROM users
	WHERE email = $1
	`

	err := ps.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
