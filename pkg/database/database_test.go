package database

import (
	"hims/pkg/config"
	"testing"

	// "gotest.tools/v3/assert"
	"github.com/stretchr/testify/assert"
)

// build dsn string from env vars
func TestBuildDsn(t *testing.T) {
	cfg := config.Config{
		DB_TYPE:     "postgres",
		DB_USER:     "user",
		DB_PASSWORD: "pass",
		DB_HOST:     "localhost",
		DB_PORT:     "5432",
		DB_NAME:     "testdb",
	}
	expected := "postgres://user:pass@localhost:5432/testdb?sslmode=disable"
	dsn := buildDsn(cfg)
	assert.Equal(t, expected, dsn)
}

// test will pass since postgres is available locally and using peer/indent auth method
func TestNew_Success(t *testing.T) {
	cfg := config.Config{
		DATABASE_URL: "postgres://postgres:pass@localhost:5432/postgres",
	}

	database, err := New(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, database)
}
