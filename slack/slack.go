package slack

import (
	"fmt"
	"strings"

	"github.com/kmurata798/goslackit/gif"
	"github.com/slack-go/slack"
)

/*
   TODO: Change @BOT_NAME to the same thing you entered when creating your Slack application.
   NOTE: command_arg_1 and command_arg_2 represent optional parameteras that you define
   in the Slack API UI
*/
/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

/*
   RespondToEvents waits for messages on the Slack client's incomingEvents channel,
   and sends a response when it detects the bot has been tagged in a message with @<botTag>.

   EDIT THIS FUNCTION IN THE SPACE INDICATED ONLY!
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			splitMessage := strings.Fields(message) //seperates string

			//leaving a blank command will display the help menu

			if message == "" {
				sendHelp(slackClient, ev.Channel)
			}

			// handling simple commands
			switch strings.ToLower(splitMessage[0]) {
			case "help":
				sendHelp(slackClient, ev.Channel)
			case "echo":
				echoMessage(slackClient, strings.Join(splitMessage[:], " "), ev.Channel)
			case "gif":
				returnGif(slackClient, splitMessage[1:], ev.Channel)
			}
		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.DesktopNotificationEvent:
			fmt.Printf("Desktop Notification: %v\n", ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		default:
		}
	}
}

const helpMessage = "Type in `@Text to Gif <command_arg_1> <command_arg_2>` to use bot commands.\n\nCommands:\n`help`\n`echo <text>`\n`gif <text>`"

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, slackChannel string) {
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

// echoMessage will just echo anything after the echo keyword.
func echoMessage(slackClient *slack.RTM, message, slackChannel string) {
	splitMessage := strings.Fields(strings.ToLower(message))

	slackClient.SendMessage(slackClient.NewOutgoingMessage(strings.Join(splitMessage[1:], " "), slackChannel))
}

// returnGif returns a gif url based on the slack user's input text
func returnGif(slackClient *slack.RTM, args []string, slackChannel string) {
	response := "Please pass in text to get a relative Gif. Example: `@Text to Gif gif I'm hungry`"
	if len(args) == 0 {
		slackClient.SendMessage(slackClient.NewOutgoingMessage(response, slackChannel))
	}

	gifURL := gif.SearchGifs(args[0])
	slackClient.SendMessage(slackClient.NewOutgoingMessage(gifURL, slackChannel))
}
