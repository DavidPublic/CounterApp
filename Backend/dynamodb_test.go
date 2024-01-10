package main

import (
    "testing"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/stretchr/testify/mock"
)

func TestCreateCounter(t *testing.T) {
    mockSvc := new(MockDynamoDBAPI)
    counter := &Counter{Name: "Test Counter", Value: 0}

    mockSvc.On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, nil)

    err := CreateCounter(mockSvc, counter)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    mockSvc.AssertExpectations(t)
}
