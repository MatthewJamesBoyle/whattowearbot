# What To Wear Bot
This is a Golang application that is intended to ping a list of recipients every morning to tell them what to wear based on the weather and their location. I am setting out to automate as many of the repetitve steps in my life as possible and checking the weather is one of the first things I do when I wake up in the morning. I also hope it is a good example app for others wanting to make a telegram bot.

## Getting it running
To get this app running, you'll need to have a telegram account. After that, message @thebotfather on telegram who will guide you through creating a bot.

You'll need to create a `.env` file in the route of the project with the credentials you just created plus the api url:
```
API_TOKEN=""
API_KEY=""
API_URL="hhttp://dataservice.accuweather.com/forecasts/v1/daily/1day/"
```

I use AccuWeather. You may need to update the parse method if you wish to use another service.

after that, just message your bot on telegram to get the conversation id in your go console. Paste this into the userslice (examples incldued in `main.go`) and the app should work!

I have this hosted on AWS Lambda using a Cloudwatch cron Trigger to run at 5:30am every morning UTC. This is a really effective (and free) way to host this. Included is a main method to run locally that you can use to test or use as a basis to use a different provider.
