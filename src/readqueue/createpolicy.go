package main

import (
	"encoding/json"
)

type Policy struct {
	Version   string    `json:"Version"`
	ID        string    `json:"Id"`
	Statement Statement `json:"Statement"`
}

type Statement struct {
	Effect    string    `json:"Effect"`
	Principal Principal `json:"Principal"`
	Action    string    `json:"Action"`
	Resource  string    `json:"Resource"`
	Condition Condition `json:"Condition"`
}

type Principal struct {
	AWS string `json:"AWS"`
}

type Condition struct {
	ArnLike ArnLike `json:"ArnLike"`
}

type ArnLike struct {
	AwsSourceArn string `json:"aws:SourceArn"`
}

func createPolicy(sqs_arn string, sns_arn string) ([]byte, error) {

	resp := &Policy{
		Version: "2008-10-17",
		ID:      "provision",
		Statement: *&Statement{
			Effect: "Allow",
			Principal: *&Principal{
				AWS: "*",
			},
			Action:   "SQS:SendMessage",
			Resource: sqs_arn,
			Condition: *&Condition{
				ArnLike: *&ArnLike{
					AwsSourceArn: sns_arn,
				},
			},
		},
	}

	return json.Marshal(resp)

}
