package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Region   string `json:"Region"`
	SnsQueue string `json:"Sns_queue"`
}

func readConfig() {
	fmt.Println(Config.Region)
	fmt.Println(Config.SnsQueue)

}
