package database

import (
	"context"
	"hims/pkg/config"
	"testing"
	"time"

	// "gotest.tools/v3/assert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
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

func TestNew_Success(t *testing.T) {
	ctx := context.Background()
	pgContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(pgContainer); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	database, err := New(config.Config{DATABASE_URL: dsn})
	require.NoError(t, err)
	require.NotNil(t, database)
	if err := database.Close(); err != nil {
		t.Errorf("failed to close database: %s", err)
	}
}
