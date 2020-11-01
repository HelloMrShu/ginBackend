package main

import (
	. "financial/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/sector/list", SectorListAPI)
	router.POST("/api/sector/save", SectorSaveAPI)
	router.POST("/api/sector/delete", SectorDeleteAPI)
	return router
}
