package main

import (
	//"encoding/json"
	"fmt"
	// "io/ioutil"
	// "os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func createQueue(sqs_arn string, sns_arn string) *sqs.CreateQueueOutput {

	svc := sqs.New(ses)
	// file, e := ioutil.ReadFile("./output.json")

	// if e != nil {
	// 	fmt.Printf("File error: %v\n", e)
	// 	os.Exit(1)
	// }

	json, e := createPolicy(sqs_arn, sns_arn)

	if e != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(e.Error())
		return nil
	}

	params := &sqs.CreateQueueInput{
		QueueName: aws.String(config.SqsQueue),
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
	fmt.Println("Queue created: " + sqs_arn)
	return resp
}

func getQueueArn(queueUrl string, session client.ConfigProvider) string {
	svc := sqs.New(ses)

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
	svc := sqs.New(ses)

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
