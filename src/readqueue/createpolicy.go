package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"os"
)

type Policy struct {
	Version   string `json:"Version"`
	ID        string `json:"Id"`
	Statement []struct {
		Effect    string `json:"Effect"`
		Principal struct {
			AWS string `json:"AWS"`
		} `json:"Principal"`
		Action    string `json:"Action"`
		Resource  string `json:"Resource"`
		Condition struct {
			ArnLike struct {
				AwsSourceArn string `json:"aws:SourceArn"`
			} `json:"ArnLike"`
		} `json:"Condition"`
	} `json:"Statement"`
}

func createPolicy(sqs_arn string, sns_arn string) *json {

	resp := &Policy{
		Version: "2008-10-17",
		Id:      "provision",
		Statement: {
			Effect: "Allow",
			Principal: {
				AWS: "*",
			},
			Action:   "SQS:SendMessage",
			Resource: sqs_arn,
			Condition: {
				ArnLike: {
					AwsSourceArn: sns_arn,
				},
			},
		},
	}

	policy, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	return policy

}
