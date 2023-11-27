package main

import (
	"YudolePlatofrmGoodgameClient/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

var Out = make(chan any, 99999)

func processMessage(message types.ChatMessage) types.ChatMessage {
	message.Text = message.Src
	message.Html = message.Src

	for _, v := range smiles {
		var smile = ":" + v.Key + ":"
		var image = v.Images.Big

		if v.Animated == 1 {
			image = v.Images.Gif
		}

		if strings.Index(message.Src, smile) >= 0 {
			message.Text = strings.ReplaceAll(message.Text, smile, "")
			message.Html = strings.ReplaceAll(message.Html, smile, fmt.Sprintf("<img src=\"%s\" alt=\"%s\" title=\"%s\"/>", image, v.Key, v.Key))
		}
	}

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
