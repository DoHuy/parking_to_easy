package router

import (
	"github.com/DoHuy/parking_to_easy/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, con *controller.Controller) {
	// Tạo mới bãi đẫu xe
	router.POST("/api/admin/create/parking", con.CreateNewParkingByAdmin) //done
	router.POST("/api/user/share/parking", con.CreateNewParkingByOwner) //done
	// Lấy thông tin bãi đậu xe
	router.GET("/api/get/all/approved/parkings", con.GetAllApprovedParkings) //done
	router.GET("/api/admin/get/all/parkings/:limit/:offset", con.GetAllParkings) //done
	router.GET("/api/owner/get/all/parkings", con.GetAllParkingsOfOwner)//done con phai sua
	router.GET("/api/get/parking/:parkingId", con.FindParkingByID) //done
	router.GET("/api/recommend/parking/radius/:rad", con.RecommendParking)
	// Cập nhật thông tin bãi xe
	router.PATCH("/api/admin/verify/parking/:id", con.VerifyParking) // done
	router.PUT("/api/owner/modify/parking/:id", con.ModifyParkingByOwner) // chua test
	router.DELETE("/api/owner/remove/parking/:id", con.RemoveParkingOfOwner) //chua test
	router.GET("/api/calculate/amount/parking/:id", con.CalculateAmountParking)//chua test
	// Login
	router.POST("/api/login", con.Login)//done
	// credential
	router.POST("/api/register", con.Register) // done
	router.GET("/api/get/all/users/:limit/:offset", con.GetAllUsers) // done
	router.GET("/api/get/detail/profile", con.GetDetailUser)  // done

	//owner
	router.GET("/api/admin/get/all/owners/:limit/:offset", con.GetAllOwners) // done
	router.GET("/api/get/owner/:id", con.GetOwnerById)//done
	router.POST("/api/create/owner", con.CreateNewOwner)//done
	router.PUT("/api/admin/disable/owner/:id", con.DisableOwner) // chua test
	// transaction
	router.POST("/api/user/create/transaction", con.CreateNewTransaction)
	router.GET("/api/user/get/all/transaction", con.GetAllTransactionOfUser)
	router.GET("/api/admin/get/all/transaction", con.GetAllTransaction)
	router.PATCH("/api/decline/transaction/:id", con.DeclineTransaction)
	router.PATCH("/api/accept/transaction/:id", con.AcceptTransaction)
	////////////////////////
	// Upload nhieu file
	router.POST("/api/files/upload", con.UploadFiles) // done
	/// rating
	router.POST("/api/rating/parking", con.RatingParking)
	router.GET("/analysis/metric/all/parkings", con.AnalysisAllParkings)
	return
}