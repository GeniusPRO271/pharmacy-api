package prescription_verification

import "gorm.io/gorm"

type PrescriptionVerificationService interface {
}

type PrescriptionVerificationDependency struct {
	Db *gorm.DB
}
