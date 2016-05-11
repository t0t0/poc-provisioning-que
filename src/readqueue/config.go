package main

import "encoding/json"

type Config struct {
	Region   string `json:"Region"`
	SnsQueue string `json:"Sns_queue"`
}

func readConfig() {

}
