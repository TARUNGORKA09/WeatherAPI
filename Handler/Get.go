package Handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TARUNGORKA09/WeatherAPI/weatherData"
)

type WeatherInfo struct {
	l *log.Logger
}

func NewWeather(m *log.Logger) *WeatherInfo {
	return &WeatherInfo{m}
}

func (data WeatherInfo) GetWeather(rw http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var location weatherData.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		fmt.Printf("Invalid Location")
	}
	long := location.Long
	lati := location.Lat
	lon := fmt.Sprintf("%f", long)
	lat := fmt.Sprintf("%s", lati)
	res, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/weather", nil)
	q := res.URL.Query()
	q.Add("lat", lat)
	q.Add("lon", lon)
	q.Add("appid", "cec178f08a0a71980085edd7163c6e29")
	if err != nil {
		fmt.Fprintf(rw, " ", err)
	}
	defer res.Body.Close()
	var data1 weatherData.WeatherData
	err = json.NewDecoder(res.Body).Decode(&data1)
	if err != nil {
		fmt.Println("unable to decode full Data")
	}
	fmt.Fprintf(rw, "Weather :", data1.Weather)
	fmt.Fprintf(rw, "Base :", data1.Base)
	fmt.Fprintf(rw, "TimeZone :", data1.Timezone)
}
