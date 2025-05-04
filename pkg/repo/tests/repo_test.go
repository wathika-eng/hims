package repo

import (
	"context"
	"hims/pkg/models"
	"hims/pkg/repo"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func TestCustomerRepo(t *testing.T) {
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
	if err != nil {
		t.Fatalf("error occured while running tests: %v", err.Error())
	}
	defer func() {
		if err := testcontainers.TerminateContainer(pgContainer); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	log.Println(connStr)
	config, err := pgx.ParseConfig(connStr)
	if err != nil {
		panic(err)
	}
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())
	repo := repo.NewRepo(db)
	repo.Up()
	// Insert a new patient
	err = repo.InsertNewPatient(&models.Patient{
		FirstName: "Mike",
		LastName:  "Go",
	})
	assert.NoError(t, err, "inserting patient should not error")

	// // Fetch all patients
	patients, err := repo.FetchPatients()
	assert.NoError(t, err, "fetching patients should not error")
	assert.Len(t, patients, 1, "there should be exactly one patient")

	// // Check if the patient data is correct
	patient := patients[0]
	assert.Equal(t, "Mike", patient.FirstName)
	assert.Equal(t, "Go", patient.LastName)

}
