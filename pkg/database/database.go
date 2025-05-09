// handles database connection
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"hims/pkg/config"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func buildDsn(cfg config.Config) string {
	// the database connection string, can be fetched from envs directly
	// or built using fmt.Sprintf(%v...)
	var dsn string
	if cfg.DATABASE_URL != "" {
		dsn = cfg.DATABASE_URL
	} else {
		// expects => #postgres://user:password@localhost:5432/dbname
		dsn = fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable", cfg.DB_TYPE, cfg.DB_USER,
			cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME,
		)
	}
	log.Printf("DB_STRING: %v", dsn)
	return dsn
}

// initializes the database using the env variables
func New(cfg config.Config) (*bun.DB, error) {
	dsn := buildDsn(cfg)
	// bun's custom connection with pgdriver
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	// return an error if db instance is nil
	if sqldb == nil {
		return nil, errors.New("error initializing a database connection")
	}
	// ping establishes a connection
	if err := sqldb.Ping(); err != nil {
		return nil, fmt.Errorf("error while pinging the db: %v", err)
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	//db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	if err := db.DB.Ping(); err != nil {
		return nil, fmt.Errorf("error while pinging the db: %v", err)
	}
	return db, nil
}
