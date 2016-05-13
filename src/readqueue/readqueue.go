package main

import (
	// "fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
)

var config Config
var ses client.ConfigProvider
var configFiles []string
var templateFiles []string

func main() {
	openConfig()
	buildFileList()
	region := aws.NewConfig().WithRegion(config.Region)
	ses = session.New(region)
	sns_arn := generateSnsArn(config)
	sqs_arn := generateSqsArn(config)
	queue := createQueue(sqs_arn, sns_arn)
	// fmt.Println("na createque")
	// sqs_arn := getQueueArn(*queue.QueueUrl)
	// fmt.Println("na getqueuearn")
	//getQueuePolicy(*queue.QueueUrl)
	subscribe(sqs_arn, sns_arn)
	readqueue(*queue.QueueUrl)
}
