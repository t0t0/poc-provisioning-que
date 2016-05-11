package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func getFile(key string, handler string, queueUrl string) {
	path := "tmp/"
	location := path + key

	file, err := os.Create(location)
	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	defer file.Close()
	downloader := s3manager.NewDownloader(session.New(&aws.Config{Region: aws.String("us-east-1")}))
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
	deleteMessage(handler, queueUrl)

}
