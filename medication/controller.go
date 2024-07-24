package medication

import "github.com/gin-gonic/gin"

type Controller struct {
	MedicationService MedicationService
}

func (controller *Controller) RegisterRoutes(router *gin.Engine) {
	router.GET("/medication", controller.GetMedication)
	router.GET("/medication/:id", controller.GetMedicationById)
	router.POST("/medication", controller.AddMedication)
	router.PUT("/medication/:id", controller.UpdateMedication)
	router.DELETE("/medication/:id", controller.DeleteMedication)
}

func (controller *Controller) GetMedication(ctx *gin.Context) {

}
func (controller *Controller) GetMedicationById(ctx *gin.Context) {

}
func (controller *Controller) AddMedication(ctx *gin.Context) {

}
func (controller *Controller) UpdateMedication(ctx *gin.Context) {

}
func (controller *Controller) DeleteMedication(ctx *gin.Context) {

}
