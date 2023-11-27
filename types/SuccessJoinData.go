package types

type SuccessJoinData struct {
	ChannelId   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelKey  string `json:"channel_key"`
	Motd        string `json:"motd"`
}
