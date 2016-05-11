package main

import (
// "fmt"
)

const sns_arn = "arn:aws:sns:us-east-1:659527370395:provision"

func main() {
	queue := createQueue("test")
	// fmt.Println("na createque")
	sqs_arn := getQueueArn(*queue.QueueUrl)
	// fmt.Println("na getqueuearn")
	//getQueuePolicy(*queue.QueueUrl)
	subscribe(sns_arn, sqs_arn)
	readqueue(*queue.QueueUrl)
}
