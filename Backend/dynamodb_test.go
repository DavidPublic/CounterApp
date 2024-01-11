package main

import (
    "testing"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/stretchr/testify/mock"
)


func TestCreateCounter(t *testing.T) {
    mockSvc := new(MockDynamoDBAPI)
    counter := &Counter{Name: "Test Counter", Value: 0}

    // Set up the expected response and error to return from the mock
    mockSvc.On("PutItem", mock.AnythingOfType("*dynamodb.PutItemInput")).Return(&dynamodb.PutItemOutput{}, nil)

    err := CreateCounter(mockSvc, counter)
    if err != nil {
        t.Errorf("CreateCounter() error = %v, wantErr %v", err, nil)
    }

    // Assert that the expectations were met
    mockSvc.AssertExpectations(t)
}
