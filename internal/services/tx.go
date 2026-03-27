package services

import (
	"context"
	"database/sql"
)

type TxManager interface {
	WithTx(ctx context.Context, fn func(tx *sql.Tx) error) error
}

type SQLTxMaanger struct {
	db *sql.DB
}

func NewSQLTxManager(db *sql.DB) *SQLTxMaanger {
	return &SQLTxMaanger{db: db}
}

func (s *SQLTxMaanger) WithTx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func () {
		if err != nil {
			tx.Rollback()
		}
	}()


	err = fn(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
