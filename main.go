package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
	"github.com/TARUNGORKA09/WeatherAPI/weatherData"
	"github.com/TARUNGORKA09/WeatherAPI/Handler"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server is starting ................")

	l := log.New(os.Stdout, "Mobile Todo", log.LstdFlags)

	mobile := Handlers.NewWeather(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getMobile/{id:[0-9]+}", mobile.GetMobileInfo)

	http.ListenAndServe(":8080", sm)

}