package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/dairyfreeZ/serverless-netease-music-api/sdk/pkg/secret"

	log "github.com/sirupsen/logrus"
)

const (
	anonymousToken = "bf8bfeabb1aa84f9c8c3906c04a04fb864322804c83f5d607e91a04eae463c9436bd1a17ec353" +
		"cf780b396507a3f7464e8a60f4bbc019437993166e004087dd32d1490298caf655c2353e58daa0bc13cc7d5c198250" +
		"968580b12c1b8817e3f5c807e650dd04abd3fb8130b7ae43fcc5b"
	endpoint = "https://music.163.com"
)

type NMClient struct {
	client *http.Client
	header http.Header
	csrf   string
}

func NewNMClient(cookies []string) (*NMClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("error creating cookie jar: %w", err)
	}
	client := &http.Client{Jar: jar}

	parsedUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint: %w", err)
	}

	var csrf string
	if len(cookies) == 0 {
		cookies = []string{
			"os=ios",
			"appver=8.7.01",
			"__remember_me=true",
			fmt.Sprintf("NMTID=%s", secret.HexStr32()),
			fmt.Sprintf("_ntes_nuid=%s", secret.HexStr32()),
			fmt.Sprintf("MUSIC_A=%s", anonymousToken),
		}
	}
	parsedHttpCookies := parseCookies(cookies)
	for _, cookie := range parsedHttpCookies {
		if cookie.Name == "__csrf" {
			csrf = cookie.Value
			break
		}
	}
	client.Jar.SetCookies(parsedUrl, parseCookies(cookies))

	nmc := &NMClient{
		client: client,
		header: http.Header{
			"User-Agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Safari/605.1.15"},
			"Content-Type":    []string{"application/x-www-form-urlencoded"},
			"Referer":         []string{"https://music.163.com"},
			"X-Real-IP":       []string{"::1"},
			"X-Forwarded-For": []string{"::1"},
		},
		csrf: csrf,
	}
	return nmc, nil
}

// POST.
func (nmc *NMClient) POST(body map[string]interface{}, path string) (string, error) {
	bodyStr, err := nmc.prepare(body)
	if err != nil {
		return "", fmt.Errorf("failed to prepare req body: %w", err)
	}

	url := fmt.Sprintf("%s/%s", endpoint, path)
	log.Infof("Posting to %s", url)
	req, err := http.NewRequest("POST", url, strings.NewReader(bodyStr))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header = nmc.header

	return nmc.send(req)
}

func (nmc *NMClient) prepare(bodyMap map[string]interface{}) (string, error) {
	if nmc.csrf != "" {
		bodyMap["csrf_token"] = nmc.csrf
	}
	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal map: %w", err)
	}
	body, err := secret.EncryptRequest(bodyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt request body: %w", err)
	}
	return body, nil
}

func (nmc *NMClient) send(req *http.Request) (string, error) {
	rsp, err := nmc.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer rsp.Body.Close()

	log.Infof("Status Code: %v", rsp.StatusCode)
	return "OK", nil
}

func parseCookies(rawCookies []string) []*http.Cookie {
	req := &http.Request{}
	req.Header = http.Header{}
	for _, rawCookie := range rawCookies {
		req.Header.Add("Cookie", rawCookie)
	}
	return req.Cookies()
}
