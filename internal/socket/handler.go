package socket

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/internal/service"
	"fmt"

	"github.com/zishang520/socket.io/v2/socket"
)

// func RegisterHandlers() {
// 	IO.On("connection", func(clients ...any) {
// 		client := clients[0].(*socket.Socket)
// 		fmt.Println("✅ New client connected:", client.Id())

// 		client.On("chat message", func(args ...any) {
// 			data := args[0].(map[string]any)
// 			content := data["content"].(string)
// 			conversationID := uint(data["conversationId"].(float64)) // HTML gửi number
// 			messageType := data["type"].(string)

// 			fmt.Printf("📩 Message received: %s (conversation %d)\n", content, conversationID)

// 			// Test broadcast lại
// 			room := fmt.Sprintf("conversation_%d", conversationID)
// 			IO.To(socket.Room(room)).Emit("chat message", map[string]any{
// 				"senderId":       123, // test cứng
// 				"conversationId": conversationID,
// 				"content":        content,
// 				"type":           messageType,
// 			})
// 		})

// 		client.On("disconnect", func(...any) {
// 			fmt.Println("❌ Client disconnected:", client.Id())
// 		})
// 	})
// }

func RegisterHandlers() {
	IO.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		fmt.Println("✅ New client connected:", client.Id())

		// Khi client join 1 room (theo conversationId)
		client.On("join room", func(args ...any) {
			room := args[0].(string)
			fmt.Println("➡️ Joining room:", room)
			client.Join(socket.Room(room))
		})
		// client.On("join room by userId", func(args ...any) {
		// 	room := args[0].(string)
		// 	fmt.Println("➡️ Joining room by userID:", room)
		// 	client.Join(socket.Room(room))
		// })

		// client.On("join room by conversationId", func(args ...any) {
		// 	room := args[0].(string)
		// 	fmt.Println("➡️ Joining room by conversationsID:", room)
		// 	client.Join(socket.Room(room))
		// })
		// Khi client gửi tin nhắn
		client.On("chat message", func(args ...any) {
			messageRepo := reporitory.NewMessageRepository()
			messageService := service.NewMessageService(messageRepo)

			data := args[0].(map[string]any)
			content := data["content"].(string)
			conversationID := uint(data["conversationId"].(float64))
			senderId := uint(data["senderId"].(float64))
			messageType := data["type"].(string)

			room := fmt.Sprintf("conversation_%d", conversationID)

			fmt.Printf("📩 Message: [%s] in room %s from id %d\n", content, room, senderId)

			_, err := messageService.Create(dto.CreateMessageInputDTO{
				Content:        content,
				ConversationId: fmt.Sprintf("%d", conversationID),
				SenderId:       fmt.Sprintf("%d", senderId)})
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			IO.To(socket.Room(room)).Emit("chat message", map[string]any{
				"senderId":       senderId, // giả lập user
				"conversationId": conversationID,
				"content":        content,
				"type":           messageType,
			})
			fmt.Println("📤 Emitted to room:", room)
		})

		client.On("disconnect", func(...any) {
			fmt.Println("❌ Client disconnected:", client.Id())
		})
	})
}
