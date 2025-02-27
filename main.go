package main

import (
	"fmt"
	"log"
	"moderator/cfg"
	"moderator/handlers"
	"net/http"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/moderate", handlers.ModerateComment)

	fmt.Printf("Server started at :%s\n", config.AppConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.ServerPort, nil))
}
