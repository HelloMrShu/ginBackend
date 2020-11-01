package main

// . "financial/database"

func main() {
	// defer DB.Close()
	router := initRouter()
	router.Run()
}
