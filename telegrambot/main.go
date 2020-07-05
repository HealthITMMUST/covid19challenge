package main

import (
	"fmt"
	"log"
	"os"

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
	/*err := godotenv.Load()
	if err != nil {
		log.Println(err)

	} */
	token = os.Getenv("TELEGRAM_TOKEN")
	//export TELEGRAM_TOKEN=1224895113:AAFdB8-X_Iav2z6v0UkFMumAxeQKeDSuCJI

	token = "1224895113:AAFdB8-X_Iav2z6v0UkFMumAxeQKeDSuCJI"
}

func main() {
	bot = tbot.New(token, tbot.WithWebhook("https://erevusha1.herokuapp.com", ":"+os.Getenv("PORT")))
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/symptoms", app.symptoms)
	bot.HandleMessage("/handling", app.handling)
	bot.HandleMessage("/helpline", app.helpline)

	log.Fatal(bot.Start())
}

func (a *application) startHandler(m *tbot.Message) {
	msg := "welcome to Erevusha bot \n Select: \n /symptoms to get basic COVID-19 symptoms \n /handle to handle infected patients \n /selftest to perfom A basic Test \n /helpline to coontact MOH for enquiries "
	a.client.SendMessage(m.Chat.ID, msg)
}
func (a *application) symptoms(m *tbot.Message) {

	msg := "Signs and symptoms of coronavirus disease 2019 (COVID-19) may appear two to 14 days after exposure. This time after exposure and before having symptoms is called the incubation period. Mild signs and symptoms can include:Fever \n Cough \n shortness of breath \n aches and pains" +
		" \nsore throat \n diarrhoea \n conjunctivitis \n headache \n loss of taste or smell \n a rash on skin, \n  discolouration of fingers or toes " +
		"SERIOUS SYMPTOMS MAY INCLUDE \n difficulty breathing or shortness of breath \n chest pain or pressure \n loss of speech or movement"
	a.client.SendMessage(m.Chat.ID, msg)
}
func (a *application) handling(m *tbot.Message) {

	msg := fmt.Sprintf("Educate patients and household members on personal hygiene, basic Infection Prevention &Control measures. \n" +
		" 1.Place the patient in a well-ventilated single room (i.e., with open windows and an open door). \n" +
		"2.Limit the movement of the patient in the house and minimize shared space. \n" +
		"3.Household members should stay in a different room or, maintain a distance of at least 1 m from the ill person (e.g., sleep in a separate bed) \n" +
		"4.Limit the number of caregivers. Ideally, assign one person who is in a good health and has no underlying chronic or immuno compromising conditions \n" +
		"5.Perform hand hygiene after any type of contact with patients or their immediate environment \n" +
		"6.when washing hands with soap and water, it is preferable to use disposable paper towels to dry hands \n " +
		"7.A medical mask should be provided to the patient and worn as much as possible \n" +
		"8.Caregivers should wear a tightly fitted medical mask that covers their mouth and nose when in the same room as the patient. Masks should not be touched or handled during use \n" +
		"9.Avoid direct contact with body fluids, particularly oral or respiratory secretions, and stool")
	a.client.SendMessage(m.Chat.ID, msg)
}

func (a *application) helpline(m *tbot.Message) {
	msg := " For further clarification Please Dial toll Free number 719 to get in touch with MOH agents \n Remember to always wear a mask when out in the public \n and remember to Sanitize frequently with an Alcohol Based Sanitizer and plenty of running water"
	a.client.SendMessage(m.Chat.ID, msg)
	msg1 := "Visit https://www.health.go.ke/ to get more updates"
	a.client.SendMessage(m.Chat.ID, msg1)
}
