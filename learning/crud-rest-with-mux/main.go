package main

import (
	"crud-rest-with-mux/controllers"
	"crud-rest-with-mux/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Load configurations from config.json using viper
	LoadAppConfig()

	//Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	//Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	//RegisterRoutes
	RegisterProductRoutes(router)

	//Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
