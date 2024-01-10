package main

import (
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/stretchr/testify/mock"
)

// MockDynamoDBAPI is a mock for DynamoDBAPI
type MockDynamoDBAPI struct {
    mock.Mock
}

func (m *MockDynamoDBAPI) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
    args := m.Called(input)
    return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}
