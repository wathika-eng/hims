package models

import (
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

func maskName(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[0]) + "."
}

func maskID(id string) string {
	if len(id) < 5 {
		return id
	}
	return id[:2] + "****" + id[len(id)-2:]
}

func maskPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-3:]
}

// table and column names generated automatically
// users table
type Doctor struct {
	bun.BaseModel
	ID             int          `bun:"id,pk,autoincrement"`
	FirstName      string       `json:"firstName" validate:"required,min=3" bun:",notnull"`
	LastName       string       `json:"lastName" validate:"required,min=3" bun:",notnull"`
	Email          string       `json:"email" validate:"required,email" bun:",unique,notnull"`
	Password       string       `json:"-" bun:",notnull"`
	Role           string       `json:"role" bun:",notnull,default:'doctor'"`
	LicenseNumber  string       `json:"licenseNumber" validate:"required,min=3" bun:",unique,notnull"`
	Specialization string       `json:"specialization" validate:"required,min=3" bun:",notnull"`
	CreatedAt      time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      bun.NullTime `json:"updatedAt"`
	DeletedAt      bun.NullTime `json:"deletedAt" bun:",nullzero,soft_delete"`
}

// Doctor input for signup — only used for JSON binding, never serialized
type DoctorSignupInput struct {
	FirstName      string `json:"firstName" validate:"required,min=3"`
	LastName       string `json:"lastName" validate:"required,min=3"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=4"`
	LicenseNumber  string `json:"licenseNumber" validate:"required,min=3"`
	Specialization string `json:"specialization" validate:"required,min=3"`
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

func (p Patient) MarshalJSON() ([]byte, error) {
	type alias struct {
		ID          int              `json:"ID"`
		FirstName   string           `json:"firstName"`
		LastName    string           `json:"lastName"`
		IDNumber    string           `json:"idNumber"`
		PhoneNumber string           `json:"phoneNumber"`
		Gender      string           `json:"gender"`
		Age         uint8            `json:"age"`
		Role        string           `json:"role"`
		DateOfBirth *time.Time       `json:"dateOfBirth"`
		CreatedAt   time.Time        `json:"createdAt"`
		UpdatedAt   bun.NullTime     `json:"updatedAt"`
		DeletedAt   bun.NullTime     `json:"deletedAt"`
		Programs    []*Program       `json:"patientPrograms"`
	}
	return json.Marshal(alias{
		ID:          p.ID,
		FirstName:   p.FirstName,
		LastName:    maskName(p.LastName),
		IDNumber:    p.IDNumber,
		PhoneNumber: maskPhone(p.PhoneNumber),
		Gender:      p.Gender,
		Age:         p.Age,
		Role:        p.Role,
		DateOfBirth: p.DateOfBirth,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
		Programs:    p.Programs,
	})
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
