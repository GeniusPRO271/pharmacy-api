package prescription

import "gorm.io/gorm"

type PrescriptionService interface {
}

type PrescriptionServiceDependency struct {
	Db *gorm.DB
}
