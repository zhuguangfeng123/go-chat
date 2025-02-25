package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

var (
	clientMap = make(map[int64]*Node)
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err)
				return
			}

		}
	}
}

func readProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(data))
	}
}

func SendMSg(userID int64, msg []byte) {
	node, ok := clientMap[userID]
	if !ok {
		log.Println("")
		return
	}
	node.DataQueue <- msg
}
