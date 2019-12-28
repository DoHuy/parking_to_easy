package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/router"
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