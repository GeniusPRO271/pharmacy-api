package prescription

import "github.com/gin-gonic/gin"

type Controller struct {
	PrescriptionService PrescriptionService
}

func (controller *Controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/prescription", controller.UploadPrescription)
	router.GET("/prescription", controller.GetPrescriptions)
	router.GET("/prescription/:id", controller.GetPrescriptionById)
	router.PUT("/prescription/:id", controller.UpdatePrescription)
	router.DELETE("/prescription/:id", controller.DelePrescription)
}

func (controller *Controller) UploadPrescription(ctx *gin.Context) {

}
func (controller *Controller) GetPrescriptions(ctx *gin.Context) {

}
func (controller *Controller) GetPrescriptionById(ctx *gin.Context) {

}
func (controller *Controller) UpdatePrescription(ctx *gin.Context) {

}
func (controller *Controller) DelePrescription(ctx *gin.Context) {

}
