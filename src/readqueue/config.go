package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Region      string `json:"Region"`
	Snstopic    string `json:"Sns_topic"`
	Environment string `json:"Environment"`
	ID          string `json:"ID"`
	SqsQueue    string `json:"Sqs_Queue"`
}

func readConfig(config Config) {
	fmt.Println(config.Region)
	fmt.Println(config.Snstopic)
	fmt.Println(config.Environment)
	fmt.Println(config.ID)
	fmt.Println(config.SqsQueue)

}

func openConfig() Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}

	decoder := json.NewDecoder(configFile)
	config := Config{}
	e := decoder.Decode(&config)
	if e != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(e.Error())
	}
	readConfig(config)
	return config
}
