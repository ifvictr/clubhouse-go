package clubhouse

import "net/http"

type Event struct {
	Channel      *string     `json:"channel"`
	Club         *Club       `json:"club"`
	Description  string      `json:"description"`
	EventID      int         `json:"event_id"`
	Hosts        []EventHost `json:"hosts"`
	IsAttending  bool        `json:"is_attending"`
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

type GetEventsForClubParams struct {
	ClubID              int         `json:"club_id"`
	InviteCode          interface{} `json:"invite_code"` // TODO: Find real type
	QueryID             *string     `json:"query_id"`
	QueryResultPosition *int        `json:"query_result_position"`
	Slug                *string     `json:"slug"`
	SourceTopicID       *int        `json:"source_topic_id"`
}

type GetEventsForClubResponse struct {
	Response
	Events []Event `json:"events"`
}

func (c *Client) GetEventsForClub(params *GetEventsForClubParams) (*GetEventsForClubResponse, *http.Response, error) {
	apiRes := new(GetEventsForClubResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("get_events_for_club").BodyJSON(params).Receive(apiRes, apiError)
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

type RsvpEventParams struct {
	EventID     int  `json:"event_id"`
	IsAttending bool `json:"is_attending"`
}

type RsvpEventResponse struct {
	Response
}

func (c *Client) RsvpEvent(params *RsvpEventParams) (*RsvpEventResponse, *http.Response, error) {
	apiRes := new(RsvpEventResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("rsvp_event").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
