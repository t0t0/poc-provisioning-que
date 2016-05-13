package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func getMessages(queueUrl string) {
	svc := sqs.New(ses)

	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queueUrl), // Required
		AttributeNames: []*string{
			aws.String("QueueAttributeName"), // Required
			// More values...
		},
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String(""), // Required
			// More values...
		},
		WaitTimeSeconds:   aws.Int64(10),
		VisibilityTimeout: aws.Int64(1),
	}

	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	l := resp.Messages
	for i := 0; i < len(l); i++ {
		v := l[i].Body
		handler := l[i].ReceiptHandle
		fmt.Println("Message received:" + *v)
		getFile(*v, *handler, queueUrl)
	}

}
