package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

const uniqueViolationCode = "23505"

type PostgresUserRepo struct {
	db *sqlx.DB
}

func NewPostgresUserRepo(dbObj *sqlx.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: dbObj}
}

func (pur *PostgresUserRepo) Create(ctx context.Context, user model.User) (int, error) {
	query := `
    INSERT INTO users (
        full_name, email, password_hash, role, created_at, updated_at
    ) VALUES (
        :full_name, :email, :password_hash, :role, :created_at, :updated_at
    )
    RETURNING id
  `

	rows, err := pur.db.NamedQueryContext(ctx, query, user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return 0, apperrors.Conflict("email is already registered", err)
		}
		return 0, apperrors.Internal("failed to create user", err)
	}

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, apperrors.Internal("failed to create user", err)
		}
	}

	return id, nil
}

func (pur *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	query := `
    SELECT id, full_name, email, password_hash, role, created_at, updated_at
    FROM users
    WHERE email = $1
    LIMIT 1
  `

	err := pur.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, apperrors.NotFound("user not found", err)
		}
		return model.User{}, apperrors.Internal("failed to get user", err)
	}

	return user, nil
}

func (pur *PostgresUserRepo) GetByID(ctx context.Context, id int) (model.User, error) {
	var user model.User

	query := `
    SELECT id, full_name, email, password_hash, role, created_at, updated_at
    FROM users
    WHERE id = $1
    LIMIT 1
  `

	err := pur.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, apperrors.NotFound("user not found", err)
		}
		return model.User{}, apperrors.Internal("failed to get user", err)
	}

	return user, nil
}
