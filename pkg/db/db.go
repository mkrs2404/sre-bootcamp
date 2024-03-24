package db

import (
	"errors"

	"github.com/mkrs2404/sre-bootcamp/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Connect(cfg config.Config) (*DB, error) {
	dsn := cfg.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, errors.New("error connecting to database")
	}

	return &DB{
		DB: db,
	}, nil
}
