package main

import (
// "fmt"
)

var config

func main() {
	region := aws.NewConfig().WithRegion(config.Region)
    session := session.New(region)
	config := openConfig()
	sns_arn := generateSnsArn(config)
	sqs_arn := generateSqsArn(config)
	createQueue(sqs_arn, sns_arn, config , session)
	// fmt.Println("na createque")
	// sqs_arn := getQueueArn(*queue.QueueUrl)
	// fmt.Println("na getqueuearn")
	//getQueuePolicy(*queue.QueueUrl)
	subscribe(sqs_arn, sns_arn)
	readqueue(*queue.QueueUrl)
}
