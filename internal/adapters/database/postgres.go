package database

import (
	"database/sql"
	"github.com/alfinkly/api-gateway/internal/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"time"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DBDSN)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	migrationsDir := "./migrations"

	if err = initMigrations(db, migrationsDir); err != nil {
		return nil, err
	}

	return db, nil
}

func initMigrations(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	err = goose.Up(db, dir)
	if err != nil {
		return err
	}
	return nil
}
