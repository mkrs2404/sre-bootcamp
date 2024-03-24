package gormimpl

import (
	"github.com/mkrs2404/sre-bootcamp/pkg/db"
	"github.com/mkrs2404/sre-bootcamp/pkg/repo/healthcheck"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// healthCheckRepository is an implementation of the healthCheckRepository
// interface that uses Gorm to query the database.
type healthCheckRepository struct {
	DB *gorm.DB
}

func NewHealthCheckRepository(db *db.DB) healthcheck.Repository {
	return &healthCheckRepository{
		DB: db.DB,
	}
}

// DBHealth returns error if the database is unavailable.
func (r *healthCheckRepository) DBHealth() error {
	sqlDB, _ := r.DB.DB()
	if err := sqlDB.Ping(); err != nil {
		logrus.Errorf("unable to ping the database, %v", err)
		return err
	}
	return nil
}
