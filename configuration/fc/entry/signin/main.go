package main

import (
	"context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/api/signin/handler"
)

var h = handler.NewHandler()

type Event struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRequest(ctx context.Context, event Event) error {
	return h.Run(handler.HandlerArgs{
		Username: event.Username,
		Password: event.Password,
	})
}

func main() {
	fc.Start(HandleRequest)
}
