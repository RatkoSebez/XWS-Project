package main

import (
	"XWS-Project/api_gateway/startup"
	"XWS-Project/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
