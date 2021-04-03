package clubhouse

import (
	"net/http"
	"time"
)

type Channel struct {
	Channel              string        `json:"channel"`
	ChannelID            int           `json:"channel_id"`
	Club                 *Club         `json:"club"`
	ClubID               *int          `json:"club_id"`
	ClubName             *string       `json:"club_name"`
	CreatorUserProfileID int           `json:"creator_user_profile_id"`
	FeatureFlags         []FeatureFlag `json:"feature_flags"`
	HasBlockedSpeakers   bool          `json:"has_blocked_speakers"`
	IsExploreChannel     bool          `json:"is_explore_channel"`
	IsPrivate            bool          `json:"is_private"`
	IsSocialMode         bool          `json:"is_social_mode"`
	NumAll               int           `json:"num_all"`
	NumOther             int           `json:"num_other"`
	NumSpeakers          int           `json:"num_speakers"`
	Topic                *string       `json:"topic"`
	URL                  string        `json:"url"`
	Users                []struct {
		IsModerator         bool       `json:"is_moderator"`
		IsSpeaker           bool       `json:"is_speaker"`
		Name                string     `json:"name"`
		PhotoURL            string     `json:"photo_url"`
		TimeJoinedAsSpeaker *time.Time `json:"time_joined_as_speaker"`
		UserID              int        `json:"user_id"`
	} `json:"users"`
	WelcomeForUserProfile *BaseUserProfile `json:"welcome_for_user_profile"`
}

type FeatureFlag string

const (
	AgoraAudioProfileSpeechStandard FeatureFlag = "AGORA_AUDIO_PROFILE_SPEECH_STANDARD"
)

type ChannelUser struct {
	BaseUserProfile
	FirstName           string     `json:"first_name"`
	IsFollowedBySpeaker bool       `json:"is_followed_by_speaker"`
	IsInvitedAsSpeaker  bool       `json:"is_invited_as_speaker"`
	IsModerator         bool       `json:"is_moderator"`
	IsNew               bool       `json:"is_new"`
	IsSpeaker           bool       `json:"is_speaker"`
	Skintone            int        `json:"skintone"`
	TimeJoinedAsSpeaker *time.Time `json:"time_joined_as_speaker"`
}

type BlockFromChannelParams struct {
	Channel string `json:"channel"`
	UserID  int    `json:"user_id"`
}

type BlockFromChannelResponse struct {
	Response
}

