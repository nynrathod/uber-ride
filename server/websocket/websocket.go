package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// MessageObject represents a basic chat message object.
type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

// SetupWebSocket registers the WebSocket endpoint and event handlers.
// Note: Do NOT call app.Listen here—server startup is handled in main.go.
func SetupWebSocket(app *fiber.App, db *gorm.DB) {
	clients := make(map[string]string)

	// Setup event handlers on the socketio instance.
	// Using socketio.On to register events.
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		fmt.Printf("Connection event 1 - User: %s\n", ep.Kws.GetStringAttribute("user_id"))
	})

	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {
		fmt.Printf("Message event - User: %s - Message: %s\n", ep.Kws.GetStringAttribute("user_id"), string(ep.Data))
		var message MessageObject
		if err := json.Unmarshal(ep.Data, &message); err != nil {
			fmt.Println("Error unmarshaling message:", err)
			return
		}
		if targetUUID, ok := clients[message.To]; ok {
			if err := ep.Kws.EmitTo(targetUUID, ep.Data); err != nil {
				fmt.Println("Error emitting message:", err)
			}
		}
	})
	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		userID := ep.Kws.GetStringAttribute("user_id")
		delete(clients, userID)
		fmt.Printf("Disconnection event - User: %s\n", userID)
	})
	socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {
		userID := ep.Kws.GetStringAttribute("user_id")
		delete(clients, userID)
		fmt.Printf("Close event - User: %s\n", userID)
	})
	socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
		fmt.Printf("Error event - User: %s\n", ep.Kws.GetStringAttribute("user_id"))
	})

	// Register the WebSocket endpoint.
	app.Get("/ws/:id", socketio.New(func(kws *socketio.Websocket) {
		userId := kws.Params("id")
		clients[userId] = kws.UUID
		kws.SetAttribute("user_id", userId)
		kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true)
		kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)))
	}))

	log.Println("✅ WebSocket endpoint registered at /ws/:id")
}
