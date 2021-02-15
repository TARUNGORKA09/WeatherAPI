package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://api.openweathermap.org/data/2.5/weather?lat=35&lon=139&appid=cec178f08a0a71980085edd7163c6e29"

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to get Weather data")
	}
	defer res.Body.Close()

}