func (c *Client) BlockFromChannel(params *BlockFromChannelParams) (*BlockFromChannelResponse, *http.Response, error) {
	apiRes := new(BlockFromChannelResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("block_from_channel").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type CreateChannelParams struct {
	ClubID       *int    `json:"club_id"`
	EventID      *int    `json:"event_id"`
	IsPrivate    bool    `json:"is_private"`
	IsSocialMode bool    `json:"is_social_mode"`
	Topic        *string `json:"topic"`
	UserIds      []int   `json:"user_ids"`
}

type CreateChannelResponse struct {
	Response
	AgoraNativeMute         bool             `json:"agora_native_mute"`
	Channel                 string           `json:"channel"`
	ChannelID               int              `json:"channel_id"`
	Club                    *Club            `json:"club"`
	ClubID                  *int             `json:"club_id"`
	ClubName                *string          `json:"club_name"`
	CreatorUserProfileID    int              `json:"creator_user_profile_id"`
	FeatureFlags            []string         `json:"feature_flags"`
	HandraisePermission     int              `json:"handraise_permission"`
	IsClubAdmin             bool             `json:"is_club_admin"`
	IsClubMember            bool             `json:"is_club_member"`
	IsHandraiseEnabled      bool             `json:"is_handraise_enabled"`
	IsPrivate               bool             `json:"is_private"`
	IsSocialMode            bool             `json:"is_social_mode"`
	PubnubEnable            bool             `json:"pubnub_enable"`
	PubnubHeartbeatInterval int              `json:"pubnub_heartbeat_interval"`
	PubnubHeartbeatValue    int              `json:"pubnub_heartbeat_value"`
	PubnubOrigin            string           `json:"pubnub_origin"`
	PubnubToken             string           `json:"pubnub_token"`
	RtmToken                string           `json:"rtm_token"`
	Token                   string           `json:"token"`
	Topic                   *string          `json:"topic"`
	URL                     string           `json:"url"`
	Users                   []ChannelUser    `json:"users"`
	WelcomeForUserProfile   *BaseUserProfile `json:"welcome_for_user_profile"`
}

func (c *Client) CreateChannel(params *CreateChannelParams) (*CreateChannelResponse, *http.Response, error) {
	apiRes := new(CreateChannelResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("create_channel").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type EndChannelParams struct {
	Channel   string `json:"channel"`
	ChannelID *int   `json:"channel_id"`
}

type EndChannelResponse struct {
	Response
}

func (c *Client) EndChannel(params *EndChannelParams) (*EndChannelResponse, *http.Response, error) {
	apiRes := new(EndChannelResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("end_channel").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetChannelParams struct {
	Channel   string `json:"channel"`
	ChannelID *int   `json:"channel_id"`
}

type GetChannelResponse struct {
	Response
	Channel               string           `json:"channel"`
	ChannelID             int              `json:"channel_id"`
	Club                  *Club            `json:"club"`
	ClubID                *int             `json:"club_id"`
	ClubName              *string          `json:"club_name"`
	CreatorUserProfileID  int              `json:"creator_user_profile_id"`
	FeatureFlags          []FeatureFlag    `json:"feature_flags"`
	HandraisePermission   int              `json:"handraise_permission"`
	IsClubAdmin           bool             `json:"is_club_admin"`
	IsClubMember          bool             `json:"is_club_member"`
	IsHandraiseEnabled    bool             `json:"is_handraise_enabled"`
	IsPrivate             bool             `json:"is_private"`
	IsSocialMode          bool             `json:"is_social_mode"`
	ShouldLeave           bool             `json:"should_leave"`
	Topic                 *string          `json:"topic"`
	URL                   string           `json:"url"`
	Users                 []ChannelUser    `json:"users"`
	WelcomeForUserProfile *BaseUserProfile `json:"welcome_for_user_profile"`
}

func (c *Client) GetChannel(params *GetChannelParams) (*GetChannelResponse, *http.Response, error) {
	apiRes := new(GetChannelResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("get_channel").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetChannelsResponse struct {
	Response
	Channels []Channel `json:"channels"`
	Events   []Event   `json:"events"`
}

func (c *Client) GetChannels() (*GetChannelsResponse, *http.Response, error) {
	apiRes := new(GetChannelsResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Get("get_channels").Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type JoinChannelParams struct {
	AttributionDetails string            `json:"attribution_details"` // Base 64 JSON
	AttributionSource  AttributionSource `json:"attribution_source"`
	Channel            string            `json:"channel"`
}

type AttributionSource string

const (
	AttributionSourceActivity AttributionSource = "activity"
	AttributionSourceFeed     AttributionSource = "feed"
)

type JoinChannelResponse struct {
	Response
	AgoraNativeMute         bool             `json:"agora_native_mute"`
	Channel                 string           `json:"channel"`
	ChannelID               int              `json:"channel_id"`
	Club                    *Club            `json:"club"`
	ClubID                  *int             `json:"club_id"`
	ClubName                *string          `json:"club_name"`
	CreatorUserProfileID    int              `json:"creator_user_profile_id"`
	FeatureFlags            []string         `json:"feature_flags"`
	HandraisePermission     int              `json:"handraise_permission"`
	IsClubAdmin             bool             `json:"is_club_admin"`
	IsClubMember            bool             `json:"is_club_member"`
	IsEmpty                 bool             `json:"is_empty"`
	IsHandraiseEnabled      bool             `json:"is_handraise_enabled"`
	IsPrivate               bool             `json:"is_private"`
	IsSocialMode            bool             `json:"is_social_mode"`
	PubnubEnable            bool             `json:"pubnub_enable"`
	PubnubHeartbeatInterval int              `json:"pubnub_heartbeat_interval"`
	PubnubHeartbeatValue    int              `json:"pubnub_heartbeat_value"`
	PubnubOrigin            string           `json:"pubnub_origin"`
	PubnubToken             string           `json:"pubnub_token"`
	RtmToken                string           `json:"rtm_token"`
	Token                   string           `json:"token"`
	Topic                   *string          `json:"topic"`
	URL                     string           `json:"url"`
	Users                   []ChannelUser    `json:"users"`
	WelcomeForUserProfile   *BaseUserProfile `json:"welcome_for_user_profile"`
}

func (c *Client) JoinChannel(params *JoinChannelParams) (*JoinChannelResponse, *http.Response, error) {
	apiRes := new(JoinChannelResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("join_channel").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type LeaveChannelParams struct {
	Channel   string `json:"channel"`
	ChannelID *int   `json:"channel_id"`
}

type LeaveChannelResponse struct {
	Response
}

func (c *Client) LeaveChannel(params *LeaveChannelParams) (*LeaveChannelResponse, *http.Response, error) {
	apiRes := new(LeaveChannelResponse)
	apiError := new(ErrorResponse)
	res, err := c.sling.New().Post("leave_channel").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
