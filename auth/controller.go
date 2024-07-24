package auth

import "github.com/gin-gonic/gin"

type Controller struct {
	AuthService AuthService
}

func (controller *Controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/login", controller.LoginUser)
	router.POST("/register", controller.RegisterUser)
	router.POST("/logout", controller.LogOutUser)
	router.POST("/refresh-token", controller.RefreshToken)
}

func (controller *Controller) LoginUser(ctx *gin.Context) {

}
func (controller *Controller) RegisterUser(ctx *gin.Context) {

}
func (controller *Controller) LogOutUser(ctx *gin.Context) {

}
func (controller *Controller) RefreshToken(ctx *gin.Context) {

}
