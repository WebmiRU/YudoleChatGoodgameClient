package main

import (
	"YudolePlatofrmGoodgameClient/types"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var smiles []types.Smile

func main() {
	smilesData, err := os.ReadFile("./data/goodgame/smiles.json")

	if err != nil {
		log.Println("Error while loading smiles")
	}

	if err := json.Unmarshal(smilesData, &smiles); err != nil {
		log.Println("Error while encoding JSON smiles data")
	}

	go Connect()

	conn, _ := net.Dial("tcp", "127.0.0.1:5801")

	fmt.Println("CONNECTED!")

	for {
		var out = <-Out
		message, _ := json.Marshal(out)
		fmt.Println("OUT:", string(message))

		if _, err := conn.Write(message); err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			break
		}
	}

	// @TODO
	fmt.Println("RECONNECT HERE")

	//for {
	//	decoder := json.NewDecoder(conn)
	//
	//	var message json.RawMessage
	//	if err := decoder.Decode(&message); err != nil {
	//		break
	//	}
	//
	//}
}
