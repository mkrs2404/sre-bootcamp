package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mkrs2404/sre-bootcamp/pkg/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	rollback = flag.Bool("rollback", false, "Rollback the last migration")
	steps    = flag.Int("steps", 1, "Number of steps to rollback")
)

func main() {
	flag.Parse()

	cfg := config.Get()
	force_version := os.Getenv("FORCE_VERSION")
	m, e := prepareMigration(cfg.MigrationPath, cfg.DSN)
	if e != nil {
		logrus.Errorf("could not prepare migration : %s", e)
		os.Exit(1)
	}

	if *rollback {
		migrateDBDown(m)
		os.Exit(0)
	}

	if force_version != "" {
		var v int
		var e error
		if v, e = strconv.Atoi(force_version); e != nil {
			logrus.Errorf("could not convert force version to int : %s", e)
			os.Exit(1)
		} else if v < 0 {
			logrus.Error("force version should be a positive integer")
			os.Exit(1)
		}
		forceVersion(m, v)
	} else {
		migrateDBUp(m)
	}
}

func prepareMigration(migrationPath, dsn string) (*migrate.Migrate, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "could not open database connection")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "could not get postgres driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationPath),
		"postgres",
		driver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not get create migration instance")
	}

	return m, nil
}

func migrateDBUp(m *migrate.Migrate) {
	logrus.Info("Starting the migration ...")
	if e := m.Up(); e != nil {
		logrus.Errorf("could not migrate up : %s", e)
	}
	logrus.Info("Migration completed.")
}

func migrateDBDown(m *migrate.Migrate) {
	logrus.Info("Starting the rollback ...")
	if e := m.Steps(-*steps); e != nil {
		logrus.Errorf("could not rollback : %s", e)
		os.Exit(1)
	}
	logrus.Info("Rollback completed.")
}

func forceVersion(m *migrate.Migrate, version int) {
	logrus.Printf("Starting the migration for version %d...", version)
	if e := m.Force(version); e != nil {
		logrus.Printf("could not force to a stable version : %s", e)
	}
}
