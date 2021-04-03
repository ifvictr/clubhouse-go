package clubhouse

type Event struct {
	Channel      *string     `json:"channel"`
	Club         *Club       `json:"club"`
	Description  string      `json:"description"`
	EventID      int         `json:"event_id"`
	Hosts        []EventHost `json:"hosts"`
	IsExpired    bool        `json:"is_expired"`
	IsMemberOnly bool        `json:"is_member_only"`
	Name         string      `json:"name"`
	TimeStart    string      `json:"time_start"`
	URL          string      `json:"url"`
}

type EventHost struct {
	BaseUserProfile
	Bio     *string `json:"bio"`
	Twitter *string `json:"twitter"`
}

// TODO
