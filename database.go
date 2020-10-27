package main

import (
	"edgarchirivella.com/sentiment-scrapper/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

var dynamoDbInstance *dynamodb.DynamoDB

func init() {
	dynamoDbRegion := os.Getenv("AWS_DYNAMODB_REGION")
	dynamoDbEndpoint := os.Getenv("AWS_DYNAMODB_ENDPOINT")

	config := &aws.Config{
		Region:   aws.String(dynamoDbRegion),
		Endpoint: aws.String(dynamoDbEndpoint),
	}

	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)

	dynamoDbInstance = svc
}

// TODO: Refactor to make more generic
func SaveNewsItem(item entity.NewsItem) error {
	tableName := os.Getenv("AWS_DYNAMODB_TABLE_NAME")

	attributeValues, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValues,
		TableName: aws.String(tableName),
	}

	_, err = dynamoDbInstance.PutItem(input)

	return err
}
