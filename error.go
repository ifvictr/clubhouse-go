package clubhouse

type ErrorResponse struct {
	Response
	Detail       *string `json:"detail"`
	ErrorMessage *string `json:"error_message"`
}

func (e ErrorResponse) Error() string {
	if e.Detail != nil {
		return *e.Detail
	}
	if e.ErrorMessage != nil {
		return *e.ErrorMessage
	}
	return ""
}

func (e ErrorResponse) Empty() bool {
	return e.Detail == nil && e.ErrorMessage == nil
}

func relevantError(httpError error, apiError ErrorResponse) error {
	if httpError != nil {
		return httpError
	}
	if !apiError.Empty() {
		return apiError
	}
	return nil
}
