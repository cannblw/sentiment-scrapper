package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func InitializeDynamodb() {
	dynamoDbRegion := os.Getenv("AWS_DYNAMODB_REGION")
	dynamoDbEndpoint := os.Getenv("AWS_DYNAMODB_ENDPOINT")

	config := &aws.Config{
		Region:   aws.String(dynamoDbRegion),
		Endpoint: aws.String(dynamoDbEndpoint),
	}

	sess := session.Must(session.NewSession(config))

	// TODO: Use this variable
	svc := dynamodb.New(sess)
}
