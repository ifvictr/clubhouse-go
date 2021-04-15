package clubhouse

import "net/http"

type CheckForUpdateParams struct {
	IsTestFlight int `url:"is_testflight"`
}

type CheckForUpdateResponse struct {
	Response
	HasUpdate bool `json:"has_update"`
}

func (c *Client) CheckForUpdate(params *CheckForUpdateParams) (*CheckForUpdateResponse, *http.Response, error) {
	apiRes := new(CheckForUpdateResponse)
	apiError := new(APIError)
	res, err := c.sling.New().Get("check_for_update").QueryStruct(params).Receive(apiRes, apiError)
	return apiRes, res, relevantError(err, *apiError)
}
