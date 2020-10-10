package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	binanceApi = "https://api.binance.com/"
)

type Response struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func binance(msg *message) {
	commandParts := strings.Split(msg.Text, " ")

	if len(commandParts) == 2 {
		result := make(chan Response, 1)
		getPriceBySymbol(result, commandParts[1])

		select {
		case p := <-result:
			if p.Msg == "Invalid symbol." {
				msg.Text = "Please enter any appropriate command\ne.g. \"binance BTCUSDT\""
			} else {
				resultStr := fmt.Sprintf("%s ---> %s", p.Symbol, p.Price)
				msg.Text = resultStr
			}

		case <-time.After(4 * time.Second):
			msg.Text = "Time out\nplease try again"

		}
	} else {
		msg.Text = "Please enter any appropriate command\ne.g. \"binance BTCUSDT\""
	}
}

func getPriceBySymbol(result chan<- Response, symbol string) {

	const pricePostFix = "api/v3/ticker/price?symbol=%s"
	urlPostFix := fmt.Sprintf(pricePostFix, symbol)
	url := binanceApi + urlPostFix
	res, err := http.Get(url)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var response Response
	json.Unmarshal(resBody, &response)
	result <- response
}
