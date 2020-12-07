package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"hello-again-go/config"
	"hello-again-go/handler"
)

func main() {
	log.Println("System loading...")
	conf, err := config.LoadConfig()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err.Error(),
			}).Errorf("Error during env var loading: %s", err.Error())
	}
	handler, err := handler.NewRequestHandler(*conf)
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err.Error(),
			}).Errorf("Error during handler initialization: %s", err.Error())
	}
	http.HandleFunc("/image", handler.RespondImage)
	log.Println("System ready! Waiting instructions...")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil); err != nil {
		log.WithFields(
			log.Fields{
				"error": err.Error(),
			}).Errorf("Error during server starting: %s", err.Error())
	}
}
