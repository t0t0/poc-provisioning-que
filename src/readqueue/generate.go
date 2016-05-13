package main

func generateSnsArn(config Config) string {

	return "arn:aws:sns:" + config.Region + ":" + config.ID + ":" + config.SnsTopic

}

func generateSqsArn(config Config) string {

	return "arn:aws:sqs:" + config.Region + ":" + config.ID + ":" + config.SqsQueue

}
