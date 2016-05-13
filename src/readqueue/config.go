package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Region       string      `json:"Region"`
	SnsTopic     string      `json:"Sns_topic"`
	Environment  string      `json:"Environment"`
	ID           string      `json:"ID"`
	SqsQueue     string      `json:"Sqs_queue"`
	Containers   []Container `json:"containers"`
	Configdir    string      `json:"configdir"`
	Templatesdir string      `json:"templatesdir"`
}

type Container struct {
	Config    []string `json:"config"`
	Templates []string `json:"templates"`
	Image     string   `json:"image"`
	Count     string   `json:"count"`
	Name      string   `json:"name"`
}

func readConfig(config Config) {
	fmt.Println(config)
	fmt.Println(config.Region)
	fmt.Println(config.SnsTopic)
	fmt.Println(config.Environment)
	fmt.Println(config.ID)
	fmt.Println(config.Containers[0].Name)

}

func openConfig() {
	configFile, err := os.Open("config.json")
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}

	decoder := json.NewDecoder(configFile)
	config = Config{}
	e := decoder.Decode(&config)
	if e != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(e.Error())
	}
	readConfig(config)
}
