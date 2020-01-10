package http

import (
	"github.com/gin-gonic/gin"
)

type RouteService struct {
	Router		*gin.Engine
	Controller	*ControllingService
}
func NewService(router *gin.Engine, con *ControllingService) *RouteService{
	return &RouteService{
		Router:		 router,
		Controller:  con,
	}
}
func (service *RouteService)Init() {
	service.Router.Use(service.Controller.Options)
	// Tạo mới bãi đẫu xe
	service.Router.POST("/api/admin/create/parking", service.Controller.CreateNewParkingByAdmin) //done
	service.Router.POST("/api/user/share/parking", service.Controller.CreateNewParkingByOwner) //done
	// Lấy thông tin bãi đậu xe
	service.Router.GET("/api/get/all/approved/parkings", service.Controller.GetAllApprovedParkings) //done
	service.Router.GET("/api/admin/get/all/parkings/:limit/:offset", service.Controller.GetAllParkings) //done
	service.Router.GET("/api/owner/get/all/parkings", service.Controller.GetAllParkingsOfOwner)//done con phai sua
	service.Router.GET("/api/get/parking/:parkingId", service.Controller.FindParkingByID) //done
	service.Router.GET("/api/recommend/parking/radius/:rad", service.Controller.RecommendParking)
	// Cập nhật thông tin bãi xe
	service.Router.PATCH("/api/admin/verify/parking/:id", service.Controller.VerifyParking) // done
	service.Router.PUT("/api/owner/modify/parking/:id", service.Controller.ModifyParkingByOwner) // chua test
	service.Router.DELETE("/api/owner/remove/parking/:id", service.Controller.RemoveParkingOfOwner) //chua test
	service.Router.GET("/api/calculate/amount/parking/:id", service.Controller.CalculateAmountParking)//chua test
	// Login
	service.Router.POST("/api/login", service.Controller.Login)//done
	// credential
	service.Router.POST("/api/register", service.Controller.Register) // done
	service.Router.GET("/api/get/all/users/:limit/:offset", service.Controller.GetAllUsers) // done
	service.Router.GET("/api/get/detail/profile", service.Controller.GetDetailUser)  // done

	//owner
	service.Router.GET("/api/admin/get/all/owners/:limit/:offset", service.Controller.GetAllOwners) // done
	service.Router.GET("/api/get/owner/:id", service.Controller.GetOwnerById)//done
	service.Router.POST("/api/create/owner", service.Controller.CreateNewOwner)//done
	service.Router.PUT("/api/admin/disable/owner/:id", service.Controller.DisableOwner) // chua test
	// transaction
	service.Router.POST("/api/user/create/transaction", service.Controller.CreateNewTransaction) // chua test
	service.Router.GET("/api/owner/get/all/transaction/:status", service.Controller.GetAllTransactionOfOwner) // chua test
	service.Router.GET("/api/user/get/transaction/:status", service.Controller.GetTransactionOfUser) // chua test
	service.Router.GET("/api/admin/get/all/transaction", service.Controller.GetAllTransaction) // chua test
	service.Router.PATCH("/api/decline/transaction/:id", service.Controller.DeclineTransaction)//
	service.Router.PATCH("/api/accept/transaction/:id", service.Controller.AcceptTransaction)
	////////////////////////
	// Upload nhieu file
	service.Router.POST("/api/files/upload", service.Controller.UploadFiles) // done
	/// rating
	service.Router.POST("/api/rating/parking", service.Controller.RatingParking)// chua tét
	//router.GET("/analysis/metric/all/parkings/:start/to/:end", con.AnalysisAllParkings) //

	// token fire base
	service.Router.POST("/api/save/token/firebase", service.Controller.SaveTokenFireBase)
	service.Router.DELETE("/api/remove/token/firebase", service.Controller.RemoveToken)
	//
	return
}
