package clubhouse

type Response struct {
	Success bool `json:"success"`
}

type PageResponse struct {
	Response
	Count    int  `json:"count"`
	Next     *int `json:"next"`
	Previous *int `json:"previous"`
}
