package main

import (
	"pharmacy-api/database"
	"pharmacy-api/prescription"
	"pharmacy-api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.InitDB()

	// init user
	userService := user.UserServiceDependency{Db: db}
	userController := user.Controller{
		UserService: userService,
	}
	userController.RegisterRoutes(router)

	// init prescription
	prescriptionService := prescription.PrescriptionServiceDependency{Db: db}
	prescriptionController := prescription.Controller{
		PrescriptionService: prescriptionService,
	}
	prescriptionController.RegisterRoutes(router)

	router.Run()
}
