package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func deleteMessage(handler string, queueUrl string) {

	svc := sqs.New(ses)

	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl), // Required
		ReceiptHandle: aws.String(handler),  // Required
	}
	_, err := svc.DeleteMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println("Message deleted from " + queueUrl)

}
