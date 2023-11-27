package types

type ChannelCountersData struct {
	ChannelId        string `json:"channel_id"`
	ClientsInChannel int    `json:"clients_in_channel"`
	UsersInChannel   int    `json:"users_in_channel"`
}
