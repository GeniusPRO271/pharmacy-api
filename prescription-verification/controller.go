package prescription_verification

import "github.com/gin-gonic/gin"

type Controller struct {
	PrescriptionVerificationService PrescriptionVerificationService
}

func (controller *Controller) RegisterRoutes(router *gin.Engine) {
	// The Routes of the user need to be added in here.
	router.POST("/prescription/:id/verify", controller.VerifyPrescription)
	router.POST("/prescription/:id/reject", controller.RejectPrescription)
}

func (controller *Controller) VerifyPrescription(ctx *gin.Context) {

}
func (controller *Controller) RejectPrescription(ctx *gin.Context) {

}
