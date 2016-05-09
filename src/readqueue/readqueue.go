package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	getMessage()
}

func getMessage() {
	svc := sqs.New(session.New())

params := &sqs.ReceiveMessageInput{
	QueueUrl: aws.String("https://sqs.us-east-1.amazonaws.com/659527370395/provision-queue"), // Required
	AttributeNames: []*string{
		aws.String("QueueAttributeName"), // Required
		// More values...
	},
	MaxNumberOfMessages: aws.Int64(10),
	MessageAttributeNames: []*string{
		aws.String("ufufyjyv"), // Required
		// More values...
	},
}


resp, err := svc.ReceiveMessage(params)

if err != nil {
	// Print the error, cast err to awserr.Error to get the Code and
	// Message from an error.
	fmt.Println(err.Error())
	return
}

l := resp.Messages
for i := 0; i < len(l); i ++ {
    v := l[i].Body
    fmt.Println(*v)
}

}