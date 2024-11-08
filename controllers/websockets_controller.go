package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebsocketsController struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewWebsocketsController() *WebsocketsController {
	return &WebsocketsController{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (wc *WebsocketsController) HandleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("websocket connection error: ", err)
		return
	}
	defer conn.Close()

	wc.mu.Lock()
	wc.clients[conn] = true
	wc.mu.Unlock()

	fmt.Println("guest connected to room")

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("receiving messages err: ", err)
		}

		err = wc.broadcastToRoom(messageType, data, conn)
		if err != nil {
			fmt.Println("Broadcast error: ", err)
			break
		}

		wc.removeClient(conn)
	}
}

func (wc *WebsocketsController) broadcastToRoom(messageType int, data []byte, sender *websocket.Conn) error {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	for client := range wc.clients {
		if client != sender {
			if err := client.WriteMessage(messageType, data); err != nil {
				fmt.Println("Broadcast error: ", err)
				client.Close()
				delete(wc.clients, client)
			}
		}
	}
	return nil
}

func (wc *WebsocketsController) removeClient(client *websocket.Conn) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	delete(wc.clients, client)
	fmt.Println("gest disconnected from the room")
}
