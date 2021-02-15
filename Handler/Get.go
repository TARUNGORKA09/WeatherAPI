package Handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/TARUNGORKA09/WeatherAPI/weatherData"
)

type WeatherInfo struct {
	l *log.Logger
}

func NewWeather(m *log.Logger) *WeatherInfo {
	return &WeatherInfo{m}
}

func (data WeatherInfo) GetWeather(rw http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	long, lati, err := parser(r)
	lon := fmt.Sprintf("%g", long)
	lat := fmt.Sprintf("%g", lati)
	req, err := url.Parse("https://api.openweathermap.org/data/2.5/weather")
	q := url.Values{}
	q.Add("lat", lat)
	q.Add("lon", lon)
	q.Add("appid", "cec178f08a0a71980085edd7163c6e29")
	req.RawQuery = q.Encode()
	if err != nil {
		fmt.Fprintf(rw, " ", err)
	}
	res, err := http.Get(req.String())
	defer res.Body.Close()
	var data1 weatherData.WeatherData
	err = json.NewDecoder(res.Body).Decode(&data1)
	if err != nil {
		fmt.Println("unable to decode full Data")
	}
	fmt.Fprintln(rw, "Weather :", data1.Weather)
	fmt.Fprintln(rw, "Base :", data1.Base)
	fmt.Fprintln(rw, "Name:", data1.Name)
	fmt.Fprintln(rw, "TimeZone :", data1.Timezone)
	fmt.Fprintln(rw, "ID :", data1.ID)
}

func parser(r *http.Request) (float64, float64, error) {

	var location weatherData.Location
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		fmt.Printf("Invalid Location")
	}
	long := location.Long
	lati := location.Lat

	return long, lati, nil
}
