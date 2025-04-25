package models

import (
	"time"

	"github.com/uptrace/bun"
)

// table and column names generated automatically
// users table
type Doctor struct {
	bun.BaseModel
	ID             int          `bun:"id,pk,autoincrement"`
	FirstName      string       `json:"firstName" bun:",notnull"`
	LastName       string       `json:"lastName" bun:",notnull"`
	Email          string       `json:"email" bun:",unique,notnull"`
	Password       string       `json:"password" bun:",notnull"`
	LicenseNumber  string       `json:"licenseNumber" bun:",unique,notnull"`
	Specialization string       `json:"specialization" bun:",notnull"`
	CreatedAt      time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      bun.NullTime `json:"updatedAt"`
	DeletedAt      time.Time    `json:"deletedAt" bun:",soft_delete"`
}

// age is uint (non-negative) and can't go beyond 255
type Patient struct {
	bun.BaseModel
	ID          int          `bun:"id,pk,autoincrement"`
	FirstName   string       `json:"firstName" bun:",notnull"`
	LastName    string       `json:"lastName" bun:",notnull"`
	IDNumber    string       `json:"idNumber" bun:",unique,notnull"`
	PhoneNumber string       `json:"phoneNumber" bun:",unique,notnull"`
	Gender      string       `json:"gender" bun:",notnull"`
	Age         uint8        `json:"age" bun:",notnull"`
	DateOfBirth time.Time    `json:"dateOfBirth,omitempty"`
	CreatedAt   time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   bun.NullTime `json:"updatedAt"`
	DeletedAt   time.Time    `json:"deletedAt" bun:",soft_delete"`
}
