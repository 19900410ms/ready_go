package main

import (
	"fmt"
	"ready_go/bitflyer"
	"ready_go/config"
	"ready_go/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
