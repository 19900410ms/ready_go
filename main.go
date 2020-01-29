package main

import (
	"log"
	"ready_go/app/controllers"
	"ready_go/config"
	"ready_go/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
	log.Println(controllers.StartWebServer())
}
