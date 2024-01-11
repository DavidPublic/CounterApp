package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

// MockDynamoDBAPI is a mock for DynamoDBAPI
type MockDynamoDBAPI struct {
	mock.Mock
}

// DeleteItem implements DynamoDBAPI.
func (*MockDynamoDBAPI) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	panic("unimplemented")
}

// GetItem implements DynamoDBAPI.
func (*MockDynamoDBAPI) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	panic("unimplemented")
}

// Scan implements DynamoDBAPI.
func (*MockDynamoDBAPI) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	panic("unimplemented")
}

// UpdateItem implements DynamoDBAPI.
func (*MockDynamoDBAPI) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	panic("unimplemented")
}

func (m *MockDynamoDBAPI) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}
