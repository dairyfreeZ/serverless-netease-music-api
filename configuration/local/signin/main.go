package main

import (
	"log"

	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/api/signin/handler"
)

var h = handler.NewHandler()

func main() {
	if err := h.Run(handler.HandlerArgs{
		Username: "",
		Password: "",
		State: handler.State{
			Location: "",
			Region:   "",
		},
		IP: "",
	}); err != nil {
		log.Fatal(err)
	}
}
