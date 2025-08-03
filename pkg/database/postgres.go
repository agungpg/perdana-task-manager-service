package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// DB is the global database instance
var DB *bun.DB

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	SSLMode  string
}

// NewConfig creates a new database configuration from environment variables
func NewConfig() *Config {
	return &Config{
		Host:     getEnvOrDefault("DATABASE_HOST", "localhost"),
		Port:     getEnvOrDefault("DATABASE_PORT", "5432"),
		Username: getEnvOrDefault("DATABASE_USERNAME", "postgres"),
		Password: getEnvOrDefault("DATABASE_PASS", ""),
		Name:     getEnvOrDefault("DATABASE_NAME", "ptm_db"),
		SSLMode:  getEnvOrDefault("DATABASE_SSLMODE", "disable"),
	}
}

// getEnvOrDefault returns the environment variable value or a default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// buildDSN constructs the database connection string
func (c *Config) buildDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Username, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

// Connect establishes a connection to the PostgreSQL database
func Connect() error {
	return ConnectWithConfig(NewConfig())
}

// ConnectWithConfig establishes a connection to the PostgreSQL database using the provided config
func ConnectWithConfig(config *Config) error {
	dsn := config.buildDSN()

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
