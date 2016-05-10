package main

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
)

func subscribe(topic_arn string, sqs_arn string) {
	
var params = {
  Protocol: "sqs", /* required */
  TopicArn: topic_arn, /* required */
  Endpoint: sqs_arn
};
sns.subscribe(params, function(err, data) {
  if (err) console.log(err, err.stack); // an error occurred
  else     console.log(data);           // successful response
});
}

