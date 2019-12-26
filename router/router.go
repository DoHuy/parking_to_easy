package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"parking_to_esay/controller"
	"parking_to_esay/middleware"
)



func InitRouter(router *gin.Engine, conn *gorm.DB) {

	// khoi tao middleware
	middleware := middleware.NewMiddleware(conn)
	//khoi tao conntroller
	controller := controller.NewController(conn)
	// Tạo mới bãi đẫu xe
	router.POST("/api/parking", middleware.ValidateParkingCreating, controller.CreateNewParking)
	////////////////////////
	return
}