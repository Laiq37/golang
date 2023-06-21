package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controller.NewApplication(database.ProductData(database.client, "Products"), database.UserData(database.client, "Users"))
}