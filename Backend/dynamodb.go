package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Initialize a DynamoDB session
var db dynamodbiface.DynamoDBAPI = dynamodb.New(session.Must(session.NewSession(&aws.Config{
	Region: aws.String("eu-north-1"), // Replace with your region
})))

// DynamoDB table name
const tableName = "CountersTable"

// CreateCounter adds a new counter to the DynamoDB table.
func CreateCounter(counter *Counter) {
	av, err := dynamodbattribute.MarshalMap(counter)
	if err != nil {
		log.Fatalf("Got error marshalling new counter item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

// IncrementCounter increments the value of a given counter.
func IncrementCounter(name string, increment int) {
    key, err := dynamodbattribute.MarshalMap(map[string]string{"Name": name})
    if err != nil {
        log.Fatalf("Got error marshalling key: %s", err)
    }

    update := map[string]*dynamodb.AttributeValue{
        ":val": {
            N: aws.String("1"),
        },
        ":zero": {
            N: aws.String("0"),
        },
    }

    input := &dynamodb.UpdateItemInput{
        Key:                       key,
        TableName:                 aws.String(tableName),
        UpdateExpression:          aws.String("set #V = if_not_exists(#V, :zero) + :val"),
        ExpressionAttributeValues: update,
        ExpressionAttributeNames:  map[string]*string{"#V": aws.String("Value")},
        ReturnValues:              aws.String("UPDATED_NEW"),
    }

    _, err = db.UpdateItem(input)
    if err != nil {
        log.Fatalf("Got error updating counter: %s", err)
    }
}


// GetAllCounters retrieves all counters from the DynamoDB table.
func GetAllCounters() []*Counter {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := db.Scan(input)
	if err != nil {
		log.Fatalf("Got error scanning table: %s", err)
	}

	var counters []*Counter
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &counters)
	if err != nil {
		log.Fatalf("Got error unmarshalling results: %s", err)
	}

	return counters
}

// GetCounter retrieves a single counter by name.
func GetCounter(name string) (*Counter, error) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"Name": name})
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var counter Counter
	err = dynamodbattribute.UnmarshalMap(result.Item, &counter)
	if err != nil {
		return nil, err
	}

	return &counter, nil
}

// DeleteCounter removes a counter from the DynamoDB table.
func DeleteCounter(name string) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"Name": name})
	if err != nil {
		log.Fatalf("Got error marshalling key: %s", err)
	}

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	_, err = db.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem: %s", err)
	}
}
