package main

import (
	"context"
	"fmt"
	"github.com/MatthewJamesBoyle/whattowearbot/weather"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var users = []weather.User{
	{
		ChatId:     27995890,
		Name:       "Matt",
		LocationId: 328328,
	},
	{
		ChatId:     27423818,
		Name:       "James",
		LocationId: 322309,
	},
	weather.User{
		ChatId:     582199789,
		Name:       "Davide",
		LocationId: 328328,
	},
}

//func main() {
//	errs := godotenv.Load()
//	if errs != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
//	if err != nil {
//		log.Panic(err)
//	}
//
//	service := weather.LocationService{
//		ApiKey: os.Getenv("API_KEY"),
//		ApiUrl: os.Getenv("API_URL"),
//	}
//
//	for _, user := range users {
//		resp, err := service.GetLocationData(user.LocationId)
//
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(resp)
//
//		bot.Debug = true
//
//		log.Printf("Authorized on account %s", bot.Self.UserName)
//
//		bot.Send(tgbotapi.NewMessage(user.ChatId, service.ShouldWear(user, resp)))
//	}
//}

//func main() {
//	lambda.Start(HandleRequest)
//}

func HandleRequest(ctx context.Context) (string, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	service := weather.LocationService{
		ApiKey: os.Getenv("API_KEY"),
		ApiUrl: os.Getenv("API_URL"),
	}

	for _, user := range users {
		resp, err := service.GetLocationData(user.LocationId)

		if err != nil {
			panic(err)
		}
		fmt.Println(resp)

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		bot.Send(tgbotapi.NewMessage(user.ChatId, service.ShouldWear(user, resp)))
	}
	return "success", nil
}

//func main() {
//	jsons := `{
//  "Headline": {
//    "EffectiveDate": "2018-07-27T14:00:00+01:00",
//    "EffectiveEpochDate": 1532696400,
//    "Severity": 1,
//    "Text": "Watch for a strong thunderstorm Friday afternoon; storms can cause flash flooding",
//    "Category": "thunderstorm",
//    "EndDate": "2018-07-27T20:00:00+01:00",
//    "EndEpochDate": 1532718000,
//    "MobileLink": "http://m.accuweather.com/en/gb/london/ec4a-2/extended-weather-forecast/328328?unit=c&lang=en-us",
//    "Link": "http://www.accuweather.com/en/gb/london/ec4a-2/daily-weather-forecast/328328?unit=c&lang=en-us"
//  },
//  "DailyForecasts": [
//    {
//      "Date": "2018-07-26T07:00:00+01:00",
//      "EpochDate": 1532584800,
//      "Sun": {
//        "Rise": "2018-07-26T05:15:00+01:00",
//        "EpochRise": 1532578500,
//        "Set": "2018-07-26T20:58:00+01:00",
//        "EpochSet": 1532635080
//      },
//      "Moon": {
//        "Rise": "2018-07-26T20:11:00+01:00",
//        "EpochRise": 1532632260,
//        "Set": "2018-07-27T04:45:00+01:00",
//        "EpochSet": 1532663100,
//        "Phase": "WaxingGibbous",
//        "Age": 13
//      },
//      "Temperature": {
//        "Minimum": {
//          "Value": 21.1,
//          "Unit": "C",
//          "UnitType": 17
//        },
//        "Maximum": {
//          "Value": 32.8,
//          "Unit": "C",
//          "UnitType": 17
//        }
//      },
//      "RealFeelTemperature": {
//        "Minimum": {
//          "Value": 20.6,
//          "Unit": "C",
//          "UnitType": 17
//        },
//        "Maximum": {
//          "Value": 32.9,
//          "Unit": "C",
//          "UnitType": 17
//        }
//      },
//      "RealFeelTemperatureShade": {
//        "Minimum": {
//          "Value": 20.6,
//          "Unit": "C",
//          "UnitType": 17
//        },
//        "Maximum": {
//          "Value": 31,
//          "Unit": "C",
//          "UnitType": 17
//        }
//      },
//      "HoursOfSun": 10.9,
//      "DegreeDaySummary": {
//        "Heating": {
//          "Value": 0,
//          "Unit": "C",
//          "UnitType": 17
//        },
//        "Cooling": {
//          "Value": 9,
//          "Unit": "C",
//          "UnitType": 17
//        }
//      },
//      "AirAndPollen": [
//        {
//          "Name": "AirQuality",
//          "Value": 75,
//          "Category": "Moderate",
//          "CategoryValue": 2,
//          "Type": "Ozone"
//        },
//        {
//          "Name": "Grass",
//          "Value": 0,
//          "Category": "Low",
//          "CategoryValue": 1
//        },
//        {
//          "Name": "Mold",
//          "Value": 0,
//          "Category": "Low",
//          "CategoryValue": 1
//        },
//        {
//          "Name": "Ragweed",
//          "Value": 0,
//          "Category": "Low",
//          "CategoryValue": 1
//        },
//        {
//          "Name": "Tree",
//          "Value": 0,
//          "Category": "Low",
//          "CategoryValue": 1
//        },
//        {
//          "Name": "UVIndex",
//          "Value": 7,
//          "Category": "High",
//          "CategoryValue": 3
//        }
//      ],
//      "Day": {
//        "Icon": 3,
//        "IconPhrase": "Partly sunny",
//        "ShortPhrase": "Partly sunny; hot",
//        "LongPhrase": "Partly sunny; hot",
//        "PrecipitationProbability": 13,
//        "ThunderstormProbability": 24,
//        "RainProbability": 13,
//        "SnowProbability": 0,
//        "IceProbability": 0,
//        "Wind": {
//          "Speed": {
//            "Value": 14.8,
//            "Unit": "km/h",
//            "UnitType": 7
//          },
//          "Direction": {
//            "Degrees": 161,
//            "Localized": "SSE",
//            "English": "SSE"
//          }
//        },
//        "WindGust": {
//          "Speed": {
//            "Value": 24.1,
//            "Unit": "km/h",
//            "UnitType": 7
//          },
//          "Direction": {
//            "Degrees": 191,
//            "Localized": "S",
//            "English": "S"
//          }
//        },
//        "TotalLiquid": {
//          "Value": 0,
//          "Unit": "mm",
//          "UnitType": 3
//        },
//        "Rain": {
//          "Value": 0,
//          "Unit": "mm",
//          "UnitType": 3
//        },
//        "Snow": {
//          "Value": 0,
//          "Unit": "cm",
//          "UnitType": 4
//        },
//        "Ice": {
//          "Value": 0,
//          "Unit": "mm",
//          "UnitType": 3
//        },
//        "HoursOfPrecipitation": 0,
//        "HoursOfRain": 0,
//        "HoursOfSnow": 0,
//        "HoursOfIce": 0,
//        "CloudCover": 29
//      },
//      "Night": {
//        "Icon": 41,
//        "IconPhrase": "Partly cloudy w/ t-storms",
//        "ShortPhrase": "A thunderstorm in the area",
//        "LongPhrase": "Partly cloudy, a shower or thunderstorm in spots; very warm",
//        "PrecipitationProbability": 42,
//        "ThunderstormProbability": 60,
//        "RainProbability": 42,
//        "SnowProbability": 0,
//        "IceProbability": 0,
//        "Wind": {
//          "Speed": {
//            "Value": 11.1,
//            "Unit": "km/h",
//            "UnitType": 7
//          },
//          "Direction": {
//            "Degrees": 267,
//            "Localized": "W",
//            "English": "W"
//          }
//        },
//        "WindGust": {
//          "Speed": {
//            "Value": 20.4,
//            "Unit": "km/h",
//            "UnitType": 7
//          },
//          "Direction": {
//            "Degrees": 203,
//            "Localized": "SSW",
//            "English": "SSW"
//          }
//        },
//        "TotalLiquid": {
//          "Value": 3,
//          "Unit": "mm",
//          "UnitType": 3
//        },
//        "Rain": {
//          "Value": 3,
//          "Unit": "mm",
//          "UnitType": 3
//        },
//        "Snow": {
//          "Value": 0,
//          "Unit": "cm",
//          "UnitType": 4
//        },
//        "Ice": {
//          "Value": 0,
//          "Unit": "mm",
//          "UnitType": 3
//        },
//        "HoursOfPrecipitation": 2,
//        "HoursOfRain": 2,
//        "HoursOfSnow": 0,
//        "HoursOfIce": 0,
//        "CloudCover": 66
//      },
//      "Sources": [
//        "AccuWeather"
//      ],
//      "MobileLink": "http://m.accuweather.com/en/gb/london/ec4a-2/daily-weather-forecast/328328?day=1&unit=c&lang=en-us",
//      "Link": "http://www.accuweather.com/en/gb/london/ec4a-2/daily-weather-forecast/328328?day=1&unit=c&lang=en-us"
//    }
//  ]
//}`
//
//	var result map[string]interface{}
//	json.Unmarshal([]byte(jsons), &result)
//	dailyForecast := result["DailyForecasts"].([]interface{})[0]
//	summary := dailyForecast.(map[string]interface{})
//	temperature := summary["Temperature"]
//
//	minTemp := temperature.(map[string]interface{})["Minimum"]
//	maxTemp := temperature.(map[string]interface{})["Maximum"]
//
//	minTempValue := minTemp.(map[string]interface{})["Value"]
//	maxTempValue := maxTemp.(map[string]interface{})["Value"]
//	sunSet := dailyForecast.(map[string]interface{})["Sun"].(map[string]interface{})["Set"]
//
//	weatherDescription := summary["Day"].(map[string]interface{})["LongPhrase"]
//	fmt.Println(minTempValue, maxTempValue, weatherDescription, sunSet)
//}
