package main

import (
	"fmt"
	"github.com/MatthewJamesBoyle/whattowearbot/weather"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var users = []weather.User{
	{
		ChatId:   27995890,
		Name:     "Matt",
		Location: "London,uk",
	},
	{
		ChatId:   27423818,
		Name:     "James",
		Location: "Cornwall,uk",
	},
	{
		ChatId:   582199789,
		Name:     "Davide",
		Location: "london,uk",
	},
}

func main() {
	errs := godotenv.Load()
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	service := weather.LocationService{
		ApiKey: os.Getenv("API_KEY"),
		ApiUrl: os.Getenv("API_URL"),
	}

	for _,user := range users {
		resp, err := service.GetLocationData(user.Location)

		if err != nil {
			panic(err)
		}
		fmt.Println(resp)

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		bot.Send(tgbotapi.NewMessage(user.ChatId, service.ShouldWear(user,resp)))
	}
}

//
//func main() {
//	jsons := `{
//"coord": {
//"lon": -0.13,
//"lat": 51.51
//},
//"weather": [
//{
//"id": 802,
//"main": "Clouds",
//"description": "scattered clouds",
//"icon": "03n"
//}
//],
//"base": "stations",
//"main": {
//"temp": 21.81,
//"pressure": 1017,
//"humidity": 53,
//"temp_min": 20,
//"temp_max": 23
//},
//"visibility": 10000,
//"wind": {
//"speed": 5.1,
//"deg": 90
//},
//"clouds": {
//"all": 44
//},
//"dt": 1532031600,
//"sys": {
//"type": 1,
//"id": 5091,
//"message": 0.0025,
//"country": "GB",
//"sunrise": 1531973211,
//"sunset": 1532030766
//},
//"id": 2643743,
//"name": "London",
//"cod": 200
//}`
//
//	var result map[string]interface{}
//	json.Unmarshal([]byte(jsons), &result)
//	 main := result["weather"].([]interface{})[0]
//	 summary := main.(map[string]interface{})
//	 fmt.Println(summary)
//
//	}
