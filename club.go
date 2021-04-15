package clubhouse

import "net/http"

type BaseClub struct {
	ClubID      int     `json:"club_id"`
	Description *string `json:"description"`
	Name        string  `json:"name"`
	PhotoURL    *string `json:"photo_url"`
}

type Club struct {
	BaseClub
	EnablePrivate       bool   `json:"enable_private"`
	IsCommunity         bool   `json:"is_community"`
	IsFollowAllowed     bool   `json:"is_follow_allowed"`
	IsMembershipPrivate bool   `json:"is_membership_private"`
	NumFollowers        int    `json:"num_followers"`
	NumMembers          int    `json:"num_members"`
	NumOnline           int    `json:"num_online"`
	Rules               []Rule `json:"rules"`
	Url                 string `json:"url"`
}

type ClubMember struct {
	BaseUserProfile
	Bio               *string `json:"bio"`
	IsAdmin           bool    `json:"is_admin"`
	IsFollower        bool    `json:"is_follower"`
	IsMember          bool    `json:"is_member"`
	IsPendingAccept   bool    `json:"is_pending_accept"`
	IsPendingApproval bool    `json:"is_pending_approval"`
}

type SearchClub struct {
	BaseClub
	ClubID int `json:"club_id,string"` // Club IDs are returned as strings in search results
}

type Rule struct {
	Desc  string `json:"desc"`
	Title string `json:"title"`
}

type FollowClubParams struct {
	ClubID        int    `json:"club_id,omitempty"`
	Slug          string `json:"slug,omitempty"`
	SourceTopicID int    `json:"source_topic_id,omitempty"`
}

type FollowClubResponse struct {
	Response
}

func (c *Client) FollowClub(params *FollowClubParams) (*FollowClubResponse, *http.Response, error) {
	apiRes := new(FollowClubResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("follow_club").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetClubParams struct {
	ClubID        int    `json:"club_id,omitempty"`
	Slug          string `json:"slug,omitempty"`
	SourceTopicID int    `json:"source_topic_id,omitempty"`
}

type GetClubResponse struct {
	Response
	AddedByUserProfile *BaseUserProfile `json:"added_by_user_profile"`
	Club               Club             `json:"club"`
	InviteLink         *string          `json:"invite_link"`
	IsAdmin            bool             `json:"is_admin"`
	IsFollower         bool             `json:"is_follower"`
	IsMember           bool             `json:"is_member"`
	IsPendingAccept    bool             `json:"is_pending_accept"`
	IsPendingApproval  bool             `json:"is_pending_approval"`
	MemberUserIds      []int            `json:"member_user_ids"`
	NumInvites         int              `json:"num_invites"`
	Topics             []Topic          `json:"topics"`
}

func (c *Client) GetClub(params *GetClubParams) (*GetClubResponse, *http.Response, error) {
	apiRes := new(GetClubResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("get_club").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetClubMembersParams struct {
	ClubID          int `url:"club_id"`
	ReturnFollowers int `url:"return_followers"`
	ReturnMembers   int `url:"return_members"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type GetClubMembersResponse struct {
	PageResponse
	Users []ClubMember `json:"users"`
}

func (c *Client) GetClubMembers(params *GetClubMembersParams) (*GetClubMembersResponse, *http.Response, error) {
	apiRes := new(GetClubMembersResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("get_club_members").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetClubsParams struct {
	IsStartableOnly bool `json:"is_startable_only"`
}

type GetClubsResponse struct {
	Response
	Clubs []struct {
		Club
		IsAdmin bool `json:"is_admin"`
	} `json:"clubs"`
}

func (c *Client) GetClubs(params *GetClubsParams) (*GetClubsResponse, *http.Response, error) {
	apiRes := new(GetClubsResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("get_clubs").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type SearchClubsParams struct {
	CofollowsOnly bool   `json:"cofollows_only" url:"-"`
	FollowersOnly bool   `json:"followers_only" url:"-"`
	FollowingOnly bool   `json:"following_only" url:"-"`
	Query         string `json:"query" url:"-"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type SearchClubsResponse struct {
	PageResponse
	Clubs   []SearchClub `json:"clubs"`
	QueryID string       `json:"query_id"`
}

func (c *Client) SearchClubs(params *SearchClubsParams) (*SearchClubsResponse, *http.Response, error) {
	apiRes := new(SearchClubsResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("search_clubs").BodyJSON(params).QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type UnfollowClubParams struct {
	ClubID        int    `json:"club_id,omitempty"`
	Slug          string `json:"slug,omitempty"`
	SourceTopicID int    `json:"source_topic_id,omitempty"`
}

type UnfollowClubResponse struct {
	Response
}

func (c *Client) UnfollowClub(params *UnfollowClubParams) (*UnfollowClubResponse, *http.Response, error) {
	apiRes := new(UnfollowClubResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("unfollow_club").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
