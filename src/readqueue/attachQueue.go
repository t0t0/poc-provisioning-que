package main

import (
	//"encoding/json"
	"fmt"
	// "io/ioutil"
	// "os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func createQueue(queueName string, sns_arn string) *sqs.CreateQueueOutput {

	svc := sqs.New(session.New())
	// file, e := ioutil.ReadFile("./output.json")

	// if e != nil {
	// 	fmt.Printf("File error: %v\n", e)
	// 	os.Exit(1)
	// }

	json, e := createPolicy("arn:aws:sqs:us-east-1:659527370395:"+queueName, sns_arn)

	if e != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(e.Error())
		return nil
	}
	fmt.Println(string(json))

	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]*string{
			"Policy": aws.String(string(json)),
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
