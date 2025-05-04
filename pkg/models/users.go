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
	DeletedAt      bun.NullTime `json:"deletedAt" bun:",nullzero,soft_delete"`
}

// age is uint (non-negative) and can't go beyond 255
type Patient struct {
	bun.BaseModel
	ID          int          `bun:"id,pk,autoincrement"`
	FirstName   string       `json:"firstName" validate:"required,min=3" bun:",notnull"`
	LastName    string       `json:"lastName" validate:"required,min=3" bun:",notnull"`
	IDNumber    string       `json:"idNumber" validate:"required,min=8" bun:",unique,notnull"`
	PhoneNumber string       `json:"phoneNumber" validate:"required,min=7" bun:",unique,notnull"`
	Gender      string       `json:"gender" validate:"required,max=6" bun:",notnull"`
	Age         uint8        `json:"age" validate:"required" bun:",notnull"`
	Role        string       `json:"role" bun:",notnull,default:'patient'"`
	DateOfBirth *time.Time   `json:"dateOfBirth"`
	CreatedAt   time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   bun.NullTime `json:"updatedAt"`
	DeletedAt   bun.NullTime `json:"deletedAt" bun:",nullzero,soft_delete"`
	Programs    []*Program   `json:"patientPrograms" bun:"m2m:patient_programs,join:Patient=Program"`
}

// a patient can have many programs
type Program struct {
	bun.BaseModel
	ID       int        `bun:"id,pk,autoincrement"`
	Name     string     `json:"program" validate:"required,min=3" bun:",unique,notnull"`
	Code     uint       `json:"programCode" validate:"required,lte=9999" bun:",unique,notnull"`
	Patients []*Patient `json:"patientsEnrolled" bun:"m2m:patient_programs,join:Program=Patient"`
	// PatientID int
	// to later track doctor who created it
	// DoctorID  int
}

// https://bun.uptrace.dev/guide/relations.html#many-to-many-relation
// intermediary table
type PatientProgram struct {
	PatientID int      `bun:",pk"`
	Patient   *Patient `bun:"rel:belongs-to,join:patient_id=id"`
	ProgramID int      `bun:",pk"`
	Program   *Program `bun:"rel:belongs-to,join:program_id=id"`
}
