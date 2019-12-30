package router

import (
	"github.com/DoHuy/parking_to_easy/controller"
	"github.com/DoHuy/parking_to_easy/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)



func InitRouter(router *gin.Engine, conn *gorm.DB) {
	// khoi tao middleware
	middleware := middleware.NewMiddleware(conn)
	//khoi tao conntroller
	controller := controller.NewController(conn)
	// Tạo mới bãi đẫu xe
	router.GET("/api/get/parking/:parkingId", controller.FindParkingByID)
	router.GET("/api/get/all/parkings/:limit/:offset", controller.GetAllParkings)
	router.POST("/api/user/create/parking", controller.CreateNewParking)
	router.POST("/api/admin/create/parking", controller.CreateNewParkingByAdmin)
	// Credential
	router.POST("/api/register", controller.CreateNewCredential)
	//router.POST("/api/login", controller.Login)
	//

	////////////////////////

	////////////////////////
	// Upload nhieu file
	//router.Use(middleware.BeforeUploadFiles())
	router.POST("/api/files/upload",middleware.BeforeUploadFiles, controller.UploadFiles)
	return
}