package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/api/signin/handler"
)

var h = handler.NewHandler()

type Event struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleRequest(ctx context.Context, event Event) error {
	return h.Run(handler.HandlerArgs{
		Username: event.Username,
		Password: event.Password,
	})
}

func main() {
	lambda.Start(handleRequest)
}
