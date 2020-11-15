package main

import (
	"financial/routers"
)

func main() {
	r := routers.InitRouter()
	//r := gin.New()
	//r.Use(costTime())

	r.Run()
}
