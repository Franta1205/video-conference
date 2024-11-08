package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"video-conference/routes"
)

func main() {
	router := gin.Default()

	routes.SetUpRoutes(router)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Could not run server: %v", err)
		}
	}()

	time.Sleep(time.Second)

	testWebSocket()
}

// Test WebSocket functionality
func testWebSocket() {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	// Send a test message to the WebSocket server
	message := []byte("Hello, WebSocket server!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Sent message: ", string(message))

	// Wait and receive a response
	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}
	fmt.Println("Received message: ", string(response))
}
