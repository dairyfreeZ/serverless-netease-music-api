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
	State    struct {
		Location string `json:"location"`
		Region   string `json:"region"`
	} `json:"state"`
	IP string `json:"ip"`
}

func handleRequest(ctx context.Context, event Event) error {
	return h.Run(handler.HandlerArgs{
		Username: event.Username,
		Password: event.Password,
		State: handler.State{
			Location: event.State.Location,
			Region:   event.State.Region,
		},
		IP: event.IP,
	})
}

func main() {
	lambda.Start(handleRequest)
}
