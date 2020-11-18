package main

import (
	"fmt"
	"log"
	"net/http"

	"hello-again-go/config"
	"hello-again-go/handler"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error during env var loading: %s", err.Error())
	}
	handler, err := handler.NewRequestHandler(*conf)
	if err != nil {
		log.Fatalf("Error during handler initialization: %s", err.Error())
	}
	http.HandleFunc("/image", handler.RespondImage)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
}
