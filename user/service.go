package user

import "gorm.io/gorm"

type UserService interface {
}

type UserServiceDependency struct {
	Db *gorm.DB
}
