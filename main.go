package main

import (
	"github.com/gin-gonic/gin"
	"parking_to_esay/model"
	"parking_to_esay/router"
)

func main() {
	r := gin.Default()
	connection, err := model.ConnectDatabase()
	if err != nil {
		panic("Lỗi kết nối database !")
	}
	router.InitRouter(r, connection)
	r.Run(":8085")
}