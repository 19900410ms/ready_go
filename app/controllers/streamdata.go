package controllers

import (
	"log"
	"ready_go/app/models"
	"ready_go/bitflyer"
	"ready_go/config"
)

func StreamIngestionData() {
	var tickerChanel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChanel)
	for ticker := range tickerChanel {
		log.Printf("action=StreamIngestionData, %v", ticker)
		for _, duration := range config.Config.Durations {
			isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated == true && duration == config.Config.TradeDuration {
				// TODO
			}
		}
	}
}
