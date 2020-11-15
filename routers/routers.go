package routers

import (
	. "financial/apis"
	"financial/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() * gin.Engine {
	r := gin.Default()

	r.Use(middlewares.RequestLog())

	r.GET("/sector/list", SectorListAPI)
	r.POST("/sector/save", SectorSaveAPI)
	r.POST("/sector/delete", SectorDeleteAPI)
	return r
}
