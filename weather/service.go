package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	ChatId   int64
	Name     string
	Location string
}

type LocationService struct {
	ApiUrl string
	ApiKey string
}

// WeatherIds from here:
// https://openweathermap.org/weather-conditions
type LocationResponse struct {
	CurrentTemp        float64
	LowTemp            float64
	HighTemp           float64
	WeatherDescription string
	WeatherId          float64
}

type Weather struct {
	main        string
	description string
}

func (locationService LocationService) GetLocationData(city string) (*LocationResponse, error) {
	reqUrl := locationService.ApiUrl + city + "&units=metric" + "&APPID=" + locationService.ApiKey
	r, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err

	}
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)
	summary := result["weather"].([]interface{})[0].(map[string]interface{})
	main := result["main"].(map[string]interface{})
	fmt.Println(summary)
	fmt.Println(main)

	return &LocationResponse{
		CurrentTemp:        main["temp"].(float64),
		LowTemp:            main["temp_min"].(float64),
		HighTemp:           main["temp_max"].(float64),
		WeatherDescription: summary["description"].(string),
		WeatherId:          summary["id"].(float64),
	}, nil
}

func (locationService LocationService) ShouldWear(user User, lr *LocationResponse) string {
	response := "Morning " + user.Name + ". Looks like the weather today in " + user.Location + " is " + lr.WeatherDescription + "."
	response += "The current Temperature is " + FloatToString(lr.CurrentTemp) + "."
	response += "There will be a high of " + FloatToString(lr.HighTemp) + " and a low of " + FloatToString(lr.LowTemp) + "."
	response += getRecommendation(lr)
	return response
}

func getRecommendation(lr *LocationResponse) string {
	recommendation := ""
	if lr.CurrentTemp < 10 && lr.HighTemp > 10 {
		recommendation = "You will need a coat or jumper this morning, but it will warm up later. "
	} else if lr.CurrentTemp > 10 && lr.HighTemp <= 15 {
		recommendation = "You should be fine with shorts today, but you might want to take a jumper too."
	} else if lr.CurrentTemp > 10 && lr.HighTemp > 15 {
		recommendation = "You should be fine with shorts today and you won't need a jumper."
	} else {
		recommendation = "You should wear the standard jeans and a jumper."
	}
	if lr.WeatherId < 800 || lr.CurrentTemp < 10 || lr.HighTemp < 10 {
		recommendation += "You will need a coat."
	}
	return recommendation

}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 0, 64)

}
