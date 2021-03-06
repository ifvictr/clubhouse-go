package clubhouse

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"

	"github.com/dghubble/sling"
	"github.com/google/uuid"
)

const (
	// BaseURL is the Clubhouse API's base endpoint.
	BaseURL = "https://www.clubhouseapi.com/api/"

	AppBuild   = 297
	AppVersion = "0.1.27"
	IOSVersion = "14.4.2"
)

type Client struct {
	sling *sling.Sling

	deviceID  uuid.UUID
	token     string
	userAgent string
	userID    int
}

type ClientOption func(*Client)

func New(opts ...ClientOption) *Client {
	client := &Client{}
	jar, _ := cookiejar.New(nil)
	httpClient := &http.Client{Jar: jar}

	client.sling = sling.New().Client(httpClient).Base(BaseURL)
	client.deviceID = uuid.New()
	client.userAgent = fmt.Sprintf("clubhouse/%d (iPhone; iOS %s; Scale/2.00)", AppBuild, IOSVersion)

	// Set default headers
	headers := map[string]string{
		"Connection": "keep-alive",
		"User-Agent": client.userAgent,
		// App-specific headers
		"CH-AppBuild":   strconv.Itoa(AppBuild),
		"CH-AppVersion": AppVersion,
		"CH-DeviceId":   strings.ToUpper(client.deviceID.String()),
		"CH-Languages":  "en-US",
		"CH-Locale":     "en_US",
		"CH-UserID":     "(null)",
	}
	for key, value := range headers {
		client.sling.Set(key, value)
	}

	// Apply options to client
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// NewFromToken is a convenience function for instantiating a Client with a token.
func NewFromToken(token string, opts ...ClientOption) *Client {
	opts = append([]ClientOption{WithToken(token)}, opts...) // Prepend token option
	return New(opts...)
}

// WithDeviceID is an option for setting the device ID.
func WithDeviceID(deviceID uuid.UUID) ClientOption {
	return func(client *Client) {
		client.deviceID = deviceID
		client.sling.Set("CH-DeviceId", strings.ToUpper(client.deviceID.String()))
	}
}

// WithUserID is an option for setting the auth token.
func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.token = token
		client.sling.Set("Authorization", fmt.Sprintf("Token %s", client.token))
	}
}

// WithUserID is an option for setting the user ID.
func WithUserID(id int) ClientOption {
	return func(client *Client) {
		client.userID = id
		client.sling.Set("CH-UserID", strconv.Itoa(client.userID))
	}
}

// Int returns a pointer to the given int value.
func Int(v int) *int {
	return &v
}

// String returns a pointer to the given string value.
func String(v string) *string {
	return &v
}
