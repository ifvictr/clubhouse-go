package clubhouse

import (
	"net/http"
	"time"
)

type BaseUserProfile struct {
	Name     *string `json:"name"`
	PhotoURL *string `json:"photo_url"`
	UserID   int     `json:"user_id"`
	Username *string `json:"username"`
}

type FollowerUserProfile struct {
	BaseUserProfile
	Bio               *string `json:"bio"`
	LastActiveMinutes int     `json:"last_active_minutes"`
	Twitter           *string `json:"twitter"`
}

type FollowingUserProfile struct {
	BaseUserProfile
	Bio     *string `json:"bio"`
	Twitter *string `json:"twitter"`
}

type SearchUserProfile struct {
	BaseUserProfile
	Bio    *string `json:"bio"`
	UserID int     `json:"user_id,string"` // User IDs are returned as strings in search results
}

type UserProfile struct {
	BaseUserProfile
	Bio                     *string           `json:"bio"`
	CanEditDisplayname      bool              `json:"can_edit_displayname"`
	CanEditName             bool              `json:"can_edit_name"`
	CanEditUsername         bool              `json:"can_edit_username"`
	CanReceiveDirectPayment bool              `json:"can_receive_direct_payment"`
	Clubs                   []Club            `json:"clubs"`
	DirectPaymentFeeFixed   float64           `json:"direct_payment_fee_fixed"`
	DirectPaymentFeeRate    float64           `json:"direct_payment_fee_rate"`
	Displayname             *string           `json:"displayname"`
	FollowsMe               bool              `json:"follows_me"`
	HasVerifiedEmail        bool              `json:"has_verified_email"`
	Instagram               *string           `json:"instagram"`
	InvitedByClub           *Club             `json:"invited_by_club"`
	InvitedByUserProfile    *BaseUserProfile  `json:"invited_by_user_profile"`
	IsBlockedByNetwork      bool              `json:"is_blocked_by_network"`
	MutualFollows           []BaseUserProfile `json:"mutual_follows"`
	MutualFollowsCount      int               `json:"mutual_follows_count"`
	NotificationType        *int              `json:"notification_type"`
	NumFollowers            int               `json:"num_followers"`
	NumFollowing            int               `json:"num_following"`
	TimeCreated             time.Time         `json:"time_created"`
	Topics                  []Topic           `json:"topics"`
	Twitter                 *string           `json:"twitter"`
	Url                     *string           `json:"url"`
}

type AddEmailParams struct {
	Email string `json:"email"`
}

type AddEmailResponse struct {
	Response
}

