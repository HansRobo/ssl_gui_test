package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/HansRobo/ssl_gui_test/pkg/ssl_messages"
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
		Port: 10006,
		IP:   net.ParseIP("224.5.23.2"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 65535)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading UDP packet:", err)
			continue
		}
		var packet ssl_messages.SSL_WrapperPacket
		err = proto.Unmarshal(buf[:n], &packet)
		if err != nil {
			fmt.Println("Error unmarshalling packet:", err)
			continue
		}
		shapeInfo := extractShapeInfo(packet)
		shapeData, err := json.Marshal(shapeInfo)
		if err != nil {
			fmt.Println("Error marshalling shape info:", err)
			continue
		}
		broadcast <- shapeData
	}
}

// extractShapeInfo関数は、SSL_WrapperPacketから図形情報を抽出します。
func extractShapeInfo(packet ssl_messages.SSL_WrapperPacket) []map[string]interface{} {
	var shapes []map[string]interface{}
	if packet.Geometry != nil {
		for _, line := range packet.Geometry.Field.FieldLines {
			shape := map[string]interface{}{
				"type":      "line",
				"x1":        line.P1.X,
				"y1":        line.P1.Y,
				"x2":        line.P2.X,
				"y2":        line.P2.Y,
				"thickness": line.Thickness,
			}
			shapes = append(shapes, shape)
		}
		for _, arc := range packet.Geometry.Field.FieldArcs {
			shape := map[string]interface{}{
				"type":      "arc",
				"centerX":   arc.Center.X,
				"centerY":   arc.Center.Y,
				"radius":    arc.Radius,
				"startAngle": arc.A1,
				"endAngle":   arc.A2,
				"thickness": arc.Thickness,
			}
			shapes = append(shapes, shape)
		}
	}
	return shapes
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
