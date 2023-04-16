package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/api/signin/handler"
)

var h = handler.NewHandler()

type Event struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRequest(ctx context.Context, eventRaw map[string]interface{}) error {
	var event Event
	// FC triggered by EventBridge of Alicloud will have nested payload.
	if payload, ok := eventRaw["payload"].(string); ok {
		if err := json.Unmarshal([]byte(payload), &event); err != nil {
			return errors.New("Failed to unmarshal event payload.")
		}
	} else {
		event.Username = eventRaw["username"].(string)
		event.Password = eventRaw["password"].(string)
	}

	if event.Username == "" {
		return errors.New("Username cannot be empty.")
	}
	if event.Password == "" {
		return errors.New("Password cannot be empty.")
	}
	return h.Run(handler.HandlerArgs{
		Username: event.Username,
		Password: event.Password,
	})
}

func main() {
	fc.Start(HandleRequest)
}
