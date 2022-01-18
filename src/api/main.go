package main

import (
	"fmt"
	"log"
	"myapp/src/appcfg"
	"myapp/src/handler/middleware"
	"myapp/src/handler/webservice"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// init logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	if err = godotenv.Load(".env"); err != nil {
		log.Panic("error load env")
	}
	appcfg.Setup()
}

func main() {
	r := gin.Default()
	createRouter(r)
	r.Run(fmt.Sprintf(":%s", appcfg.ServicePort))
}

func createRouter(r *gin.Engine) {
	mainAPI := r.Group("/api")
	mainAPI.GET("/covid-stat", middleware.DIMiddleware(), webservice.APICovidHandler)
}
