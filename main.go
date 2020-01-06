package main

import (
	auth2 "github.com/DoHuy/parking_to_easy/auth"
	controller2 "github.com/DoHuy/parking_to_easy/controller"
	middleware2 "github.com/DoHuy/parking_to_easy/middleware"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/router"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	r := gin.Default()
	static_path, _ := filepath.Abs("./")
	r.Use(static.Serve("/", static.LocalFile(filepath.Join(static_path, "resource/images"), false)))
	auth := auth2.NewAuth()
	dao, err := mysql.NewDAO()
	if err != nil {
		panic(err.Error())
	}
	middleware := middleware2.NewMiddleware(dao, auth)
	controller:= controller2.NewController(dao, middleware, auth)
	router.InitRouter(r, controller)
	r.Run(":8085")
}