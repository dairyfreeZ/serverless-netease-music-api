package main

import (
	"log"

	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/api/signin/handler"
)

var h = handler.NewHandler()

func main() {
	if err := h.Run(handler.HandlerArgs{
		Username: "your_user_name",
		Password: "your_md5_encoded_password",
	}); err != nil {
		log.Fatal(err)
	}
}
