package main

// Counter represents a counter with an ID, name, and value.
type Counter struct {
    ID    int    `json:"id" dynamodbav:"ID"`
    Name  string `json:"name" dynamodbav:"Name"`
    Value int    `json:"value" dynamodbav:"Value"`
}
