package types

type MessageData struct {
	ChannelId  string   `json:"channel_id"`
	UserId     int      `json:"user_id"`
	UserName   string   `json:"user_name"`
	UserRights int      `json:"user_rights"`
	Premium    int      `json:"premium"`
	Premiums   []string `json:"premiums"`
	Staff      int      `json:"staff"`
	Color      string   `json:"color"`
	Icon       string   `json:"icon"`
	Role       string   `json:"role"`
	Mobile     int      `json:"mobile"`
	GgPlusTier int      `json:"gg_plus_tier"`
	IsStatus   int      `json:"isStatus"`
	MessageId  int64    `json:"message_id"`
	Timestamp  int      `json:"timestamp"`
	Text       string   `json:"text"`
	Regtime    int      `json:"regtime"`
}
