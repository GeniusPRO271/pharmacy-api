package database

import "gorm.io/gorm"

type UserDBModel struct {
	gorm.Model
	FullName     string `gorm:"type:varchar(100);not null"`
	RUN          string `gorm:"type:varchar(12);unique;not null"`
	Email        string `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	PhoneNumber  string `gorm:"type:varchar(20)"`
	Address      string `gorm:"type:text"`
}
