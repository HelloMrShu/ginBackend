package main

import (
	"financial/components"
	"financial/routers"
)

func main() {
	components.LogToFile()
	r := routers.InitRouter()

	r.Run()
}
