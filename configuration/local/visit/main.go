package main

import (
	"log"

	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/api/visit/handler"
)

var h = handler.NewHandler()

func main() {
	_, err := h.Run(handler.HandlerArgs{
		Username: "",
		Password: "",
		Path:     "",
		State: handler.State{
			Location: "",
			Region:   "",
		},
		IP: "",
	})
	if err != nil {
		log.Fatal(err)
	}
}
