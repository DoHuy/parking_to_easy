package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/router"
	"github.com/gin-contrib/static"
	"path/filepath"
)

func main() {
	r := gin.Default()
	static_path, _ := filepath.Abs("./")
	r.Use(static.Serve("/", static.LocalFile(filepath.Join(static_path, "resource/images"), false)))
	connection, err := model.ConnectDatabase()
	if err != nil {
		panic("Lỗi kết nối database !")
	}
	router.InitRouter(r, connection)
	r.Run(":8085")
}