func (c *Client) AddEmail(params *UpdateNameParams) (*AddEmailResponse, *http.Response, error) {
	apiRes := new(AddEmailResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("add_email").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type FollowParams struct {
	Source        int    `json:"source"` // TODO: Document all sources
	SourceTopicID *int   `json:"source_topic_id"`
	UserIDS       *[]int `json:"user_ids"`
	UserID        int    `json:"user_id"`
}

type FollowResponse struct {
	Response
}

func (c *Client) Follow(params *FollowParams) (*FollowResponse, *http.Response, error) {
	apiRes := new(FollowResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("follow").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetFollowersParams struct {
	UserID int `url:"user_id"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type GetFollowersResponse struct {
	PageResponse
	Users []FollowerUserProfile `json:"users"`
}

func (c *Client) GetFollowers(params *GetFollowersParams) (*GetFollowersResponse, *http.Response, error) {
	apiRes := new(GetFollowersResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Get("get_followers").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetFollowingParams struct {
	UserID int `url:"user_id"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type GetFollowingResponse struct {
	PageResponse
	Clubs []Club                 `json:"clubs"`
	Users []FollowingUserProfile `json:"users"`
}

func (c *Client) GetFollowing(params *GetFollowingParams) (*GetFollowingResponse, *http.Response, error) {
	apiRes := new(GetFollowingResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Get("get_following").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetProfileParams struct {
	UserID   int    `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type GetProfileResponse struct {
	Response
	UserProfile UserProfile `json:"user_profile"`
}

func (c *Client) GetProfile(params *GetProfileParams) (*GetProfileResponse, *http.Response, error) {
	apiRes := new(GetProfileResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("get_profile").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type InviteToAppParams struct {
	Message     *string `json:"message"`
	Name        *string `json:"name"`
	PhoneNumber string  `json:"phone_number"`
}

type InviteToAppResponse struct {
	Response
}

func (c *Client) InviteToApp(params *InviteToAppParams) (*InviteToAppResponse, *http.Response, error) {
	apiRes := new(InviteToAppResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("invite_to_app").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type MeParams struct {
	ReturnBlockedIds   bool   `json:"return_blocked_ids"`
	TimezoneIdentifier string `json:"timezone_identifier"`
	ReturnFollowingIds bool   `json:"return_following_ids"`
}

type MeResponse struct {
	Response
	AccessToken                  string   `json:"access_token"`
	ActionableNotificationsCount int      `json:"actionable_notifications_count"`
	AuthToken                    string   `json:"auth_token"`
	BlockedIds                   *[]int   `json:"blocked_ids"`
	Email                        string   `json:"email"`
	FeatureFlags                 []string `json:"feature_flags"`
	FollowingIds                 *[]int   `json:"following_ids"`
	HasUnreadNotifications       bool     `json:"has_unread_notifications"`
	IsAdmin                      bool     `json:"is_admin"`
	NotificationsEnabled         bool     `json:"notifications_enabled"`
	NumInvites                   int      `json:"num_invites"`
	RefreshToken                 string   `json:"refresh_token"`
	ServiceStatus                *struct {
		ButtonTitle string `json:"button_title"`
		Message     string `json:"message"`
		Title       string `json:"title"`
		URL         string `json:"url"`
	} `json:"service_status"`
	UserProfile BaseUserProfile `json:"user_profile"`
}

func (c *Client) Me(params *MeParams) (*MeResponse, *http.Response, error) {
	apiRes := new(MeResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("me").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type SearchUsersParams struct {
	CofollowsOnly bool   `json:"cofollows_only" url:"-"`
	FollowersOnly bool   `json:"followers_only" url:"-"`
	FollowingOnly bool   `json:"following_only" url:"-"`
	Query         string `json:"query" url:"-"`

	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type SearchUsersResponse struct {
	PageResponse
	QueryID string              `json:"query_id"`
	Users   []SearchUserProfile `json:"users"`
}

func (c *Client) SearchUsers(params *SearchUsersParams) (*SearchUsersResponse, *http.Response, error) {
	apiRes := new(SearchUsersResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("search_users").BodyJSON(params).QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type UnfollowParams struct {
	UserID   int    `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type UnfollowResponse struct {
	Response
}

func (c *Client) Unfollow(params *UnfollowParams) (*UnfollowResponse, *http.Response, error) {
	apiRes := new(UnfollowResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("unfollow").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type UpdateBioParams struct {
	Bio string `json:"bio"`
}

type UpdateBioResponse struct {
	Response
}

func (c *Client) UpdateBio(params *UpdateBioParams) (*UpdateBioResponse, *http.Response, error) {
	apiRes := new(UpdateBioResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("update_bio").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type UpdateNameParams struct {
	Name string `json:"name"`
}

type UpdateNameResponse struct {
	Response
	ErrorMessage *string `json:"error_message"`
}

func (c *Client) UpdateName(params *UpdateNameParams) (*UpdateNameResponse, *http.Response, error) {
	apiRes := new(UpdateNameResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("update_name").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type UpdateUsernameParams struct {
	TwitterToken  *string `json:"twitter_token"`
	TwitterSecret *string `json:"twitter_secret"`
	Username      string  `json:"username"`
}

type UpdateUsernameResponse struct {
	Response
	ErrorMessage *string `json:"error_message"`
}

func (c *Client) UpdateUsername(params *UpdateUsernameParams) (*UpdateUsernameResponse, *http.Response, error) {
	apiRes := new(UpdateUsernameResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("update_username").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
