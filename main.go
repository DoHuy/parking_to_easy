package main

import (
	"github.com/DoHuy/parking_to_easy/business_logic"
	"github.com/DoHuy/parking_to_easy/http"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	r := gin.Default()
	static_path, _ := filepath.Abs("./")
	r.Use(static.Serve("/", static.LocalFile(filepath.Join(static_path, "resource/images"), false)))
	factoryService, err := business_logic.NewFactory()
	if err != nil {
		panic(err)
	}
	middleware := http.NewMiddleware(factoryService)
	controller:= http.NewControllingService(factoryService, middleware)
	routingService := http.NewService(r, controller)
	routingService.Init()
	r.Run(":8085")
}
