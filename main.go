package main

import (
	"fmt"
	"log"

	// "time"
	"strings"
)

func main() {
	content := configReader("config.json")
	ws, _ := connectToSlackRtmApi(content.SlackToken)
	fmt.Println("Bot is ready ")

	for {
		msg, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		if msg.Type == "message" {
			if strings.HasPrefix(msg.Text, "binance") {
				binance(&msg)
			} else {
				msg.Text = "Please enter any appropriate command\ne.g. \"binance BTCUSDT\""
			}
			sendMessage(ws, msg)
		}

	}

}
