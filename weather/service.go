package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ChatId     int64
	Name       string
	LocationId int64
}

type LocationService struct {
	ApiUrl string
	ApiKey string
}

// WeatherIds from here:
// https://openweathermap.org/weather-conditions
type LocationResponse struct {
	LowTemp            float64
	HighTemp           float64
	WeatherDescription string
	SunSet             string
}

func (locationService LocationService) GetLocationData(locationId int64) (*LocationResponse, error) {
	reqUrl := locationService.ApiUrl + strconv.FormatInt(locationId, 10) + "?apikey=" + locationService.ApiKey + "&details=true&metric=true"
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
	fmt.Println(result)

	dailyForecast := result["DailyForecasts"].([]interface{})[0]
	summary := dailyForecast.(map[string]interface{})
	temperature := summary["Temperature"]

	minTemp := temperature.(map[string]interface{})["Minimum"]
	maxTemp := temperature.(map[string]interface{})["Maximum"]

	minTempValue := minTemp.(map[string]interface{})["Value"]
	maxTempValue := maxTemp.(map[string]interface{})["Value"]

	weatherDescription := summary["Day"].(map[string]interface{})["LongPhrase"]
	sunSet := dailyForecast.(map[string]interface{})["Sun"].(map[string]interface{})["Set"]

	return &LocationResponse{
		LowTemp:            minTempValue.(float64),
		HighTemp:           maxTempValue.(float64),
		WeatherDescription: weatherDescription.(string),
		SunSet:             sunSet.(string),
	}, nil
}

func (locationService LocationService) ShouldWear(user User, lr *LocationResponse) string {
	response := "Morning " + user.Name + ". Looks like the weather today is " + lr.WeatherDescription + "."
	response += "Today there will be a high of " + FloatToString(lr.HighTemp) + " and a low of " + FloatToString(lr.LowTemp) + "."
	response += getRecommendation(lr)

	t1, _ := time.Parse(time.RFC3339, lr.SunSet)
	response += "Sunset will be at " + t1.Format("3:04PM") + "."
	return response
}

func getRecommendation(lr *LocationResponse) string {
	recommendation := ""
	if lr.HighTemp <= 10 {
		recommendation = "You should wear a jumper and jeans today. "
	} else if lr.HighTemp <= 15 {
		recommendation = "You should be fine with shorts today, but might want a jumper.."
	} else if lr.HighTemp > 15 {
		recommendation = "You should be fine with shorts today and you won't need a jumper."
	} else {
		recommendation = "You should wear the standard jeans and a jumper."
	}
	if strings.Contains(lr.WeatherDescription, "rain") {
		recommendation += "It is going to rain at some point today so you should take a coat or Umbrella."
	}
	return recommendation

}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 0, 64)

}
