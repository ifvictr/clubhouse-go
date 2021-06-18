package clubhouse

import "net/http"

type Notification struct {
	Channel        *string     `json:"channel,omitempty"`
	EventID        *int        `json:"event_id,omitempty"`
	IsUnread       bool        `json:"is_unread"`
	Message        string      `json:"message"`
	NotificationID int         `json:"notification_id"`
	TimeCreated    string      `json:"time_created"`
	Type           int         `json:"type"`
	UserProfile    UserProfile `json:"user_profile"`
}

type GetActionableNotificationsResponse struct {
	PageResponse
	Notifications []Notification `json:"notifications"`
}

func (c *Client) GetActionableNotifications() (*GetNotificationsResponse, *http.Response, error) {
	apiRes := new(GetNotificationsResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("get_actionable_notifications").Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetNotificationsParams struct {
	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type GetNotificationsResponse struct {
	Response
	Count         int            `json:"count"`
	Disabled      bool           `json:"disabled"`
	Notifications []Notification `json:"notifications"`
}

func (c *Client) GetNotifications(params *GetNotificationsParams) (*GetNotificationsResponse, *http.Response, error) {
	apiRes := new(GetNotificationsResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("get_notifications").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
