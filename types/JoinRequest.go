package types

type JoinRequest struct {
	Type string          `json:"type"`
	Data JoinRequestData `json:"data"`
}

type JoinRequestData struct {
	ChannelId string `json:"channel_id"`
	Hidden    int    `json:"hidden"`
	Mobile    bool   `json:"mobile"`
	Reload    bool   `json:"reload"`
}
