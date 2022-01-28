package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

func sendSlack(msg string) {
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	api := slack.New(token)
	channelID, timestamp, err := api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)

}
