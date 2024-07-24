package user

import "github.com/gin-gonic/gin"

type Controller struct {
	UserService UserService
}

func (controller *Controller) RegisterRoutes(router *gin.Engine) {
	// The Routes of the user need to be added in here.
	router.GET("/v1/users", controller.GetUsers)
	router.GET("/v1/user/:userID", controller.GetUserById)
	router.PUT("/v1/user/:userID", controller.UpdateUser)
	router.DELETE("/v1/user/:userID", controller.DeleteUser)
}

func (controller *Controller) GetUsers(ctx *gin.Context) {

}
func (controller *Controller) GetUserById(ctx *gin.Context) {

}
func (controller *Controller) UpdateUser(ctx *gin.Context) {

}
func (controller *Controller) VerifyUser(ctx *gin.Context) {

}
func (controller *Controller) DeleteUser(ctx *gin.Context) {

}
