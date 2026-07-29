package repository

import (
	"context"
	"errors"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

const uniqueViolationCode = "23505"

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(dbObj *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: dbObj}
}

func (pur *PostgresUserRepo) Create(ctx context.Context, user model.User) (int, error) {
	err := pur.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return 0, apperrors.Conflict("email is already registered", err)
		}
		return 0, apperrors.Internal("failed to create user", err)
	}

	return user.ID, nil
}

func (pur *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := pur.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, apperrors.NotFound("user not found", err)
		}
		return model.User{}, apperrors.Internal("failed to get user", err)
	}

	return user, nil
}

func (pur *PostgresUserRepo) GetByID(ctx context.Context, id int) (model.User, error) {
	var user model.User

	err := pur.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, apperrors.NotFound("user not found", err)
		}
		return model.User{}, apperrors.Internal("failed to get user", err)
	}

	return user, nil
}
