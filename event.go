package clubhouse

import "net/http"

type Event struct {
	Channel      *string     `json:"channel"`
	Club         *Club       `json:"club"`
	Description  string      `json:"description"`
	EventID      int         `json:"event_id"`
	Hosts        []EventHost `json:"hosts"`
	IsExpired    bool        `json:"is_expired"`
	IsMemberOnly bool        `json:"is_member_only"`
	Name         string      `json:"name"`
	TimeStart    string      `json:"time_start"`
	URL          string      `json:"url"`
}

type EventHost struct {
	BaseUserProfile
	Bio     *string `json:"bio"`
	Twitter *string `json:"twitter"`
}

type GetEventsParams struct {
	IsFiltered bool `url:"is_filtered"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type GetEventsResponse struct {
	PageResponse
	Events []struct {
		Event
		ClubIsFollower int `json:"club_is_follower"`
		ClubIsMember   int `json:"club_is_member"`
	} `json:"events"`
}

func (c *Client) GetEvents(params *GetEventsParams) (*GetEventsResponse, *http.Response, error) {
	apiRes := new(GetEventsResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("get_events").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetEventsForUserParams struct {
	UserID int `url:"user_id"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type GetEventsForUserResponse struct {
	PageResponse
	Events []struct {
		Event
		ClubIsFollower int `json:"club_is_follower"`
		ClubIsMember   int `json:"club_is_member"`
	} `json:"events"`
}

func (c *Client) GetEventsForUser(params *GetEventsForUserParams) (*GetEventsForUserResponse, *http.Response, error) {
	apiRes := new(GetEventsForUserResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("get_events_for_user").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
