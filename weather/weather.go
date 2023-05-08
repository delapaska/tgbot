package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiKey = "0684c950bf2d40880f862810f27ae6ae"

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}

func TakeTempMoscow() float64 {
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=0684c950bf2d40880f862810f27ae6ae", "moscow"))
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}
	s := weather.Main.Temp
	return s
}

func TakeWindSpeedMoscow() float64 {
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=0684c950bf2d40880f862810f27ae6ae", "moscow"))
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}
	s := weather.Wind.Speed
	return s
}

func TakeTempSPB() float64 {
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=0684c950bf2d40880f862810f27ae6ae", "Saint Petersburg"))
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}
	s := weather.Main.Temp
	return s
}

func TakeWindSpeedSPB() float64 {
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=0684c950bf2d40880f862810f27ae6ae", "Saint Petersburg"))
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}
	s := weather.Wind.Speed
	return s
}
