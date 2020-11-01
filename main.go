package main

import (
	. "financial/apis"
	db "financial/database"

	"github.com/gin-gonic/gin"
)

func main() {
	defer db.Eloquent.Close()
	r := gin.Default()

	r.GET("/api/sector/list", SectorListAPI)

	r.Run()
}
