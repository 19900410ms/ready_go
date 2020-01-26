package main

import (
	"log"
	"ready_go/config"
	"ready_go/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
