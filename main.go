package main

import (
	"busuanzi/app/controller"
	"busuanzi/app/middleware"
	"busuanzi/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// debug
	if !config.C.Web.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// middleware
	if config.C.Web.Log {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	r.Use(middleware.AccessControl())

	// web
	r.LoadHTMLFiles("dist/index.html")
	r.StaticFile("/js", "dist/busuanzi.js")

	// router
	r.GET("/", controller.Index)
	r.POST("/api", controller.ApiHandler)
	r.OPTIONS("/api", controller.ApiHandler)
	r.GET("/ping", controller.PingHandler)
	r.NoRoute(controller.NoRouteHandler)

	// start server
	log.SetOutput(gin.DefaultWriter)
	log.Println("server listen on port:", config.C.Web.Address)
	err := r.Run(config.C.Web.Address)
	if err != nil {
		log.Fatalf("web服务启动失败: %s", err)
	}
}
