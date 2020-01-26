package main

import (
	"fmt"
	"ready_go/94_config/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
