package main

import (
	"DeviceConnect/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"fmt"

	"github.com/joho/godotenv"
	"github.com/okcredit/go-common/hack/auth"
)

func main() {

	err := godotenv.Load(".env")
	fmt.Println(err)
	if err != nil {
		fmt.Println("Cannot find .env file")
		return
	}
	serverPort := os.Getenv("PORT")
	serverPort = ":" + serverPort

	if serverPort == ":" {
		fmt.Println("Unable to find SERVER_PORT from environmental variables")
		return
	}

	// db,err := driver.Connect()

	// accountController := controller.AccountController{DB: db}
	// filterController := controller.FilterController{DB: db}

	fmt.Print(err)

	router := mux.NewRouter().StrictSlash(true)
	routes.HandleAccountRoutes(router)
	router.Use(auth.HttpMiddleware)
	log.Fatal(http.ListenAndServe(serverPort, cors.AllowAll().Handler(router)))
}
