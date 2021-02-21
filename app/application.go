package app

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())
	mapUrls()
	log.Fatalln(router.Run(":8080"))
}
