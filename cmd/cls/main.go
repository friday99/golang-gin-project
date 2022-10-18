package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oranzh/cls/internal/middlewares"
	"github.com/oranzh/cls/internal/routers"
	"github.com/oranzh/cls/pkg/db"
	"github.com/oranzh/cls/pkg/router"
	"github.com/oranzh/cls/pkg/setting"
	"log"
)

func main() {
	// init setting
	err := setting.New()
	if err != nil {
		log.Fatalf("setting.New err: %v", err)
	}
	//db connect
	db.Setup()
	defer db.CloseDB()
	run()
}

// run Gin server
func run() {
	r := gin.Default()
	// global middleware
	r.Use(middlewares.RequestIDMiddleware)

	// load routers
	router.Include(routers.LoadWeb, routers.LoadApi)
	router.Init(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("r.Run err: %v", err)
	}
}
