package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Descrition string `json:"password"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")

	name := request.QueryStringParameters["name"]
	//body := fmt.Sprintf("Hello %s!", name)

	body, err := json.Marshal(Event{
		Id:         "askdhasldjkhask",
		Title:      name,
		Descrition: "netlify-function-example",
	})

	if err != nil {
		panic(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(body),
		IsBase64Encoded: false,
	}, nil
}

func main() {

	lambda.Start(handler)
}
