package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func subscribe(topic_arn string, sqs_arn string) {

	svc := sns.New(session.New())

	params := &sns.SubscribeInput{
		Protocol: aws.String("sqs"),     // Required
		TopicArn: aws.String(topic_arn), // Required
		Endpoint: aws.String(sqs_arn),
	}
	resp, err := svc.Subscribe(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
