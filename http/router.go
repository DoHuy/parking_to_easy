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
	service.Router.PUT("/api/owner/modify/parking/:id", service.Controller.ModifyParkingByOwner) // done
	service.Router.DELETE("/api/owner/remove/parking/:id", service.Controller.RemoveParkingOfOwner) //chua test
	service.Router.GET("/api/calculate/amount/parking/:id", service.Controller.CalculateAmountParking)//done
	// Login
	service.Router.POST("/api/login", service.Controller.Login)//done
	// credential
	service.Router.POST("/api/register", service.Controller.Register) // done
	service.Router.GET("/api/get/all/users/:limit/:offset", service.Controller.GetAllUsers) // done
	service.Router.GET("/api/get/detail/profile", service.Controller.GetDetailUser)  // done

	//owner
	service.Router.GET("/api/admin/get/all/owners", service.Controller.GetAllOwners) // done
	service.Router.GET("/api/get/owner/:id", service.Controller.GetOwnerById)//done
	service.Router.POST("/api/create/owner", service.Controller.CreateNewOwner)//done
	service.Router.PATCH("/api/admin/disable/owner/:id", service.Controller.DisableOwner) // done
	// transaction
	service.Router.POST("/api/user/create/transaction", service.Controller.CreateNewTransaction) // done / chua them goroutine va firebase
	service.Router.GET("/api/owner/get/transactions/parking/:parkingId/state/:status", service.Controller.GetAllTransactionOfOwner) // done
	service.Router.GET("/api/user/get/all/transaction/:status", service.Controller.GetTransactionOfUser) // done
	service.Router.GET("/api/admin/get/all/transactions", service.Controller.GetAllTransaction) // done
	service.Router.PATCH("/api/change/transaction", service.Controller.ChangeStateTransaction)// done // can phai gui thong bao qua firebase
	//service.Router.PATCH("/api/accept/transaction/:id", service.Controller.AcceptTransaction)
	//service.Router.PATCH("/api/process/transaction/:id", service.Controller.ProcessTransaction)//
	//service.Router.PATCH("/api/finish/transaction/:id", service.Controller.FinishTransaction)
	////////////////////////
	// Upload nhieu file
	service.Router.POST("/api/files/upload", service.Controller.UploadFiles) // done
	/// rating
	service.Router.POST("/api/rating/parking", service.Controller.RatingParking)// done
	service.Router.GET("/analysis/metric/all/transactions/:start/to/:end", service.Controller.AnalysisTransaction)// done

	// token fire base
	service.Router.POST("/api/save/token/firebase", service.Controller.SaveTokenFireBase)// done
	service.Router.DELETE("/api/remove/token/firebase", service.Controller.RemoveToken)// done
	//
	return
}
