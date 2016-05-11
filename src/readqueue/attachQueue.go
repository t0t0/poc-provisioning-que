package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func createQueue(queueName string) *sqs.CreateQueueOutput {

	svc := sqs.New(session.New())
	file, e := ioutil.ReadFile("./output.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]*string{
			"Policy": aws.String(string(file)),
		}, // Required
	}
	resp, err := svc.CreateQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	// Pretty-print the response data.
	return resp
}

func getQueueArn(queueUrl string) string {
	svc := sqs.New(session.New())

	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(queueUrl), // Required
		AttributeNames: []*string{
			aws.String("QueueArn"), // Required
			// More values...
		},
	}
	resp, err := svc.GetQueueAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return " "
	}

	// Pretty-print the response data.
	return *resp.Attributes["QueueArn"]
}

func getQueuePolicy(queueUrl string) {
	svc := sqs.New(session.New())

	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(queueUrl), // Required
		AttributeNames: []*string{
			aws.String("Policy"), // Required
			// More values...
		},
	}
	resp, err := svc.GetQueueAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println("error from getQueuePolicy")
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(*resp.Attributes["Policy"])
}
