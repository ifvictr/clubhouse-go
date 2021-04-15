package clubhouse

type APIError struct {
	Response
	Detail       *string `json:"detail"`
	ErrorMessage *string `json:"error_message"`
}

func (e APIError) Error() string {
	if e.Detail != nil {
		return *e.Detail
	}
	if e.ErrorMessage != nil {
		return *e.ErrorMessage
	}
	return ""
}

func (e APIError) Empty() bool {
	return e.Detail == nil && e.ErrorMessage == nil
}

func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if !apiError.Empty() {
		return apiError
	}
	return nil
}
