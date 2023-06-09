package handler

import (
	"fmt"
	"strings"

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
	State    State
	IP       string
}

type State struct {
	Location string
	Region   string
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Run(args HandlerArgs) error {
	// Set up client.
	client, err := request.NewNMClient(args.State.Location, args.State.Region, args.IP)
	if err != nil {
		return fmt.Errorf("failed to init client: %w", err)
	}

	// Login.
	homepageDetail, _ := client.GET("#")
	if !strings.Contains(homepageDetail, args.Username) {
		log.Infof("No state provided or cookies expired, attempt to login.")
		body := map[string]interface{}{
			"username": args.Username,
			"password": args.Password,
		}
		_, err := client.POST(body, loginPath)
		if err != nil {
			return fmt.Errorf("failed to login: %w", err)
		}
		log.Info("Login succeeded.")
	} else {
		log.Infof("Cookies are fresh, no login required.")
	}

	// Sign in.
	body := map[string]interface{}{
		"type": typeMobile,
	}
	rsp, err := client.POST(body, signinPath)
	if err != nil {
		return fmt.Errorf("failed to sign in: %w", err)
	}
	log.Infof("Sign in: %s", rsp)

	if args.State.Location != "" {
		log.Infof("Exporting state to: %s", args.State.Location)
		if err := client.ExportState(args.State.Location, args.State.Region); err != nil {
			return err
		}
	}

	return nil
}
