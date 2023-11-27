package main

import (
	"YudolePlatofrmGoodgameClient/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

var Out = make(chan any, 99999)

func processMessage(message types.ChatMessage) types.ChatMessage {

	return message
}

func Connect() {
	conn, _, err := websocket.DefaultDialer.Dial("wss://chat-1.goodgame.ru/chat2/", nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer conn.Close()

	for {
		var message types.Base
		conn.ReadJSON(&message)

		switch message.Type {
		case "welcome":
			var joinRequest = types.JoinRequest{
				Type: "join",
				Data: types.JoinRequestData{
					ChannelId: "53029",
					Hidden:    0,
					Mobile:    false,
					Reload:    false,
				},
			}

			conn.WriteJSON(joinRequest)
			break

		case "success_join":
			var successJoinData types.SuccessJoinData
			json.Unmarshal(message.Data, &successJoinData)
			// Do nothing
			break

		case "message":
			var messageData types.MessageData
			json.Unmarshal(message.Data, &messageData)

			var m = types.ChatMessage{
				Type: "chat/message",
				Src:  messageData.Text,
				Text: "",
				Html: "",
				User: types.User{
					Name: messageData.UserName,
					Meta: types.UserMeta{
						Avatar: "",
						Badges: "",
					},
				},
			}

			Out <- processMessage(m)
			break

		case "channel_counters":
			var channelCountersData types.ChannelCountersData
			json.Unmarshal(message.Data, &channelCountersData)
			break

		case "private_message":

			break

		case "premium":

			break

		case "gifted_premiums":

			break

		default:
			break
		}
	}
}
