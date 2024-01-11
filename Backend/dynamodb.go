package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"log"
)

// DynamoDBAPI is an interface for DynamoDB operations
type DynamoDBAPI interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

var db DynamoDBAPI = dynamodb.New(session.Must(session.NewSession(&aws.Config{
	Region: aws.String("eu-north-1"), // Replace with your AWS region
})))

const tableName = "CountersTable"

func CreateCounter(svc DynamoDBAPI, counter *Counter) error {
	av, err := dynamodbattribute.MarshalMap(counter)
	if err != nil {
		log.Printf("Error marshalling new counter item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("CountersTable"), // Replace with your table name
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Error calling PutItem: %s", err)
		return err
	}

	return nil
}

func IncrementCounter(name string, increment int) error {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"Name": name})
	if err != nil {
		log.Printf("Error marshalling key: %s", err)
		return err
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
		log.Printf("Error updating counter: %s", err)
		return err
	}

	return nil
}

func GetAllCounters() ([]*Counter, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := db.Scan(input)
	if err != nil {
		log.Printf("Error scanning table: %s", err)
		return nil, err
	}

	var counters []*Counter
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &counters)
	if err != nil {
		log.Printf("Error unmarshalling results: %s", err)
		return nil, err
	}

	return counters, nil
}

func GetCounter(name string) (*Counter, error) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"Name": name})
	if err != nil {
		log.Printf("Error marshalling key: %s", err)
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	result, err := db.GetItem(input)
	if err != nil {
		log.Printf("Error getting item: %s", err)
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var counter Counter
	err = dynamodbattribute.UnmarshalMap(result.Item, &counter)
	if err != nil {
		log.Printf("Error unmarshalling item: %s", err)
		return nil, err
	}

	return &counter, nil
}

func DeleteCounter(name string) error {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"Name": name})
	if err != nil {
		log.Printf("Error marshalling key: %s", err)
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	_, err = db.DeleteItem(input)
	if err != nil {
		log.Printf("Error deleting item: %s", err)
		return err
	}

	return nil
}
