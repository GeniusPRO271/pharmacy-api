package database

import (
	"gorm.io/gorm"
)

type ClientDBModel struct {
	gorm.Model
	FullName     string `gorm:"type:varchar(100);not null"`
	RUN          string `gorm:"type:varchar(12);unique;not null"`
	Email        string `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	PhoneNumber  string `gorm:"type:varchar(20)"`
	Address      string `gorm:"type:text"`
}

type VerifierDBModel struct {
	gorm.Model
	FullName     string `gorm:"type:varchar(100);not null"`
	RUN          string `gorm:"type:varchar(12);unique;not null"`
	Email        string `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	PhoneNumber  string `gorm:"type:varchar(20)"`
	Role         string `gorm:"type:varchar(50)"`
}

type PrescriptionDBModel struct {
	gorm.Model
	ClientID         uint          `gorm:"not null"`
	PrescriptionFile string        `gorm:"type:varchar(255);not null"`
	Status           string        `gorm:"type:varchar(50);default:'pending'"`
	Client           ClientDBModel `gorm:"foreignkey:ClientID"`
}

type VerificationLogDBModel struct {
	gorm.Model
	PrescriptionID uint                `gorm:"not null"`
	VerifierID     uint                `gorm:"not null"`
	Action         string              `gorm:"type:varchar(50);not null"`
	Prescription   PrescriptionDBModel `gorm:"foreignkey:PrescriptionID"`
	Verifier       VerifierDBModel     `gorm:"foreignkey:VerifierID"`
}
