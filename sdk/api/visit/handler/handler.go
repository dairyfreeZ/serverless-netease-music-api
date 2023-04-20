package handler

import (
	"fmt"
	"strings"

	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/pkg/request"

	log "github.com/sirupsen/logrus"
)

const (
	loginPath  = "weapi/login"
	typeMobile = 0
)

type HandlerArgs struct {
	Username string
	Password string
	Path     string
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

func (h *Handler) Run(args HandlerArgs) (string, error) {
	// Set up client.
	client, err := request.NewNMClient(args.State.Location, args.State.Region, args.IP)
	if err != nil {
		return "", fmt.Errorf("failed to init client: %w", err)
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
			return "", fmt.Errorf("failed to login: %w", err)
		}
		log.Info("Login succeeded.")
	} else {
		log.Infof("Cookies are fresh, no login required.")
	}

	// Visit.
	detail, err := client.GET(args.Path)
	if err != nil {
		return "", err
	}

	// Save state.
	if args.State.Location != "" {
		log.Infof("Exporting state to: %s", args.State.Location)
		if err := client.ExportState(args.State.Location, args.State.Region); err != nil {
			return "", err
		}
	}

	return detail, nil
}
