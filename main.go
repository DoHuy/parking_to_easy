package main

import (
	auth2 "github.com/DoHuy/parking_to_easy/business_logic/auth"
	"github.com/DoHuy/parking_to_easy/http"
	"github.com/DoHuy/parking_to_easy/mysql"
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
	middleware := http.NewMiddleware(dao, auth)
	controller:= http.NewControllingService(dao, middleware, auth)
	routingService := http.NewService(r, controller)
	routingService.Init()
	r.Run(":8085")
}