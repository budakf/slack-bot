package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Content struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	SlackToken string `json:"slack-token"`
}

func configReader(fileName string) (content Content) {

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("msg")
		log.Fatal(err)
	}

	byteContent, err := ioutil.ReadAll(file)
	json.Unmarshal(byteContent, &content)

	return
}
