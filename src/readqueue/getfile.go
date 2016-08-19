package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func getFile(key string, handler string, queueUrl string) {
	isfile, whatfile := check(key)
	if isfile {
		location := "./"
		if whatfile == "config" {
			location = config.Configdir + "/" + key
		}
		if whatfile == "template" {
			location = config.Templatesdir + "/" + key
		}

		file, err := os.Create(location)
		if err != nil {
			log.Fatal("Failed to create file", err)
		}
		defer file.Close()
		downloader := s3manager.NewDownloader(ses)
		numBytes, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String("poc-provision"),
				Key:    aws.String(key),
			})
		if err != nil {
			fmt.Println("Failed to download file", err)
			return
		}
		fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")

	}
	deleteMessage(handler, queueUrl)

}

func check(key string) (bool, string) {

	for _, e := range configFiles {
		if strings.Compare(key, e) == 0 {
			fmt.Println(key, "is a config file i require")
			return true, "config"
		}
	}
	for _, e := range templateFiles {
		if strings.Compare(key, e) == 0 {
			fmt.Println(key, "is a template file i require")
			return true, "template"
		}
	}
	fmt.Println(key, "is a not a file i require")
	return false, "no file"
}

func buildFileList() {
	//fmt.Println(len(config.Containers))

	for _, e := range config.Containers {
		//fmt.Println(e.Config)
		configFiles = append(configFiles, e.Config...)
		//fmt.Println(e.Templates)
		templateFiles = append(templateFiles, e.Templates...)
	}
	fmt.Println("config files:")
	fmt.Println(configFiles)
	fmt.Println("template files:")
	fmt.Println(templateFiles)

}
