package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

// メイン関数は、WebSocketサーバーを起動し、UDPパケットを処理するための関数をゴルーチンとして実行します。
func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleUDP()
	go handleMessages()
	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// handleConnections関数は、WebSocket接続を処理し、クライアントからのメッセージを受信します。
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			delete(clients, ws)
			break
		}
		fmt.Println("Received message:", string(msg))
	}
}

// handleUDP関数は、UDPサーバーを起動し、受信したUDPパケットをブロードキャストチャネルに送信します。
func handleUDP() {
	addr := net.UDPAddr{
		Port: 8081,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading UDP packet:", err)
			continue
		}
		fmt.Println("Received UDP packet:", string(buf[:n]))
		broadcast <- buf[:n]
	}
}

// handleMessages関数は、ブロードキャストチャネルからメッセージを受信し、すべてのクライアントに送信します。
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Error writing message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
