// package repo provides interactions to the database
package repo

import (
	"fmt"

	"github.com/uptrace/bun"
)

// define repo struct referencing our db already initialized
type Repo struct {
	db *bun.DB
}

// constructor for new Repo
func NewRepo(db *bun.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// get the db stats
func (r *Repo) Stats() map[string]string {
	stats := r.db.DB.Stats()

	return map[string]string{
		"MaxOpenConnections": fmt.Sprintf("%d", stats.MaxOpenConnections),
		"OpenConnections":    fmt.Sprintf("%d", stats.OpenConnections),
		"InUse":              fmt.Sprintf("%d", stats.InUse),
		"Idle":               fmt.Sprintf("%d", stats.Idle),
		"WaitCount":          fmt.Sprintf("%d", stats.WaitCount),
		"WaitDuration":       stats.WaitDuration.String(),
		"MaxIdleClosed":      fmt.Sprintf("%d", stats.MaxIdleClosed),
		"MaxLifetimeClosed":  fmt.Sprintf("%d", stats.MaxLifetimeClosed),
	}
}
