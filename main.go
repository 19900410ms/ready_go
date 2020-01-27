package main

import (
	"fmt"
	"ready_go/app/controllers"
	"ready_go/app/models"
	"ready_go/config"
	"ready_go/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
}
