package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ddb *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}

// CompleteTodo handles marking an existing todo as completed.
func CompleteTodo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("CompletTodo")

	var (
		id        = request.PathParameters["id"]
		tableName = aws.String(os.Getenv("TODOS_TABLE_NAME"))
		done      = "done"
	)

	// Update row
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		UpdateExpression: aws.String("set #d = :d"),
		ExpressionAttributeNames: map[string]*string{
			"#d": &done,
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":d": {
				BOOL: aws.Bool(true),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
		TableName:    tableName,
	}

	_, err := ddb.UpdateItem(input)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(CompleteTodo)
}
