package main

import (
	"fmt"
	"rest-api-practice/controllers"
	"rest-api-practice/database"
	"rest-api-practice/routers"
)

func main() {
	db, err := database.InitializeDB()
	if err != nil {
		fmt.Println("Aborting operation due to:", err)
		return
	}
	fmt.Println("connected to db!")

	ctr := controllers.InitializeNewControllerWithDB(db)

	r := routers.InitializeRouter(":8080", ctr)
	r.Run()
}
