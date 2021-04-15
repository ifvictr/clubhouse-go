package clubhouse

import "net/http"

type CallPhoneNumberAuthParams struct {
	PhoneNumber string `json:"phone_number"`
}

type CallPhoneNumberAuthResponse struct {
	Response
	ErrorMessage *string `json:"error_message"`
}

func (c *Client) CallPhoneNumberAuth(params *CallPhoneNumberAuthParams) (*CallPhoneNumberAuthResponse, *http.Response, error) {
	apiRes := new(CallPhoneNumberAuthResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("call_phone_number_auth").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type CompletePhoneNumberAuthParams struct {
	DeviceToken      *string `json:"device_token,omitempty"`
	PhoneNumber      string  `json:"phone_number"`
	VerificationCode string  `json:"verification_code"`
}

type CompletePhoneNumberAuthResponse struct {
	Response
	AccessToken               *string          `json:"access_token,omitempty"`
	AuthToken                 *string          `json:"auth_token,omitempty"`
	IsOnboarding              *bool            `json:"is_onboarding,omitempty"`
	IsVerified                bool             `json:"is_verified"`
	IsWaitlisted              *bool            `json:"is_waitlisted,omitempty"`
	NumberOfAttemptsRemaining *int             `json:"number_of_attempts_remaining,omitempty"`
	RefreshToken              *string          `json:"refresh_token,omitempty"`
	UserProfile               *BaseUserProfile `json:"user_profile,omitempty"`
}

func (c *Client) CompletePhoneNumberAuth(params *CompletePhoneNumberAuthParams) (*CompletePhoneNumberAuthResponse, *http.Response, error) {
	apiRes := new(CompletePhoneNumberAuthResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("complete_phone_number_auth").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type ResendPhoneNumberAuthParams struct {
	PhoneNumber string `json:"phone_number"`
}

type ResendPhoneNumberAuthResponse struct {
	Response
	ErrorMessage *string `json:"error_message"`
}

func (c *Client) ResendPhoneNumberAuth(params *ResendPhoneNumberAuthParams) (*ResendPhoneNumberAuthResponse, *http.Response, error) {
	apiRes := new(ResendPhoneNumberAuthResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("resend_phone_number_auth").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type StartPhoneNumberAuthParams struct {
	PhoneNumber string `json:"phone_number"`
}

type StartPhoneNumberAuthResponse struct {
	Response
	ErrorMessage *string `json:"error_message"`
	IsBlocked    bool    `json:"is_blocked"`
}

func (c *Client) StartPhoneNumberAuth(params *StartPhoneNumberAuthParams) (*StartPhoneNumberAuthResponse, *http.Response, error) {
	apiRes := new(StartPhoneNumberAuthResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("start_phone_number_auth").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
