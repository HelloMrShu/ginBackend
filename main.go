package main

import (
	db "financial/database"
)

func main() {
	defer db.Eloquent.Close()
	router := initRouter()
	router.Run()
}
