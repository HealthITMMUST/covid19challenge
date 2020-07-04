package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)

	}
	token = os.Getenv("TELEGRAM_TOKEN")
	//export TELEGRAM_TOKEN=1224895113:AAFdB8-X_Iav2z6v0UkFMumAxeQKeDSuCJI

	//token = "1224895113:AAFdB8-X_Iav2z6v0UkFMumAxeQKeDSuCJI"
}

func main() {
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/symptoms", app.symptoms)

	log.Fatal(bot.Start())
}

func (a *application) startHandler(m *tbot.Message) {
	msg := "this is a bot whose purpose is to Serve COVID-19 FACTS"
	a.client.SendMessage(m.Chat.ID, msg)
}

func (a *application) symptoms(m *tbot.Message) {

	msg := "Signs and symptoms of coronavirus disease 2019 (COVID-19) may appear two to 14 days after exposure. This time after exposure and before having symptoms is called the incubation period.%0A Common signs and symptoms can include:%0A Fever Cough"
	a.client.SendMessage(m.Chat.ID, msg)
}
