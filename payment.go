package clubhouse

import "net/http"

type GetStripeEphemeralKeyParams struct {
	StripeVersion string `json:"stripe_version"`
}

type GetStripeEphemeralKeyResponse struct {
	Response
	AssociatedObjects []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"associated_objects"`
	Created  int    `json:"created"`
	Expires  int    `json:"expires"`
	ID       string `json:"id"`
	Livemode bool   `json:"livemode"`
	Object   string `json:"object"`
	Secret   string `json:"secret"`
}

func (c *Client) GetStripeEphemeralKey(params *GetStripeEphemeralKeyParams) (*GetStripeEphemeralKeyResponse, *http.Response, error) {
	apiRes := new(GetStripeEphemeralKeyResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("get_stripe_ephemeral_key").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type SendDirectPaymentParams struct {
	Amount          int     `json:"amount"` // Amount is in the smallest common currency unit. E.g. $1.00 = 100
	Channel         *string `json:"channel"`
	PaymentMethodID string  `json:"payment_method_id"`
	ReceivingUserID int     `json:"receiving_user_id"`
}

type SendDirectPaymentResponse struct {
	Response
	ClientSecret     *string     `json:"client_secret"`
	ConnectAccountID interface{} `json:"connect_account_id"` // TODO: Find real type
	NeedsAction      bool        `json:"needs_action"`
}

func (c *Client) SendDirectPayment(params *SendDirectPaymentParams) (*SendDirectPaymentResponse, *http.Response, error) {
	apiRes := new(SendDirectPaymentResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Post("send_direct_payment").BodyJSON(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
