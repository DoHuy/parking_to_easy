package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/DoHuy/parking_to_easy/controller"
	"github.com/DoHuy/parking_to_easy/middleware"
)



func InitRouter(router *gin.Engine, conn *gorm.DB) {
	// khoi tao middleware
	middleware := middleware.NewMiddleware(conn)
	//khoi tao conntroller
	controller := controller.NewController(conn)
	// Tạo mới bãi đẫu xe
	apiParking := router.Group("/api")
	apiParking.GET("/parkings/:parkingId", controller.FindParkingByID)
	//apiParking.GET("/parkings/all/:limit/:offset", controller.GetAllParkings)
	////////////////////////
	// Upload nhieu file
	router.POST("/api/files/upload", middleware.BeforeUploadFiles, controller.UploadFiles)
	return
}