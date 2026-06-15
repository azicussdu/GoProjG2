package repository

import (
	"fmt"
	"time"

	"github.com/azicussdu/GoProjG2/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

// sqlx.DB strukturanyn obektin kaitaru ushin
func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}
