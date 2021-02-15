package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TARUNGORKA09/WeatherAPI/Handler"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server is starting ................")

	l := log.New(os.Stdout, "Mobile Todo", log.LstdFlags)

	mobile := Handler.NewWeather(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getWeather", mobile.GetWeather)

	http.ListenAndServe(":8080", sm)

}
