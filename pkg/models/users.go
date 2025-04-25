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
	FirstName      string       `json:"firstName" validate:"required,min=3" bun:",notnull"`
	LastName       string       `json:"lastName" validate:"required,min=3" bun:",notnull"`
	Email          string       `json:"email" validate:"required,email" bun:",unique,notnull"`
	Password       string       `json:"password" validate:"required,min=4" bun:",notnull"`
	Role           string       `json:"role" bun:",notnull,default:'doctor'"`
	LicenseNumber  string       `json:"licenseNumber" validate:"required,min=3" bun:",unique,notnull"`
	Specialization string       `json:"specialization" validate:"required,min=3" bun:",notnull"`
	CreatedAt      time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      bun.NullTime `json:"updatedAt"`
	DeletedAt      time.Time    `json:"deletedAt" bun:",soft_delete"`
}

// age is uint (non-negative) and can't go beyond 255
type Patient struct {
	bun.BaseModel
	ID          int          `bun:"id,pk,autoincrement"`
	FirstName   string       `json:"firstName" validate:"required,min=3" bun:",notnull"`
	LastName    string       `json:"lastName" validate:"required,min=3" bun:",notnull"`
	IDNumber    string       `json:"idNumber" validate:"required,min=8,max=10" bun:",unique,notnull"`
	PhoneNumber string       `json:"phoneNumber" validate:"required,min=3,max=15" bun:",unique,notnull"`
	Gender      string       `json:"gender" validate:"required,max=6" bun:",notnull"`
	Age         uint8        `json:"age" validate:"required,max=3" bun:",notnull"`
	Role        string       `json:"role" bun:",notnull,default:'patient'"`
	DateOfBirth time.Time    `json:"dateOfBirth,omitempty"`
	CreatedAt   time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   bun.NullTime `json:"updatedAt"`
	DeletedAt   time.Time    `json:"deletedAt" bun:",soft_delete"`
	Programs    []*Program   `json:"patientPrograms" bun:"rel:has-many,join:id=patient_id"`
}

// a patient can have many programs
type Program struct {
	bun.BaseModel
	ID        int    `bun:"id,pk,autoincrement"`
	Name      string `json:"program" validate:"required,min=3" bun:",unique,notnull"`
	Code      uint   `json:"programCode" validate:"required,max=3" bun:",unique,notnull"`
	PatientID int
}
