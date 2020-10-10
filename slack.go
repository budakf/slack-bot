package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/net/websocket"
)

const (
	slackRtmUrl = "https://slack.com/api/rtm.connect?token=%s"
	slackApi    = "https://api.slack.com/"
)

type self struct{
	Id string `json:"id"`
}

type resObject struct {
	Ok    	string 	`json:"ok"`
	Error 	string 	`json:"error"`
	Url   	string 	`json:"url"`
	SelfId  self 	`json:"self"`
}

type message struct {
	Channel	string 	`json:"channel"`
	Text	string 	`json:"text"`
	Type	string 	`json:"type"`
}



func getMessage(ws *websocket.Conn) (msg message, err error) {
	err = websocket.JSON.Receive(ws, &msg)
	return
}

func sendMessage(ws *websocket.Conn, msg message) (err error) {
	err = websocket.JSON.Send(ws, msg)
	return
}

func getWsUrlFromSlack(slackRtmUrl, token string) ( wsUrl, id string, err error) {
	
	url := fmt.Sprintf(slackRtmUrl, token)
	res, err := http.Get(url)
	if err != nil{
		return 
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return 
	}

	var resObj resObject
	json.Unmarshal(resBody, &resObj)
	wsUrl = resObj.Url
	id = resObj.SelfId.Id
	return 
}

func connectToSlackRtmApi(token string) (ws *websocket.Conn, id string) {

	wsUrl, id, err := getWsUrlFromSlack(slackRtmUrl, token)
	if err != nil {
		log.Fatal(err)
	}

	ws, err = websocket.Dial(wsUrl, "", slackApi)
	if err != nil {
		log.Fatal(err)
	}

	return
}

