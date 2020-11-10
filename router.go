package main

import (
	. "financial/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/sector/list", SectorListAPI)
	router.POST("/sector/save", SectorSaveAPI)
	router.POST("/sector/delete", SectorDeleteAPI)
	return router
}
