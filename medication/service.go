package medication

import "gorm.io/gorm"

type MedicationService interface {
}

type MedicationServiceDependency struct {
	Db *gorm.DB
}
