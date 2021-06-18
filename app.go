package clubhouse

import "net/http"

type CheckForUpdateParams struct {
	IsTestFlight int `url:"is_testflight"`
}

type CheckForUpdateResponse struct {
	Response
	AppBuild    *int    `json:"app_build,omitempty"`
	AppURL      *string `json:"app_url,omitempty"`
	AppVersion  *string `json:"app_version,omitempty"`
	HasUpdate   bool    `json:"has_update"`
	IsMandatory *bool   `json:"is_mandatory,omitempty"`
}

func (c *Client) CheckForUpdate(params *CheckForUpdateParams) (*CheckForUpdateResponse, *http.Response, error) {
	apiRes := new(CheckForUpdateResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("check_for_update").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}

type GetSettingsResponse struct {
	Response
	NotificationsEnableOther     bool `json:"notifications_enable_other"`
	NotificationsEnableRoom      bool `json:"notifications_enable_room"`
	NotificationsEnableSendFewer bool `json:"notifications_enable_send_fewer"`
	NotificationsEnableTrending  bool `json:"notifications_enable_trending"`
	NotificationsFrequency       int  `json:"notifications_frequency"`
	NotificationsIsPaused        bool `json:"notifications_is_paused"`
}

func (c *Client) GetSettings() (*GetSettingsResponse, *http.Response, error) {
	apiRes := new(GetSettingsResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("get_settings").Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
