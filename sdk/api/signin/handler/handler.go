package handler

import (
	"fmt"
	"net/http"

	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/pkg/request"

	log "github.com/sirupsen/logrus"
)

const (
	loginPath  = "weapi/login"
	signinPath = "weapi/point/dailyTask"
	typeMobile = 0
)

type HandlerArgs struct {
	Username string
	Password string
}

type Handler struct {
	header http.Header
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Run(args HandlerArgs) error {
	// TODO: Get cookies from remote storage.

	// Set up client w/o cookies.
	client, err := request.NewNMClient(nil)
	if err != nil {
		return fmt.Errorf("failed to init client: %w", err)
	}

	// Login.
	body := map[string]interface{}{
		"username": args.Username,
		"password": args.Password,
	}
	rsp, err := client.POST(body, loginPath)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	log.Infof("login: %s", rsp)

	// Sign in.
	body = map[string]interface{}{
		"type": typeMobile,
	}
	rsp, err = client.POST(body, signinPath)
	if err != nil {
		return fmt.Errorf("failed to sign in: %w", err)
	}
	log.Infof("sign in: %s", rsp)

	return nil
}
