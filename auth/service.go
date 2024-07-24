package auth

import "gorm.io/gorm"

type AuthService interface {
}

type AuthServiceDependency struct {
	Db *gorm.DB
}
