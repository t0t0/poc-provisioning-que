package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)


func readqueue() {
	svc := sqs.New(session.New())

params := &sqs.ListQueuesInput{
	QueueNamePrefix: aws.String(""),
}
resp, err := svc.ListQueues(params)

if err != nil {
	// Print the error, cast err to awserr.Error to get the Code and
	// Message from an error.
	fmt.Println(err.Error())
	return
}

// Pretty-print the response data.
fmt.Println(resp)

}