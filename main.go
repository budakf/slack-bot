package main

import (
	"log"
	"fmt"
	// "time"
	"strings"
)


func main(){
	content := configReader("config.json")
	ws, _ := connectToSlackRtmApi( content.SlackToken )
	fmt.Println("Bot is ready ")
	
	for{
		msg, err := getMessage(ws)		
		if err != nil {
			log.Fatal(err)
		}

		if msg.Type == "message"{
			if strings.HasPrefix(msg.Text, "binance"){
				binance(&msg)
			}
			sendMessage(ws, msg)
		}

	}	

}


