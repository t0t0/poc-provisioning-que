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
	arn, e := svc.Subscribe(params)

	if e != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(e.Error())
		return
	}

	vars := &sns.SetSubscriptionAttributesInput{
		AttributeName:   aws.String("RawMessageDelivery"), // Required
		SubscriptionArn: aws.String(*arn.SubscriptionArn), // Required
		AttributeValue:  aws.String("true"),
	}

	_, err := svc.SetSubscriptionAttributes(vars)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}
}
