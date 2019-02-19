package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"opentribe/controllers"
	"os"
)

func main() {

	router := mux.NewRouter()

	//router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/signup", controllers.Signup)
	router.HandleFunc("/profile", controllers.Profile)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./html"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
