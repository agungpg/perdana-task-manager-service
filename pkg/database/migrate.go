package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// RunMigrations executes database migrations in the given direction ("up" or "down").
func RunMigrations(direction string) (err error) {
	if direction == "" {
		direction = "up"
	}

	cfg := NewConfig()

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("init postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("init migrator: %w", err)
	}
	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			err = errors.Join(err, sourceErr)
		}
		if dbErr != nil {
			err = errors.Join(err, dbErr)
		}
	}()

	switch direction {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		err = fmt.Errorf("invalid migration direction %q", direction)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
